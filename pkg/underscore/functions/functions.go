package functions

import (
	"reflect"
	"sync"
	"time"
)

// Bind creates a function that, when called, invokes func with the this binding of thisArg and prepends any additional arguments to those provided to the bound function
func Bind(fn interface{}, thisArg interface{}, args ...interface{}) func(...interface{}) []reflect.Value {
	return func(callArgs ...interface{}) []reflect.Value {
		// Combine the bound args with the call args
		combinedArgs := append(args, callArgs...)

		// Convert args to reflect.Value
		reflectArgs := make([]reflect.Value, 0, len(combinedArgs)+1)

		// Add thisArg as the first argument if it's not nil
		if thisArg != nil {
			reflectArgs = append(reflectArgs, reflect.ValueOf(thisArg))
		}

		// Add the rest of the arguments
		for _, arg := range combinedArgs {
			reflectArgs = append(reflectArgs, reflect.ValueOf(arg))
		}

		// Call the function
		return reflect.ValueOf(fn).Call(reflectArgs)
	}
}

// BindAll binds methods of an object to the object itself
func BindAll(obj interface{}, methodNames ...string) {
	objVal := reflect.ValueOf(obj)

	// If no method names are provided, bind all methods
	if len(methodNames) == 0 {
		objType := objVal.Type()
		for i := 0; i < objType.NumMethod(); i++ {
			method := objType.Method(i)
			methodNames = append(methodNames, method.Name)
		}
	}

	// Bind each method
	for _, name := range methodNames {
		method := objVal.MethodByName(name)
		if method.IsValid() {
			bound := Bind(method.Interface(), obj)

			// Set the bound method back to the object
			// Note: This requires the object to be a pointer to a struct with exported fields
			if objVal.Kind() == reflect.Ptr && objVal.Elem().Kind() == reflect.Struct {
				field := objVal.Elem().FieldByName(name)
				if field.IsValid() && field.CanSet() {
					field.Set(reflect.ValueOf(bound))
				}
			}
		}
	}
}

// Partial partially applies arguments to a function
func Partial(fn interface{}, args ...interface{}) func(...interface{}) []reflect.Value {
	return func(callArgs ...interface{}) []reflect.Value {
		// Combine the partial args with the call args
		combinedArgs := append(args, callArgs...)

		// Convert args to reflect.Value
		reflectArgs := make([]reflect.Value, len(combinedArgs))
		for i, arg := range combinedArgs {
			reflectArgs[i] = reflect.ValueOf(arg)
		}

		// Call the function
		return reflect.ValueOf(fn).Call(reflectArgs)
	}
}

// Memoize creates a function that memoizes the result of func
type memoizeCache struct {
	mu    sync.RWMutex
	cache map[string]interface{}
}

func Memoize(fn interface{}, hasher func(...interface{}) string) func(...interface{}) []reflect.Value {
	cache := &memoizeCache{
		cache: make(map[string]interface{}),
	}

	if hasher == nil {
		// Default hasher just converts the first argument to a string
		hasher = func(args ...interface{}) string {
			if len(args) == 0 {
				return ""
			}
			return reflect.ValueOf(args[0]).String()
		}
	}

	return func(args ...interface{}) []reflect.Value {
		key := hasher(args...)

		// Check if the result is already cached
		cache.mu.RLock()
		if result, ok := cache.cache[key]; ok {
			cache.mu.RUnlock()
			return []reflect.Value{reflect.ValueOf(result)}
		}
		cache.mu.RUnlock()

		// Call the function
		reflectArgs := make([]reflect.Value, len(args))
		for i, arg := range args {
			reflectArgs[i] = reflect.ValueOf(arg)
		}

		result := reflect.ValueOf(fn).Call(reflectArgs)

		// Cache the result
		if len(result) > 0 {
			cache.mu.Lock()
			cache.cache[key] = result[0].Interface()
			cache.mu.Unlock()
		}

		return result
	}
}

// Delay delays the execution of func until after wait milliseconds have elapsed
func Delay(fn interface{}, wait int, args ...interface{}) {
	go func() {
		time.Sleep(time.Duration(wait) * time.Millisecond)

		// Convert args to reflect.Value
		reflectArgs := make([]reflect.Value, len(args))
		for i, arg := range args {
			reflectArgs[i] = reflect.ValueOf(arg)
		}

		// Call the function
		reflect.ValueOf(fn).Call(reflectArgs)
	}()
}

// Defer defers executing the function until the current call stack has cleared
func Defer(fn interface{}, args ...interface{}) {
	go func() {
		// Convert args to reflect.Value
		reflectArgs := make([]reflect.Value, len(args))
		for i, arg := range args {
			reflectArgs[i] = reflect.ValueOf(arg)
		}

		// Call the function
		reflect.ValueOf(fn).Call(reflectArgs)
	}()
}

// Throttle creates a function that, when invoked repeatedly, will only actually call the original function at most once per every wait milliseconds
func Throttle(fn interface{}, wait int) func(...interface{}) []reflect.Value {
	var lastCall time.Time
	var mutex sync.Mutex
	var result []reflect.Value

	return func(args ...interface{}) []reflect.Value {
		mutex.Lock()
		defer mutex.Unlock()

		now := time.Now()

		// If this is the first call or enough time has elapsed since the last call
		if lastCall.IsZero() || now.Sub(lastCall) >= time.Duration(wait)*time.Millisecond {
			// Convert args to reflect.Value
			reflectArgs := make([]reflect.Value, len(args))
			for i, arg := range args {
				reflectArgs[i] = reflect.ValueOf(arg)
			}

			// Call the function
			result = reflect.ValueOf(fn).Call(reflectArgs)
			lastCall = now
		}

		return result
	}
}

