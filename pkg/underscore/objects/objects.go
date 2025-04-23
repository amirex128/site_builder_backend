package objects

import (
	"fmt"
	"math"
	"reflect"
	"regexp"
	"sort"
	"strings"
	"time"
)

// Keys returns all the keys of an object
func Keys(obj interface{}) []string {
	val := reflect.ValueOf(obj)
	
	if val.Kind() == reflect.Map {
		keys := val.MapKeys()
		result := make([]string, 0, len(keys))
		
		for _, key := range keys {
			if key.Kind() == reflect.String {
				result = append(result, key.String())
			}
		}
		
		return result
	} else if val.Kind() == reflect.Struct {
		typ := val.Type()
		result := make([]string, 0, typ.NumField())
		
		for i := 0; i < typ.NumField(); i++ {
			field := typ.Field(i)
			if field.IsExported() {
				result = append(result, field.Name)
			}
		}
		
		return result
	}
	
	return []string{}
}

// AllKeys returns all the keys of an object, including inherited ones
func AllKeys(obj interface{}) []string {
	// In Go, there's no direct equivalent to JavaScript's prototype chain
	// So we'll just return all keys, including unexported ones for structs
	val := reflect.ValueOf(obj)
	
	if val.Kind() == reflect.Map {
		keys := val.MapKeys()
		result := make([]string, 0, len(keys))
		
		for _, key := range keys {
			if key.Kind() == reflect.String {
				result = append(result, key.String())
			}
		}
		
		return result
	} else if val.Kind() == reflect.Struct {
		typ := val.Type()
		result := make([]string, 0, typ.NumField())
		
		for i := 0; i < typ.NumField(); i++ {
			field := typ.Field(i)
			result = append(result, field.Name)
		}
		
		return result
	}
	
	return []string{}
}

// Values returns all the values of an object
func Values(obj interface{}) []interface{} {
	val := reflect.ValueOf(obj)
	
	if val.Kind() == reflect.Map {
		keys := val.MapKeys()
		result := make([]interface{}, 0, len(keys))
		
		for _, key := range keys {
			result = append(result, val.MapIndex(key).Interface())
		}
		
		return result
	} else if val.Kind() == reflect.Struct {
		result := make([]interface{}, 0, val.NumField())
		
		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)
			if field.CanInterface() {
				result = append(result, field.Interface())
			}
		}
		
		return result
	}
	
	return []interface{}{}
}

// MapObject maps an object to a new object with the same keys
func MapObject(obj interface{}, iteratee func(value interface{}, key string, obj interface{}) interface{}) map[string]interface{} {
	val := reflect.ValueOf(obj)
	result := make(map[string]interface{})
	
	if val.Kind() == reflect.Map {
		for _, key := range val.MapKeys() {
			if key.Kind() == reflect.String {
				keyStr := key.String()
				value := val.MapIndex(key).Interface()
				result[keyStr] = iteratee(value, keyStr, obj)
			}
		}
	} else if val.Kind() == reflect.Struct {
		typ := val.Type()
		for i := 0; i < typ.NumField(); i++ {
			field := typ.Field(i)
			if field.IsExported() {
				fieldVal := val.Field(i)
				if fieldVal.CanInterface() {
					value := fieldVal.Interface()
					result[field.Name] = iteratee(value, field.Name, obj)
				}
			}
		}
	}
	
	return result
}

// Pairs converts an object into a list of [key, value] pairs
func Pairs(obj interface{}) [][]interface{} {
	val := reflect.ValueOf(obj)
	
	if val.Kind() == reflect.Map {
		keys := val.MapKeys()
		result := make([][]interface{}, 0, len(keys))
		
		for _, key := range keys {
			if key.Kind() == reflect.String {
				pair := []interface{}{key.String(), val.MapIndex(key).Interface()}
				result = append(result, pair)
			}
		}
		
		return result
	} else if val.Kind() == reflect.Struct {
		typ := val.Type()
		result := make([][]interface{}, 0, val.NumField())
		
		for i := 0; i < typ.NumField(); i++ {
			field := typ.Field(i)
			if field.IsExported() {
				fieldVal := val.Field(i)
				if fieldVal.CanInterface() {
					pair := []interface{}{field.Name, fieldVal.Interface()}
					result = append(result, pair)
				}
			}
		}
		
		return result
	}
	
	return [][]interface{}{}
}

