package elasticsearch

import "time"

// Option -.
type Option func(*Elasticsearch)

// ConnectTimeout -.
func ConnectTimeout(timeout time.Duration) Option {
	return func(es *Elasticsearch) {
		es.connectTimeout = timeout
	}
}

// HealthcheckInterval -.
func HealthcheckInterval(interval time.Duration) Option {
	return func(es *Elasticsearch) {
		es.healthcheckInterval = interval
	}
}

// SniffEnabled -.
func SniffEnabled(enabled bool) Option {
	return func(es *Elasticsearch) {
		es.sniffEnabled = enabled
	}
}

// EnableIndex -.
func EnableIndex(index string) Option {
	return func(es *Elasticsearch) {
		if es.indexes == nil {
			es.indexes = make(map[string]bool)
		}
		es.indexes[index] = true
	}
}

// Retries -.
func Retries(retries int) Option {
	return func(es *Elasticsearch) {
		es.retries = retries
	}
}

// RetryBackoff -.
func RetryBackoff(backoff time.Duration) Option {
	return func(es *Elasticsearch) {
		es.retryBackoff = backoff
	}
}
