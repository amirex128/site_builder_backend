package arrays

import (
	"reflect"
	_ "sort"
)

// First returns the first n elements of an array
func First[T any](arr []T, n ...int) []T {
	if len(arr) == 0 {
		return []T{}
	}

	count := 1
	if len(n) > 0 && n[0] > 0 {
		count = n[0]
	}

	if count >= len(arr) {
		return arr
	}

	return arr[:count]
}

// Initial returns everything but the last n elements of an array
func Initial[T any](arr []T, n ...int) []T {
	if len(arr) == 0 {
		return []T{}
	}

	count := 1
	if len(n) > 0 && n[0] > 0 {
		count = n[0]
	}

	if count >= len(arr) {
		return []T{}
	}

	return arr[:len(arr)-count]
}

// Last returns the last n elements of an array
func Last[T any](arr []T, n ...int) []T {
	if len(arr) == 0 {
		return []T{}
	}

	count := 1
	if len(n) > 0 && n[0] > 0 {
		count = n[0]
	}

	if count >= len(arr) {
		return arr
	}

	return arr[len(arr)-count:]
}

// Rest returns everything but the first n elements of an array
func Rest[T any](arr []T, n ...int) []T {
	if len(arr) == 0 {
		return []T{}
	}

	count := 1
	if len(n) > 0 && n[0] > 0 {
		count = n[0]
	}

	if count >= len(arr) {
		return []T{}
	}

	return arr[count:]
}

// Compact returns a copy of the array with all falsy values removed
func Compact[T comparable](arr []T) []T {
	var zero T
	result := make([]T, 0)

	for _, item := range arr {
		if item != zero {
			result = append(result, item)
		}
	}

	return result
}

// Flatten flattens a nested array
func Flatten(arr interface{}, shallow ...bool) []interface{} {
	isShallow := false
	if len(shallow) > 0 {
		isShallow = shallow[0]
	}

	result := make([]interface{}, 0)
	flattenHelper(arr, &result, isShallow, 0)

	return result
}

func flattenHelper(arr interface{}, result *[]interface{}, shallow bool, depth int) {
	val := reflect.ValueOf(arr)

	if val.Kind() == reflect.Slice || val.Kind() == reflect.Array {
		for i := 0; i < val.Len(); i++ {
			item := val.Index(i).Interface()
			itemVal := reflect.ValueOf(item)

			if (itemVal.Kind() == reflect.Slice || itemVal.Kind() == reflect.Array) && (!shallow || depth == 0) {
				flattenHelper(item, result, shallow, depth+1)
			} else {
				*result = append(*result, item)
			}
		}
	} else {
		*result = append(*result, arr)
	}
}

// Without returns a copy of the array with all instances of the specified values removed
func Without[T comparable](arr []T, values ...T) []T {
	result := make([]T, 0)

	for _, item := range arr {
		exclude := false
		for _, value := range values {
			if item == value {
				exclude = true
				break
			}
		}

		if !exclude {
			result = append(result, item)
		}
	}

	return result
}

// Union returns the union of the arrays
func Union[T comparable](arrs ...[]T) []T {
	seen := make(map[T]bool)
	result := make([]T, 0)

	for _, arr := range arrs {
		for _, item := range arr {
			if !seen[item] {
				seen[item] = true
				result = append(result, item)
			}
		}
	}

	return result
}

// Intersection returns the intersection of the arrays
func Intersection[T comparable](arrs ...[]T) []T {
	if len(arrs) == 0 {
		return []T{}
	}

	if len(arrs) == 1 {
		return arrs[0]
	}

	// Count occurrences of each item
	counts := make(map[T]int)
	for _, arr := range arrs {
		// Use a set to avoid counting duplicates in the same array
		seen := make(map[T]bool)
		for _, item := range arr {
			if !seen[item] {
				seen[item] = true
				counts[item]++
			}
		}
	}

	// Add items that appear in all arrays
	result := make([]T, 0)
	for item, count := range counts {
		if count == len(arrs) {
			result = append(result, item)
		}
	}

	return result
}

// Difference returns the values from array that are not present in the other arrays
func Difference[T comparable](arr []T, others ...[]T) []T {
	if len(arr) == 0 {
		return []T{}
	}

	// Create a set of all items in other arrays
	exclude := make(map[T]bool)
	for _, other := range others {
		for _, item := range other {
			exclude[item] = true
		}
	}

	// Add items from arr that are not in the exclude set
	result := make([]T, 0)
	for _, item := range arr {
		if !exclude[item] {
			result = append(result, item)
		}
	}

	return result
}

// Uniq returns a duplicate-free version of the array
func Uniq[T comparable](arr []T) []T {
	seen := make(map[T]bool)
	result := make([]T, 0)

	for _, item := range arr {
		if !seen[item] {
			seen[item] = true
			result = append(result, item)
		}
	}

	return result
}