// Invert creates an object with the keys and values swapped
func Invert(obj interface{}) map[string]string {
	val := reflect.ValueOf(obj)
	result := make(map[string]string)
	
	if val.Kind() == reflect.Map {
		for _, key := range val.MapKeys() {
			if key.Kind() == reflect.String {
				keyStr := key.String()
				value := val.MapIndex(key)
				if value.Kind() == reflect.String {
					result[value.String()] = keyStr
				} else {
					result[value.String()] = keyStr
				}
			}
		}
	} else if val.Kind() == reflect.Struct {
		typ := val.Type()
		for i := 0; i < typ.NumField(); i++ {
			field := typ.Field(i)
			if field.IsExported() {
				fieldVal := val.Field(i)
				if fieldVal.CanInterface() {
					result[fieldVal.String()] = field.Name
				}
			}
		}
	}
	
	return result
}

// Create creates a new object with the given prototype
// Note: Go doesn't have prototypes like JavaScript, so this is a simplified version
func Create(prototype interface{}, properties map[string]interface{}) map[string]interface{} {
	// Create a new object with the properties from the prototype
	result := make(map[string]interface{})
	
	// Copy properties from the prototype
	protoVal := reflect.ValueOf(prototype)
	if protoVal.Kind() == reflect.Map {
		for _, key := range protoVal.MapKeys() {
			if key.Kind() == reflect.String {
				result[key.String()] = protoVal.MapIndex(key).Interface()
			}
		}
	} else if protoVal.Kind() == reflect.Struct {
		typ := protoVal.Type()
		for i := 0; i < typ.NumField(); i++ {
			field := typ.Field(i)
			if field.IsExported() {
				fieldVal := protoVal.Field(i)
				if fieldVal.CanInterface() {
					result[field.Name] = fieldVal.Interface()
				}
			}
		}
	}
	
	// Add the new properties
	for key, value := range properties {
		result[key] = value
	}
	
	return result
}

// Functions returns a sorted list of function names in an object
func Functions(obj interface{}) []string {
	val := reflect.ValueOf(obj)
	result := make([]string, 0)
	
	if val.Kind() == reflect.Struct || val.Kind() == reflect.Ptr {
		typ := val.Type()
		for i := 0; i < typ.NumMethod(); i++ {
			method := typ.Method(i)
			result = append(result, method.Name)
		}
	}
	
	sort.Strings(result)
	return result
}

// FindKey returns the first key where the predicate returns true
func FindKey(obj interface{}, predicate func(value interface{}, key string, obj interface{}) bool) (string, bool) {
	val := reflect.ValueOf(obj)
	
	if val.Kind() == reflect.Map {
		for _, key := range val.MapKeys() {
			if key.Kind() == reflect.String {
				keyStr := key.String()
				value := val.MapIndex(key).Interface()
				if predicate(value, keyStr, obj) {
					return keyStr, true
				}
			}
		}
	} else if val.Kind() == reflect.Struct {
		typ := val.Type()
		for i := 0; i < typ.NumField(); i++ {
			field := typ.Field(i)
			if field.IsExported() {
				fieldVal := val.Field(i)
				if fieldVal.CanInterface() {
					value := fieldVal.Interface()
					if predicate(value, field.Name, obj) {
						return field.Name, true
					}
				}
			}
		}
	}
	
	return "", false
}

