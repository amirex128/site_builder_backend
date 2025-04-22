// Package elasticsearch implements elasticsearch connection.
package elasticsearch

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"site_builder_backend/configs"
	"site_builder_backend/pkg/logger"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

const (
	_defaultConnectTimeout      = 5 * time.Second
	_defaultHealthcheckInterval = 10 * time.Second
	_defaultRetries             = 3
	_defaultRetryBackoff        = 1 * time.Second
)

// Elasticsearch -.
type Elasticsearch struct {
	client              *elasticsearch.Client
	logger              logger.Logger
	esConfig            configs.Elasticsearch
	indexes             map[string]bool
	connectTimeout      time.Duration
	healthcheckInterval time.Duration
	sniffEnabled        bool
	retries             int
	retryBackoff        time.Duration
}

// New creates a new Elasticsearch instance
func New(cfg configs.Elasticsearch, log logger.Logger, opts ...Option) (*Elasticsearch, error) {
	es := &Elasticsearch{
		logger:              log,
		esConfig:            cfg,
		indexes:             make(map[string]bool),
		connectTimeout:      _defaultConnectTimeout,
		healthcheckInterval: _defaultHealthcheckInterval,
		sniffEnabled:        cfg.SniffEnabled,
		retries:             _defaultRetries,
		retryBackoff:        _defaultRetryBackoff,
	}

	// Apply options
	for _, opt := range opts {
		opt(es)
	}

	// Register enabled indexes
	for _, index := range cfg.EnabledIndexes {
		es.indexes[index] = true
	}

	// Initialize Elasticsearch client
	esConfig := elasticsearch.Config{
		Addresses: cfg.Addresses,
		Username:  cfg.Username,
		Password:  cfg.Password,
		Transport: &http.Transport{
			DisableKeepAlives: false,
		},
		RetryOnStatus: []int{502, 503, 504, 429},
		DisableRetry:  false,
		MaxRetries:    es.retries,
		RetryBackoff: func(i int) time.Duration {
			return es.retryBackoff * time.Duration(i)
		},
		Logger: &esLogger{
			errorLogger: newESErrorLogger(log),
			infoLogger:  newESInfoLogger(log),
		},
	}

	client, err := elasticsearch.NewClient(esConfig)
	if err != nil {
		return nil, fmt.Errorf("elasticsearch - failed to create client: %w", err)
	}
	es.client = client

	// Verify connection
	ctx, cancel := context.WithTimeout(context.Background(), es.connectTimeout)
	defer cancel()

	res, err := client.Info(
		client.Info.WithContext(ctx),
	)
	if err != nil {
		return nil, fmt.Errorf("elasticsearch - connection failed: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, fmt.Errorf("elasticsearch - connection failed with status: %s", res.Status())
	}

	// Parse response to get version
	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, fmt.Errorf("elasticsearch - error parsing response body: %w", err)
	}

	version := r["version"].(map[string]interface{})["number"].(string)
	log.Info("Elasticsearch connected successfully. Version: %s", version)

	return es, nil
}

// GetClient returns the Elasticsearch client
func (es *Elasticsearch) GetClient() *elasticsearch.Client {
	return es.client
}

// IsIndexEnabled checks if an index is enabled
func (es *Elasticsearch) IsIndexEnabled(index string) bool {
	enabled, exists := es.indexes[index]
	return exists && enabled
}

// EnabledIndexes returns a list of enabled indexes
func (es *Elasticsearch) EnabledIndexes() []string {
	indexes := make([]string, 0, len(es.indexes))
	for index, enabled := range es.indexes {
		if enabled {
			indexes = append(indexes, index)
		}
	}
	return indexes
}

// CreateIndexIfNotExists creates an index if it doesn't exist
func (es *Elasticsearch) CreateIndexIfNotExists(ctx context.Context, index string, mapping string) error {
	if !es.IsIndexEnabled(index) {
		return fmt.Errorf("elasticsearch - index %s is not enabled", index)
	}

	// Check if index exists
	req := esapi.IndicesExistsRequest{
		Index: []string{index},
	}
	res, err := req.Do(ctx, es.client)
	if err != nil {
		return fmt.Errorf("elasticsearch - failed to check if index exists: %w", err)
	}
	defer res.Body.Close()

	// If status code is 404, index doesn't exist
	if res.StatusCode == 404 {
		// Create index
		createReq := esapi.IndicesCreateRequest{
			Index: index,
			Body:  strings.NewReader(mapping),
		}
		createRes, err := createReq.Do(ctx, es.client)
		if err != nil {
			return fmt.Errorf("elasticsearch - failed to create index: %w", err)
		}
		defer createRes.Body.Close()

		if createRes.IsError() {
			return fmt.Errorf("elasticsearch - failed to create index: %s", createRes.String())
		}
		es.logger.Info("Elasticsearch index %s created", index)
	} else if res.IsError() {
		return fmt.Errorf("elasticsearch - failed to check if index exists: %s", res.String())
	}

	return nil
}

// Close closes the elasticsearch client
func (es *Elasticsearch) Close() {
	// Elasticsearch client doesn't require explicit closing
}

// Logger implementation for go-elasticsearch
type esLogger struct {
	errorLogger *esErrorLogger
	infoLogger  *esInfoLogger
}

func (l *esLogger) LogRoundTrip(
	req *http.Request,
	res *http.Response,
	err error,
	start time.Time,
	dur time.Duration,
) error {
	if err != nil {
		l.errorLogger.Printf("Request failed: %s %s: %s", req.Method, req.URL.String(), err)
	} else {
		l.infoLogger.Printf(
			"%s %s [status:%d duration:%s]",
			req.Method,
			req.URL.String(),
			res.StatusCode,
			dur.Truncate(time.Millisecond),
		)
	}
	return nil
}

func (l *esLogger) RequestBodyEnabled() bool  { return false }
func (l *esLogger) ResponseBodyEnabled() bool { return false }

// Custom loggers to integrate with our logger interface
type esErrorLogger struct {
	logger logger.Logger
}

func newESErrorLogger(log logger.Logger) *esErrorLogger {
	return &esErrorLogger{logger: log}
}

func (l *esErrorLogger) Printf(format string, v ...interface{}) {
	l.logger.Error("Elasticsearch: "+format, v...)
}

type esInfoLogger struct {
	logger logger.Logger
}

func newESInfoLogger(log logger.Logger) *esInfoLogger {
	return &esInfoLogger{logger: log}
}

func (l *esInfoLogger) Printf(format string, v ...interface{}) {
	l.logger.Debug("Elasticsearch: "+format, v...)
}