// Debounce creates a function that will delay invoking func until after wait milliseconds have elapsed since the last time it was invoked
func Debounce(fn interface{}, wait int) func(...interface{}) {
	var timer *time.Timer
	var mutex sync.Mutex

	return func(args ...interface{}) {
		mutex.Lock()
		defer mutex.Unlock()

		// Cancel the previous timer
		if timer != nil {
			timer.Stop()
		}

		// Start a new timer
		timer = time.AfterFunc(time.Duration(wait)*time.Millisecond, func() {
			// Convert args to reflect.Value
			reflectArgs := make([]reflect.Value, len(args))
			for i, arg := range args {
				reflectArgs[i] = reflect.ValueOf(arg)
			}

			// Call the function
			reflect.ValueOf(fn).Call(reflectArgs)
		})
	}
}

// Once creates a function that is restricted to invoking func once
func Once(fn interface{}) func(...interface{}) []reflect.Value {
	var once sync.Once
	var result []reflect.Value

	return func(args ...interface{}) []reflect.Value {
		once.Do(func() {
			// Convert args to reflect.Value
			reflectArgs := make([]reflect.Value, len(args))
			for i, arg := range args {
				reflectArgs[i] = reflect.ValueOf(arg)
			}

			// Call the function
			result = reflect.ValueOf(fn).Call(reflectArgs)
		})

		return result
	}
}

// After creates a function that invokes func once it's called n or more times
func After(n int, fn interface{}) func(...interface{}) []reflect.Value {
	var count int
	var mutex sync.Mutex

	return func(args ...interface{}) []reflect.Value {
		mutex.Lock()
		defer mutex.Unlock()

		count++
		if count >= n {
			// Convert args to reflect.Value
			reflectArgs := make([]reflect.Value, len(args))
			for i, arg := range args {
				reflectArgs[i] = reflect.ValueOf(arg)
			}

			// Call the function
			return reflect.ValueOf(fn).Call(reflectArgs)
		}

		return nil
	}
}

// Before creates a function that invokes func, with the this binding and arguments of the created function, while it's called less than n times
func Before(n int, fn interface{}) func(...interface{}) []reflect.Value {
	var count int
	var mutex sync.Mutex
	var result []reflect.Value

	return func(args ...interface{}) []reflect.Value {
		mutex.Lock()
		defer mutex.Unlock()

		count++
		if count < n {
			// Convert args to reflect.Value
			reflectArgs := make([]reflect.Value, len(args))
			for i, arg := range args {
				reflectArgs[i] = reflect.ValueOf(arg)
			}

			// Call the function
			result = reflect.ValueOf(fn).Call(reflectArgs)
		}

		return result
	}
}

// Wrap creates a function that provides value to the wrapper function as its first argument
func Wrap(value interface{}, wrapper interface{}) func(...interface{}) []reflect.Value {
	return func(args ...interface{}) []reflect.Value {
		// Prepend the value to the arguments
		newArgs := make([]interface{}, len(args)+1)
		newArgs[0] = value
		copy(newArgs[1:], args)

		// Convert args to reflect.Value
		reflectArgs := make([]reflect.Value, len(newArgs))
		for i, arg := range newArgs {
			reflectArgs[i] = reflect.ValueOf(arg)
		}

		// Call the wrapper function
		return reflect.ValueOf(wrapper).Call(reflectArgs)
	}
}

// Negate creates a function that negates the result of the predicate func
func Negate(predicate interface{}) func(...interface{}) bool {
	return func(args ...interface{}) bool {
		// Convert args to reflect.Value
		reflectArgs := make([]reflect.Value, len(args))
		for i, arg := range args {
			reflectArgs[i] = reflect.ValueOf(arg)
		}

		// Call the predicate function
		result := reflect.ValueOf(predicate).Call(reflectArgs)

		// Negate the result
		if len(result) > 0 && result[0].Kind() == reflect.Bool {
			return !result[0].Bool()
		}

		return false
	}
}

// Compose creates a function that is the composition of the provided functions
func Compose(funcs ...interface{}) func(...interface{}) []reflect.Value {
	return func(args ...interface{}) []reflect.Value {
		if len(funcs) == 0 {
			return nil
		}

		// Convert initial args to reflect.Value
		reflectArgs := make([]reflect.Value, len(args))
		for i, arg := range args {
			reflectArgs[i] = reflect.ValueOf(arg)
		}

		// Call the first function
		result := reflect.ValueOf(funcs[len(funcs)-1]).Call(reflectArgs)

		// Call each function in reverse order, passing the result of the previous function
		for i := len(funcs) - 2; i >= 0; i-- {
			// Convert the result to a slice of reflect.Value
			newArgs := make([]reflect.Value, len(result))
			copy(newArgs, result)

			// Call the next function
			result = reflect.ValueOf(funcs[i]).Call(newArgs)
		}

		return result
	}
}

// RestArguments creates a function that invokes func with the this binding of thisArg and arguments of the created function
func RestArguments(fn interface{}) func(...interface{}) []reflect.Value {
	return func(args ...interface{}) []reflect.Value {
		// Convert args to reflect.Value
		reflectArgs := make([]reflect.Value, len(args))
		for i, arg := range args {
			reflectArgs[i] = reflect.ValueOf(arg)
		}

		// Call the function
		return reflect.ValueOf(fn).Call(reflectArgs)
	}
}