// Extend copies all properties from the source objects to the destination object
func Extend(destination interface{}, sources ...interface{}) {
	destVal := reflect.ValueOf(destination)
	
	// Only works if destination is a pointer to a map or struct
	if destVal.Kind() != reflect.Ptr {
		return
	}
	
	destVal = destVal.Elem()
	
	if destVal.Kind() == reflect.Map {
		// For maps, we can directly set values
		for _, source := range sources {
			sourceVal := reflect.ValueOf(source)
			if sourceVal.Kind() == reflect.Map {
				for _, key := range sourceVal.MapKeys() {
					destVal.SetMapIndex(key, sourceVal.MapIndex(key))
				}
			} else if sourceVal.Kind() == reflect.Struct {
				typ := sourceVal.Type()
				for i := 0; i < typ.NumField(); i++ {
					field := typ.Field(i)
					if field.IsExported() {
						fieldVal := sourceVal.Field(i)
						if fieldVal.CanInterface() {
							destVal.SetMapIndex(reflect.ValueOf(field.Name), fieldVal)
						}
					}
				}
			}
		}
	} else if destVal.Kind() == reflect.Struct {
		// For structs, we need to find matching fields
		for _, source := range sources {
			sourceVal := reflect.ValueOf(source)
			if sourceVal.Kind() == reflect.Map {
				for _, key := range sourceVal.MapKeys() {
					if key.Kind() == reflect.String {
						keyStr := key.String()
						field := destVal.FieldByName(keyStr)
						if field.IsValid() && field.CanSet() {
							sourceField := sourceVal.MapIndex(key)
							if sourceField.Type().AssignableTo(field.Type()) {
								field.Set(sourceField)
							}
						}
					}
				}
			} else if sourceVal.Kind() == reflect.Struct {
				typ := sourceVal.Type()
				for i := 0; i < typ.NumField(); i++ {
					field := typ.Field(i)
					if field.IsExported() {
						destField := destVal.FieldByName(field.Name)
						if destField.IsValid() && destField.CanSet() {
							sourceField := sourceVal.Field(i)
							if sourceField.Type().AssignableTo(destField.Type()) {
								destField.Set(sourceField)
							}
						}
					}
				}
			}
		}
	}
}

// ExtendOwn copies all own properties from the source objects to the destination object
func ExtendOwn(destination interface{}, sources ...interface{}) {
	// In Go, there's no distinction between "own" and inherited properties
	// So this is the same as Extend
	Extend(destination, sources...)
}

// Pick returns a copy of the object with only the whitelisted properties
func Pick(obj interface{}, keys ...string) map[string]interface{} {
	val := reflect.ValueOf(obj)
	result := make(map[string]interface{})
	
	// Create a set of keys for faster lookup
	keySet := make(map[string]bool)
	for _, key := range keys {
		keySet[key] = true
	}
	
	if val.Kind() == reflect.Map {
		for _, key := range val.MapKeys() {
			if key.Kind() == reflect.String {
				keyStr := key.String()
				if keySet[keyStr] {
					result[keyStr] = val.MapIndex(key).Interface()
				}
			}
		}
	} else if val.Kind() == reflect.Struct {
		typ := val.Type()
		for i := 0; i < typ.NumField(); i++ {
			field := typ.Field(i)
			if field.IsExported() && keySet[field.Name] {
				fieldVal := val.Field(i)
				if fieldVal.CanInterface() {
					result[field.Name] = fieldVal.Interface()
				}
			}
		}
	}
	
	return result
}

// Omit returns a copy of the object without the blacklisted properties
func Omit(obj interface{}, keys ...string) map[string]interface{} {
	val := reflect.ValueOf(obj)
	result := make(map[string]interface{})
	
	// Create a set of keys for faster lookup
	keySet := make(map[string]bool)
	for _, key := range keys {
		keySet[key] = true
	}
	
	if val.Kind() == reflect.Map {
		for _, key := range val.MapKeys() {
			if key.Kind() == reflect.String {
				keyStr := key.String()
				if !keySet[keyStr] {
					result[keyStr] = val.MapIndex(key).Interface()
				}
			}
		}
	} else if val.Kind() == reflect.Struct {
		typ := val.Type()
		for i := 0; i < typ.NumField(); i++ {
			field := typ.Field(i)
			if field.IsExported() && !keySet[field.Name] {
				fieldVal := val.Field(i)
				if fieldVal.CanInterface() {
					result[field.Name] = fieldVal.Interface()
				}
			}
		}
	}
	
	return result
}

