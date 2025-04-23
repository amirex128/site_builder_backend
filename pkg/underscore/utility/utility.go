package utility

import (
	"bytes"
	"html/template"
	"math/rand"
	"reflect"
	"strings"
	"time"
)

// NoConflict returns a reference to the Underscore object
// Note: This is a no-op in Go since there's no global namespace like in JavaScript
func NoConflict() string {
	return "underscore"
}

// Identity returns the same value that is used as the argument
func Identity[T any](value T) T {
	return value
}

// Constant returns a function that returns value
func Constant[T any](value T) func() T {
	return func() T {
		return value
	}
}

// Noop does nothing and returns nothing
func Noop() {
	// Do nothing
}

// Times invokes the iteratee n times, returning an array of the results
func Times[T any](n int, iteratee func(int) T) []T {
	result := make([]T, n)
	for i := 0; i < n; i++ {
		result[i] = iteratee(i)
	}
	return result
}

// Random returns a random integer between min and max (inclusive)
func Random(min, max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(max-min+1) + min
}

// Mixin adds functions to the underscore object
// Note: This is a no-op in Go since we're not extending a global object
func Mixin(obj map[string]interface{}) map[string]interface{} {
	return obj
}

// Iteratee returns a function that can be used to iterate over a collection
func Iteratee(value interface{}) func(interface{}) bool {
	if value == nil {
		return func(obj interface{}) bool {
			return obj != nil
		}
	}
	
	val := reflect.ValueOf(value)
	
	// If it's a function, return it
	if val.Kind() == reflect.Func {
		return func(obj interface{}) bool {
			args := []reflect.Value{reflect.ValueOf(obj)}
			result := val.Call(args)
			if len(result) > 0 && result[0].Kind() == reflect.Bool {
				return result[0].Bool()
			}
			return false
		}
	}
	
	// If it's a map, check if the object matches the properties
	if val.Kind() == reflect.Map {
		return func(obj interface{}) bool {
			objVal := reflect.ValueOf(obj)
			
			for _, key := range val.MapKeys() {
				if key.Kind() == reflect.String {
					keyStr := key.String()
					
					var objValue interface{}
					if objVal.Kind() == reflect.Map {
						mapVal := objVal.MapIndex(key)
						if !mapVal.IsValid() {
							return false
						}
						objValue = mapVal.Interface()
					} else if objVal.Kind() == reflect.Struct {
						field := objVal.FieldByName(keyStr)
						if !field.IsValid() || !field.CanInterface() {
							return false
						}
						objValue = field.Interface()
					} else {
						return false
					}
					
					// Check if the values are equal
					if !reflect.DeepEqual(objValue, val.MapIndex(key).Interface()) {
						return false
					}
				}
			}
			
			return true
		}
	}
	
	// If it's a string, return a function that checks if the object has that property
	if val.Kind() == reflect.String {
		return func(obj interface{}) bool {
			objVal := reflect.ValueOf(obj)
			
			if objVal.Kind() == reflect.Map {
				mapVal := objVal.MapIndex(val)
				return mapVal.IsValid()
			} else if objVal.Kind() == reflect.Struct {
				field := objVal.FieldByName(val.String())
				return field.IsValid()
			}
			
			return false
		}
	}
	
	// For other types, check for equality
	return func(obj interface{}) bool {
		return reflect.DeepEqual(obj, value)
	}
}

// UniqueId returns a unique ID with an optional prefix
func UniqueId(prefix ...string) string {
	id := time.Now().UnixNano()
	
	if len(prefix) > 0 {
		return prefix[0] + "_" + string(id)
	}
	
	return string(id)
}

// Escape escapes HTML special characters in a string
func Escape(str string) string {
	return template.HTMLEscapeString(str)
}

// Unescape unescapes HTML special characters in a string
func Unescape(str string) string {
	return template.JSEscapeString(str)
}

// Result extracts the value of a property from an object
func Result(obj interface{}, property string, defaultValue ...interface{}) interface{} {
	val := reflect.ValueOf(obj)
	
	if val.Kind() == reflect.Map {
		propVal := val.MapIndex(reflect.ValueOf(property))
		if propVal.IsValid() {
			// If the property is a function, call it
			if propVal.Kind() == reflect.Func {
				result := propVal.Call([]reflect.Value{val})
				if len(result) > 0 {
					return result[0].Interface()
				}
			} else {
				return propVal.Interface()
			}
		}
	} else if val.Kind() == reflect.Struct {
		field := val.FieldByName(property)
		if field.IsValid() && field.CanInterface() {
			return field.Interface()
		}
		
		method := val.MethodByName(property)
		if method.IsValid() {
			result := method.Call([]reflect.Value{})
			if len(result) > 0 {
				return result[0].Interface()
			}
		}
	}
	
	// Return the default value if provided
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	
	return nil
}

// Now returns the current timestamp
func Now() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// Template compiles a template into a function that can be evaluated
func Template(text string, settings ...map[string]interface{}) func(map[string]interface{}) string {
	// Default settings
	templateSettings := map[string]string{
		"interpolate": "{{(.+?)}}",
		"escape":      "{{-(.+?)}}",
		"evaluate":    "{{=(.+?)}}",
	}
	
	// Override with user settings
	if len(settings) > 0 {
		for key, value := range settings[0] {
			if strValue, ok := value.(string); ok {
				templateSettings[key] = strValue
			}
		}
	}
	
	// Replace template tags with Go template syntax
	processedText := text
	
	// Replace evaluate tags
	processedText = strings.ReplaceAll(processedText, templateSettings["evaluate"], "{{$1}}")
	
	// Replace escape tags
	processedText = strings.ReplaceAll(processedText, templateSettings["escape"], "{{html $1}}")
	
	// Replace interpolate tags
	processedText = strings.ReplaceAll(processedText, templateSettings["interpolate"], "{{$1}}")
	
	// Compile the template
	tmpl, err := template.New("underscore").Funcs(template.FuncMap{
		"html": template.HTMLEscapeString,
	}).Parse(processedText)
	
	if err != nil {
		return func(map[string]interface{}) string {
			return err.Error()
		}
	}
	
	return func(data map[string]interface{}) string {
		var buf bytes.Buffer
		err := tmpl.Execute(&buf, data)
		if err != nil {
			return err.Error()
		}
		return buf.String()
	}
} 