// Zip merges together the values of each of the arrays with the values at the corresponding position
func Zip[T any](arrs ...[]T) [][]T {
	if len(arrs) == 0 {
		return [][]T{}
	}

	// Find the length of the longest array
	maxLen := 0
	for _, arr := range arrs {
		if len(arr) > maxLen {
			maxLen = len(arr)
		}
	}

	// Zip the arrays
	result := make([][]T, maxLen)
	for i := 0; i < maxLen; i++ {
		result[i] = make([]T, len(arrs))
		for j, arr := range arrs {
			if i < len(arr) {
				result[i][j] = arr[i]
			}
		}
	}

	return result
}

// Unzip is the opposite of Zip
func Unzip[T any](arr [][]T) [][]T {
	if len(arr) == 0 {
		return [][]T{}
	}

	// Find the length of the longest inner array
	maxLen := 0
	for _, inner := range arr {
		if len(inner) > maxLen {
			maxLen = len(inner)
		}
	}

	// Unzip the array
	result := make([][]T, maxLen)
	for i := 0; i < maxLen; i++ {
		result[i] = make([]T, len(arr))
		for j, inner := range arr {
			if i < len(inner) {
				result[i][j] = inner[i]
			}
		}
	}

	return result
}

// Object converts arrays of key-value pairs into an object
func Object[K comparable, V any](pairs [][]interface{}) map[K]V {
	result := make(map[K]V)

	for _, pair := range pairs {
		if len(pair) >= 2 {
			if key, ok := pair[0].(K); ok {
				if val, ok := pair[1].(V); ok {
					result[key] = val
				}
			}
		}
	}

	return result
}

// Chunk splits an array into groups of size n
func Chunk[T any](arr []T, size int) [][]T {
	if size <= 0 {
		return [][]T{}
	}

	result := make([][]T, 0)
	for i := 0; i < len(arr); i += size {
		end := i + size
		if end > len(arr) {
			end = len(arr)
		}
		result = append(result, arr[i:end])
	}

	return result
}

// IndexOf returns the index at which value can be found in the array, or -1 if not found
func IndexOf[T comparable](arr []T, value T, fromIndex ...int) int {
	start := 0
	if len(fromIndex) > 0 && fromIndex[0] >= 0 {
		start = fromIndex[0]
	}

	for i := start; i < len(arr); i++ {
		if arr[i] == value {
			return i
		}
	}

	return -1
}

// LastIndexOf returns the index of the last occurrence of value in the array, or -1 if not found
func LastIndexOf[T comparable](arr []T, value T, fromIndex ...int) int {
	end := len(arr) - 1
	if len(fromIndex) > 0 && fromIndex[0] >= 0 && fromIndex[0] < len(arr) {
		end = fromIndex[0]
	}

	for i := end; i >= 0; i-- {
		if arr[i] == value {
			return i
		}
	}

	return -1
}

// SortedIndex returns the index at which value should be inserted into array to maintain order
func SortedIndex[T any](arr []T, value T, iteratee func(T) float64) int {
	low := 0
	high := len(arr)

	valueScore := iteratee(value)

	for low < high {
		mid := (low + high) / 2
		if iteratee(arr[mid]) < valueScore {
			low = mid + 1
		} else {
			high = mid
		}
	}

	return low
}

// FindIndex returns the index of the first element that passes the predicate test
func FindIndex[T any](arr []T, predicate func(T, int, []T) bool) int {
	for i, item := range arr {
		if predicate(item, i, arr) {
			return i
		}
	}

	return -1
}

// FindLastIndex returns the index of the last element that passes the predicate test
func FindLastIndex[T any](arr []T, predicate func(T, int, []T) bool) int {
	for i := len(arr) - 1; i >= 0; i-- {
		if predicate(arr[i], i, arr) {
			return i
		}
	}

	return -1
}

// Range returns a list of integers from start up to (but not including) stop, incremented by step
func Range(start, stop, step int) []int {
	if step == 0 {
		step = 1
	}

	// If only one argument is provided, interpret it as the stop value
	if stop == 0 && step == 1 {
		stop = start
		start = 0
	}

	// Handle negative step
	if step < 0 && start < stop {
		return []int{}
	}

	// Handle positive step
	if step > 0 && start > stop {
		return []int{}
	}

	// Calculate the number of elements
	n := int(float64(stop-start) / float64(step))
	if n <= 0 {
		return []int{}
	}

	result := make([]int, 0, n)
	for i := start; (step > 0 && i < stop) || (step < 0 && i > stop); i += step {
		result = append(result, i)
	}

	return result
}