// Defaults fills in undefined properties in object with values from the defaults objects
func Defaults(obj interface{}, defaults ...interface{}) {
	objVal := reflect.ValueOf(obj)
	
	// Only works if obj is a pointer to a map or struct
	if objVal.Kind() != reflect.Ptr {
		return
	}
	
	objVal = objVal.Elem()
	
	if objVal.Kind() == reflect.Map {
		// For maps, we can directly set values
		for _, defaultObj := range defaults {
			defaultVal := reflect.ValueOf(defaultObj)
			if defaultVal.Kind() == reflect.Map {
				for _, key := range defaultVal.MapKeys() {
					// Only set if the key doesn't exist or is zero value
					if !objVal.MapIndex(key).IsValid() {
						objVal.SetMapIndex(key, defaultVal.MapIndex(key))
					}
				}
			} else if defaultVal.Kind() == reflect.Struct {
				typ := defaultVal.Type()
				for i := 0; i < typ.NumField(); i++ {
					field := typ.Field(i)
					if field.IsExported() {
						key := reflect.ValueOf(field.Name)
						// Only set if the key doesn't exist or is zero value
						if !objVal.MapIndex(key).IsValid() {
							fieldVal := defaultVal.Field(i)
							if fieldVal.CanInterface() {
								objVal.SetMapIndex(key, fieldVal)
							}
						}
					}
				}
			}
		}
	} else if objVal.Kind() == reflect.Struct {
		// For structs, we need to find matching fields
		for _, defaultObj := range defaults {
			defaultVal := reflect.ValueOf(defaultObj)
			if defaultVal.Kind() == reflect.Map {
				for _, key := range defaultVal.MapKeys() {
					if key.Kind() == reflect.String {
						keyStr := key.String()
						field := objVal.FieldByName(keyStr)
						if field.IsValid() && field.CanSet() {
							// Only set if the field is zero value
							if field.IsZero() {
								sourceField := defaultVal.MapIndex(key)
								if sourceField.Type().AssignableTo(field.Type()) {
									field.Set(sourceField)
								}
							}
						}
					}
				}
			} else if defaultVal.Kind() == reflect.Struct {
				typ := defaultVal.Type()
				for i := 0; i < typ.NumField(); i++ {
					field := typ.Field(i)
					if field.IsExported() {
						destField := objVal.FieldByName(field.Name)
						if destField.IsValid() && destField.CanSet() {
							// Only set if the field is zero value
							if destField.IsZero() {
								sourceField := defaultVal.Field(i)
								if sourceField.Type().AssignableTo(destField.Type()) {
									destField.Set(sourceField)
								}
							}
						}
					}
				}
			}
		}
	}
}

// Clone creates a shallow copy of the object
func Clone(obj interface{}) interface{} {
	val := reflect.ValueOf(obj)
	
	if val.Kind() == reflect.Map {
		// Create a new map of the same type
		mapType := val.Type()
		newMap := reflect.MakeMap(mapType)
		
		// Copy all key-value pairs
		for _, key := range val.MapKeys() {
			newMap.SetMapIndex(key, val.MapIndex(key))
		}
		
		return newMap.Interface()
	} else if val.Kind() == reflect.Struct {
		// Create a new struct of the same type
		newStruct := reflect.New(val.Type()).Elem()
		
		// Copy all fields
		for i := 0; i < val.NumField(); i++ {
			if newStruct.Field(i).CanSet() {
				newStruct.Field(i).Set(val.Field(i))
			}
		}
		
		return newStruct.Interface()
	} else if val.Kind() == reflect.Slice || val.Kind() == reflect.Array {
		// Create a new slice of the same type
		sliceType := val.Type()
		newSlice := reflect.MakeSlice(sliceType, val.Len(), val.Cap())
		
		// Copy all elements
		reflect.Copy(newSlice, val)
		
		return newSlice.Interface()
	}
	
	// For other types, just return the value
	return obj
}

// Tap invokes interceptor with the object, and then returns the object
func Tap(obj interface{}, interceptor func(obj interface{})) interface{} {
	interceptor(obj)
	return obj
}

// ToPath converts a string path to an array of path segments
func ToPath(path string) []string {
	// Split the path by dots and brackets
	// This is a simplified version that doesn't handle all edge cases
	path = strings.ReplaceAll(path, "[", ".")
	path = strings.ReplaceAll(path, "]", "")
	return strings.Split(path, ".")
}

// Has checks if path is a direct property of object
func Has(obj interface{}, path string) bool {
	val := reflect.ValueOf(obj)
	segments := ToPath(path)
	
	if len(segments) == 0 {
		return false
	}
	
	// Check only the first segment for direct property
	segment := segments[0]
	
	if val.Kind() == reflect.Map {
		for _, key := range val.MapKeys() {
			if key.Kind() == reflect.String && key.String() == segment {
				return true
			}
		}
	} else if val.Kind() == reflect.Struct {
		field := val.FieldByName(segment)
		return field.IsValid()
	}
	
	return false
}

// Get retrieves the value at path of object
func Get(obj interface{}, path string) interface{} {
	val := reflect.ValueOf(obj)
	segments := ToPath(path)
	
	for _, segment := range segments {
		if val.Kind() == reflect.Map {
			val = val.MapIndex(reflect.ValueOf(segment))
			if !val.IsValid() {
				return nil
			}
		} else if val.Kind() == reflect.Struct {
			val = val.FieldByName(segment)
			if !val.IsValid() {
				return nil
			}
		} else if val.Kind() == reflect.Slice || val.Kind() == reflect.Array {
			// Try to convert segment to an index
			var index int
			_, err := fmt.Sscanf(segment, "%d", &index)
			if err != nil || index < 0 || index >= val.Len() {
				return nil
			}
			val = val.Index(index)
		} else {
			return nil
		}
	}
	
	if !val.CanInterface() {
		return nil
	}
	
	return val.Interface()
}

// Property creates a function that returns the value at path of object
func Property(path string) func(obj interface{}) interface{} {
	return func(obj interface{}) interface{} {
		return Get(obj, path)
	}
}

// PropertyOf creates a function that returns the value at a given path of object
func PropertyOf(obj interface{}) func(path string) interface{} {
	return func(path string) interface{} {
		return Get(obj, path)
	}
}

// Matcher returns a function that checks if an object matches the given properties
func Matcher(attrs map[string]interface{}) func(obj interface{}) bool {
	return func(obj interface{}) bool {
		val := reflect.ValueOf(obj)
		
		for key, attrValue := range attrs {
			var objValue interface{}
			
			if val.Kind() == reflect.Map {
				mapVal := val.MapIndex(reflect.ValueOf(key))
				if !mapVal.IsValid() {
					return false
				}
				objValue = mapVal.Interface()
			} else if val.Kind() == reflect.Struct {
				field := val.FieldByName(key)
				if !field.IsValid() || !field.CanInterface() {
					return false
				}
				objValue = field.Interface()
			} else {
				return false
			}
			
			// Check if the values are equal
			if !reflect.DeepEqual(objValue, attrValue) {
				return false
			}
		}
		
		return true
	}
}

// IsEqual performs a deep comparison between two objects
func IsEqual(a, b interface{}) bool {
	return reflect.DeepEqual(a, b)
}

// IsMatch checks if object contains all the key-value pairs in attrs
func IsMatch(obj interface{}, attrs map[string]interface{}) bool {
	return Matcher(attrs)(obj)
}

// IsEmpty checks if an object is empty
func IsEmpty(obj interface{}) bool {
	val := reflect.ValueOf(obj)
	
	switch val.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice, reflect.String:
		return val.Len() == 0
	case reflect.Bool:
		return !val.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return val.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return val.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return val.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return val.IsNil()
	}
	
	return false
}

// IsElement checks if obj is a DOM element
func IsElement(obj interface{}) bool {
	// Go doesn't have DOM elements, so this always returns false
	return false
}

// IsArray checks if obj is an array
func IsArray(obj interface{}) bool {
	val := reflect.ValueOf(obj)
	return val.Kind() == reflect.Array || val.Kind() == reflect.Slice
}

// IsObject checks if obj is an object
func IsObject(obj interface{}) bool {
	if obj == nil {
		return false
	}
	
	val := reflect.ValueOf(obj)
	return val.Kind() == reflect.Map || val.Kind() == reflect.Struct || val.Kind() == reflect.Ptr
}

// IsArguments checks if obj is an arguments object
func IsArguments(obj interface{}) bool {
	// Go doesn't have arguments objects, so this always returns false
	return false
}

// IsFunction checks if obj is a function
func IsFunction(obj interface{}) bool {
	val := reflect.ValueOf(obj)
	return val.Kind() == reflect.Func
}

// IsString checks if obj is a string
func IsString(obj interface{}) bool {
	val := reflect.ValueOf(obj)
	return val.Kind() == reflect.String
}

// IsNumber checks if obj is a number
func IsNumber(obj interface{}) bool {
	val := reflect.ValueOf(obj)
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64:
		return true
	}
	return false
}

// IsFinite checks if obj is a finite number
func IsFinite(obj interface{}) bool {
	val := reflect.ValueOf(obj)
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return true
	case reflect.Float32, reflect.Float64:
		f := val.Float()
		return !math.IsInf(f, 0) && !math.IsNaN(f)
	}
	return false
}

// IsBoolean checks if obj is a boolean
func IsBoolean(obj interface{}) bool {
	val := reflect.ValueOf(obj)
	return val.Kind() == reflect.Bool
}

// IsDate checks if obj is a Date
func IsDate(obj interface{}) bool {
	val := reflect.ValueOf(obj)
	if val.Kind() == reflect.Struct {
		_, ok := obj.(time.Time)
		return ok
	}
	return false
}

// IsRegExp checks if obj is a RegExp
func IsRegExp(obj interface{}) bool {
	val := reflect.ValueOf(obj)
	if val.Kind() == reflect.Struct {
		_, ok := obj.(*regexp.Regexp)
		return ok
	}
	return false
}

// IsError checks if obj is an Error
func IsError(obj interface{}) bool {
	_, ok := obj.(error)
	return ok
}

// IsSymbol checks if obj is a Symbol
func IsSymbol(obj interface{}) bool {
	// Go doesn't have symbols, so this always returns false
	return false
}

// IsMap checks if obj is a Map
func IsMap(obj interface{}) bool {
	val := reflect.ValueOf(obj)
	return val.Kind() == reflect.Map
}

// IsWeakMap checks if obj is a WeakMap
func IsWeakMap(obj interface{}) bool {
	// Go doesn't have WeakMaps, so this always returns false
	return false
}

// IsSet checks if obj is a Set
func IsSet(obj interface{}) bool {
	// Go doesn't have Sets, so this always returns false
	return false
}

// IsWeakSet checks if obj is a WeakSet
func IsWeakSet(obj interface{}) bool {
	// Go doesn't have WeakSets, so this always returns false
	return false
}

// IsArrayBuffer checks if obj is an ArrayBuffer
func IsArrayBuffer(obj interface{}) bool {
	// Go doesn't have ArrayBuffers, so this always returns false
	return false
}

// IsDataView checks if obj is a DataView
func IsDataView(obj interface{}) bool {
	// Go doesn't have DataViews, so this always returns false
	return false
}

// IsTypedArray checks if obj is a TypedArray
func IsTypedArray(obj interface{}) bool {
	// Go doesn't have TypedArrays, so this always returns false
	return false
}

// IsNaN checks if obj is NaN
func IsNaN(obj interface{}) bool {
	val := reflect.ValueOf(obj)
	if val.Kind() == reflect.Float32 || val.Kind() == reflect.Float64 {
		return math.IsNaN(val.Float())
	}
	return false
}

// IsNull checks if obj is null
func IsNull(obj interface{}) bool {
	return obj == nil
}

// IsUndefined checks if obj is undefined
func IsUndefined(obj interface{}) bool {
	// Go doesn't have undefined, so this checks for nil
	return obj == nil
} 