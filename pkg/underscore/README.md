# Underscore.go

A Go implementation of [Underscore.js](https://underscorejs.org/), the popular utility library and some useful tools from world. This library provides functional programming helpers without extending any built-in objects, offering utility functions for working with collections, arrays, objects, functions, and more.

## Features

The library is organized into several packages:

- **collections**: Functions for working with collections (arrays, slices, maps)
- **arrays**: Functions for working with arrays and slices
- **functions**: Functions for working with functions (binding, throttling, debouncing, etc.)
- **objects**: Functions for working with objects (maps and structs)
- **utility**: Utility functions (random, template, etc.)

## Installation

```bash
go get git.snappfood.ir/backend/go/packages/sf-underscore
```

## Usage

Here's a quick overview of the library:

```go
import (
    "fmt"
    
    "git.snappfood.ir/backend/go/packages/sf-underscore/arrays"
    "git.snappfood.ir/backend/go/packages/sf-underscore/collections"
    "git.snappfood.ir/backend/go/packages/sf-underscore/functions"
    "git.snappfood.ir/backend/go/packages/sf-underscore/objects"
    "git.snappfood.ir/backend/go/packages/sf-underscore/utility"
)
```

## Collections

### each
Iterates over a list of elements, yielding each in turn to an iteratee function.

```go
people := []Person{
    {Name: "John", Age: 30},
    {Name: "Jane", Age: 25},
}

collections.Each(people, func(person Person, index int, _ interface{}) bool {
    fmt.Printf("%d: %s is %d years old\n", index, person.Name, person.Age)
    return true // continue iteration
})
// Output:
// 0: John is 30 years old
// 1: Jane is 25 years old
```

### map
Produces a new array of values by mapping each value in list through a transformation function.

```go
names := collections.Map(people, func(person Person, _ int, _ interface{}) string {
    return person.Name
})
fmt.Println(names) // [John Jane]
```

### reduce
Boils down a list of values into a single value by applying an iteratee function.

```go
totalAge := collections.Reduce(people, func(result int, person Person, _ int, _ interface{}) int {
    return result + person.Age
}, 0)
fmt.Println(totalAge) // 55
```

### reduceRight
Like reduce, but processes the list from right to left.

```go
reversedNames := collections.ReduceRight(people, func(result string, person Person, _ int, _ interface{}) string {
    if result == "" {
        return person.Name
    }
    return result + ", " + person.Name
}, "")
fmt.Println(reversedNames) // Jane, John
```

### find
Returns the first element in the list that passes the predicate test.

```go
jane, found := collections.Find(people, func(person Person, _ int, _ interface{}) bool {
    return person.Name == "Jane"
})
if found {
    fmt.Println(jane.Name, jane.Age) // Jane 25
}
```

### filter
Returns all elements in the list that pass the predicate test.

```go
adults := collections.Filter(people, func(person Person, _ int, _ interface{}) bool {
    return person.Age >= 30
})
fmt.Println(adults) // [{John 30}]
```

### where
Looks through each value in the list, returning an array of all the values that match the properties.

```go
type User struct {
    Name     string
    Role     string
    Active   bool
}

users := []User{
    {Name: "John", Role: "admin", Active: true},
    {Name: "Jane", Role: "admin", Active: false},
    {Name: "Bob", Role: "user", Active: true},
}

admins := collections.Where(users, map[string]interface{}{
    "Role": "admin",
})
fmt.Println(admins) // [{John admin true} {Jane admin false}]
```

### findWhere
Returns the first value in the list that matches the properties.

```go
activeAdmin, found := collections.FindWhere(users, map[string]interface{}{
    "Role": "admin",
    "Active": true,
})
if found {
    fmt.Println(activeAdmin.Name) // John
}
```

### reject
Returns the values in list without the elements that pass the predicate test.

```go
nonAdults := collections.Reject(people, func(person Person, _ int, _ interface{}) bool {
    return person.Age >= 30
})
fmt.Println(nonAdults) // [{Jane 25}]
```

### every
Returns true if all elements in the list pass the predicate test.

```go
allActive := collections.Every(users, func(user User, _ int, _ interface{}) bool {
    return user.Active
})
fmt.Println(allActive) // false
```

### some
Returns true if at least one element in the list passes the predicate test.

```go
anyAdmin := collections.Some(users, func(user User, _ int, _ interface{}) bool {
    return user.Role == "admin"
})
fmt.Println(anyAdmin) // true
```

### contains
Returns true if the list contains the specified value.

```go
names := []string{"John", "Jane", "Bob"}
hasJane := collections.Contains(names, "Jane")
fmt.Println(hasJane) // true
```

### invoke
Calls the method named by methodName on each value in the list.

```go
type Counter struct {
    value int
}

func (c *Counter) Increment() int {
    c.value++
    return c.value
}

counters := []*Counter{{value: 0}, {value: 10}}
results := collections.Invoke(counters, "Increment")
fmt.Println(results) // [1 11]
```

### pluck
Extracts a list of property values from a list of objects.

```go
ages := collections.Pluck(people, "Age")
fmt.Println(ages) // [30 25]
```

### max
Returns the maximum value in list, based on the result of the provided iteratee function.

```go
oldest, found := collections.Max(people, func(person Person) float64 {
    return float64(person.Age)
})
if found {
    fmt.Println(oldest.Name) // John
}
```

### min
Returns the minimum value in list, based on the result of the provided iteratee function.

```go
youngest, found := collections.Min(people, func(person Person) float64 {
    return float64(person.Age)
})
if found {
    fmt.Println(youngest.Name) // Jane
}
```

### sortBy
Returns a sorted copy of list, ranked in ascending order by the results of running iteratee on each value.

```go
sortedByAge := collections.SortBy(people, func(person Person) float64 {
    return float64(person.Age)
})
fmt.Println(sortedByAge) // [{Jane 25} {John 30}]
```

### groupBy
Groups the list's values by the result of the iteratee function.

```go
groupedByRole := collections.GroupBy(users, func(user User) string {
    return user.Role
})
fmt.Println(groupedByRole) // map[admin:[{John admin true} {Jane admin false}] user:[{Bob user true}]]
```

### indexBy
Given a list and an iteratee function that returns a key for each element in the list, returns an object with the values indexed by the keys.

```go
indexedByName := collections.IndexBy(users, func(user User) string {
    return user.Name
})
fmt.Println(indexedByName["John"]) // {John admin true}
```

### countBy
Sorts a list into groups and returns a count for the number of objects in each group.

```go
countByRole := collections.CountBy(users, func(user User) string {
    return user.Role
})
fmt.Println(countByRole) // map[admin:2 user:1]
```

### shuffle
Returns a shuffled copy of the list.

```go
shuffled := collections.Shuffle([]int{1, 2, 3, 4, 5})
fmt.Println(shuffled) // e.g., [3, 1, 5, 2, 4]
```

### sample
Returns n random elements from the list.

```go
twoRandomUsers := collections.Sample(users, 2)
fmt.Println(twoRandomUsers) // e.g., [{Bob user true} {John admin true}]
```

### toArray
Converts the list to a Go slice.

```go
usersMap := map[string]User{
    "john": {Name: "John", Role: "admin", Active: true},
    "jane": {Name: "Jane", Role: "admin", Active: false},
}
usersSlice := collections.ToArray(usersMap)
fmt.Println(usersSlice) // [{John admin true} {Jane admin false}]
```

### size
Returns the number of values in the list.

```go
size := collections.Size(users)
fmt.Println(size) // 3
```

### partition
Splits the collection into two arrays: one whose elements pass the predicate test and one whose elements do not.

```go
active, inactive := collections.Partition(users, func(user User, _ int, _ interface{}) bool {
    return user.Active
})[0], collections.Partition(users, func(user User, _ int, _ interface{}) bool {
    return user.Active
})[1]
fmt.Println(active)   // [{John admin true} {Bob user true}]
fmt.Println(inactive) // [{Jane admin false}]
```

## Arrays

### first
Returns the first n elements of an array.

```go
numbers := []int{1, 2, 3, 4, 5}
first := arrays.First(numbers)
fmt.Println(first) // [1]

firstThree := arrays.First(numbers, 3)
fmt.Println(firstThree) // [1, 2, 3]
```

### initial
Returns everything but the last n elements of an array.

```go
initial := arrays.Initial(numbers)
fmt.Println(initial) // [1, 2, 3, 4]

allButLastTwo := arrays.Initial(numbers, 2)
fmt.Println(allButLastTwo) // [1, 2, 3]
```

### last
Returns the last n elements of an array.

```go
last := arrays.Last(numbers)
fmt.Println(last) // [5]

lastThree := arrays.Last(numbers, 3)
fmt.Println(lastThree) // [3, 4, 5]
```

### rest
Returns everything but the first n elements of an array.

```go
rest := arrays.Rest(numbers)
fmt.Println(rest) // [2, 3, 4, 5]

allButFirstTwo := arrays.Rest(numbers, 2)
fmt.Println(allButFirstTwo) // [3, 4, 5]
```

### compact
Returns a copy of the array with all falsy values removed.

```go
mixedNumbers := []int{0, 1, 2, 0, 3, 0, 4}
compacted := arrays.Compact(mixedNumbers)
fmt.Println(compacted) // [1, 2, 3, 4]
```

### flatten
Flattens a nested array.

```go
nested := []interface{}{1, []interface{}{2, []interface{}{3, 4}}, 5}
flattened := arrays.Flatten(nested)
fmt.Println(flattened) // [1, 2, 3, 4, 5]

shallowFlattened := arrays.Flatten(nested, true)
fmt.Println(shallowFlattened) // [1, 2, [3, 4], 5]
```

### without
Returns a copy of the array with all instances of the specified values removed.

```go
withoutEvenNumbers := arrays.Without(numbers, 2, 4)
fmt.Println(withoutEvenNumbers) // [1, 3, 5]
```

### union
Returns the union of the arrays; all the unique items from all of the passed-in arrays.

```go
moreNumbers := []int{4, 5, 6, 7}
union := arrays.Union(numbers, moreNumbers)
fmt.Println(union) // [1, 2, 3, 4, 5, 6, 7]
```

### intersection
Returns the intersection of the arrays; all the items present in all of the passed-in arrays.

```go
intersection := arrays.Intersection(numbers, moreNumbers)
fmt.Println(intersection) // [4, 5]
```

### difference
Returns the values from array that are not present in other arrays.

```go
difference := arrays.Difference(numbers, moreNumbers)
fmt.Println(difference) // [1, 2, 3]
```

### uniq
Returns a duplicate-free version of the array.

```go
duplicates := []int{1, 2, 2, 3, 3, 3, 4, 5, 5}
unique := arrays.Uniq(duplicates)
fmt.Println(unique) // [1, 2, 3, 4, 5]
```

### zip
Merges together the values of each of the arrays with the values at the corresponding position.

```go
names := []string{"John", "Jane", "Bob"}
ages := []int{30, 25, 40}
zipped := arrays.Zip(names, ages)
fmt.Println(zipped) // [[John 30] [Jane 25] [Bob 40]]
```

### unzip
The opposite of zip. Separates zipped arrays into their individual component arrays.

```go
unzipped := arrays.Unzip(zipped)
fmt.Println(unzipped) // [[John Jane Bob] [30 25 40]]
```

### chunk
Splits an array into groups of specified size.

```go
chunks := arrays.Chunk(numbers, 2)
fmt.Println(chunks) // [[1 2] [3 4] [5]]
```

### indexOf
Returns the index at which the value can be found in the array, or -1 if not found.

```go
index := arrays.IndexOf(names, "Jane")
fmt.Println(index) // 1
```

### lastIndexOf
Returns the index of the last occurrence of the value in the array, or -1 if not found.

```go
repeatedNames := []string{"John", "Jane", "John", "Bob"}
lastIndex := arrays.LastIndexOf(repeatedNames, "John")
fmt.Println(lastIndex) // 2
```

### sortedIndex
Returns the index at which a value should be inserted into a sorted array to maintain order.

```go
sortedAges := []int{20, 30, 40, 50}
insertIndex := arrays.SortedIndex(sortedAges, 35, func(age int) float64 {
    return float64(age)
})
fmt.Println(insertIndex) // 2
```

### findIndex
Returns the index of the first element that passes the predicate test.

```go
userIndex := arrays.FindIndex(users, func(user User, _ int, _ []User) bool {
    return user.Role == "admin" && user.Active
})
fmt.Println(userIndex) // 0
```

### findLastIndex
Returns the index of the last element that passes the predicate test.

```go
lastUserIndex := arrays.FindLastIndex(users, func(user User, _ int, _ []User) bool {
    return user.Role == "admin"
})
fmt.Println(lastUserIndex) // 1
```

### range
Generates a list of integers, starting with start, incrementing by step, and stopping before end.

```go
rangeOneToTen := arrays.Range(1, 11, 1)
fmt.Println(rangeOneToTen) // [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]

rangeTwoToTwentyByTwo := arrays.Range(2, 21, 2)
fmt.Println(rangeTwoToTwentyByTwo) // [2, 4, 6, 8, 10, 12, 14, 16, 18, 20]
```

## Functions

### bind
Creates a function that invokes func with the this binding of thisArg and prepends arguments.

```go
type Calculator struct {
    Multiplier int
}

func (c *Calculator) Multiply(n int) int {
    return c.Multiplier * n
}

calc := &Calculator{Multiplier: 5}
timesTwo := functions.Bind(calc.Multiply, calc, 2)
fmt.Println(timesTwo()) // 10
```

### bindAll
Binds methods of an object to the object itself.

```go
type Counter struct {
    Count int
    Increment func()
}

counter := &Counter{Count: 0}
increment := func() {
    counter.Count++
}

functions.BindAll(counter, "Increment")
counter.Increment()
fmt.Println(counter.Count) // 1
```

### partial
Partially applies arguments to a function.

```go
add := func(a, b, c int) int {
    return a + b + c
}

addFive := functions.Partial(add, 5)
fmt.Println(addFive(10, 15)) // 30
```

### memoize
Creates a function that memoizes the result of func.

```go
fibonacci := func(n int) int {
    if n <= 1 {
        return n
    }
    return fibonacci(n-1) + fibonacci(n-2)
}

hasher := func(args ...interface{}) string {
    return fmt.Sprintf("%d", args[0])
}

memoizedFib := functions.Memoize(fibonacci, hasher)
fmt.Println(memoizedFib(40)) // Fast calculation due to memoization
```

### delay
Delays the execution of func until after wait milliseconds have elapsed.

```go
functions.Delay(func(message string) {
    fmt.Println(message)
}, 1000, "This message appears after 1 second")
time.Sleep(1500 * time.Millisecond) // Wait for the delayed function to execute
```

### defer
Defers executing the function until the current call stack has cleared.

```go
functions.Defer(func() {
    fmt.Println("This will be executed after the current function completes")
})
```

### throttle
Creates a function that, when invoked repeatedly, will only actually call the original function at most once per every wait milliseconds.

```go
counter := 0
increment := func() {
    counter++
    fmt.Printf("Counter: %d\n", counter)
}

throttledIncrement := functions.Throttle(increment, 500)

// Call repeatedly
for i := 0; i < 5; i++ {
    throttledIncrement()
    time.Sleep(100 * time.Millisecond)
}
// Will only execute once or twice due to throttling
```

### debounce
Creates a function that will delay invoking func until after wait milliseconds have elapsed since the last time it was invoked.

```go
counter := 0
increment := func() {
    counter++
    fmt.Printf("Counter: %d\n", counter)
}

debouncedIncrement := functions.Debounce(increment, 500)

// Call repeatedly
for i := 0; i < 5; i++ {
    debouncedIncrement()
    time.Sleep(100 * time.Millisecond)
}
time.Sleep(600 * time.Millisecond)
// Counter will only be incremented once
```

### once
Creates a function that is restricted to invoking func once.

```go
counter := 0
incrementOnce := functions.Once(func() {
    counter++
})

incrementOnce()
incrementOnce()
incrementOnce()
fmt.Println(counter) // 1
```

### after
Creates a function that invokes func once it's called n or more times.

```go
counter := 0
incrementAfterThreeCalls := functions.After(3, func() {
    counter++
})

incrementAfterThreeCalls()
incrementAfterThreeCalls()
incrementAfterThreeCalls() // First time counter will be incremented
incrementAfterThreeCalls() // Second time counter will be incremented
fmt.Println(counter) // 2
```

### before
Creates a function that invokes func, with the this binding and arguments of the created function, while it's called less than n times.

```go
counter := 0
incrementBeforeThreeCalls := functions.Before(3, func() {
    counter++
})

incrementBeforeThreeCalls() // First time counter will be incremented
incrementBeforeThreeCalls() // Second time counter will be incremented
incrementBeforeThreeCalls() // Third time counter will NOT be incremented
fmt.Println(counter) // 2
```

### wrap
Creates a function that provides value to the wrapper function as its first argument.

```go
hello := func(name string) string {
    return "Hello, " + name
}

greet := functions.Wrap("John", hello)
fmt.Println(greet()) // "Hello, John"
```

### negate
Creates a function that negates the result of the predicate func.

```go
isEven := func(n int) bool {
    return n%2 == 0
}

isOdd := functions.Negate(isEven)
fmt.Println(isOdd(3)) // true
```

### compose
Creates a function that is the composition of the provided functions.

```go
doubleIt := func(n int) int {
    return n * 2
}

addOne := func(n int) int {
    return n + 1
}

doubleAndAddOne := functions.Compose(addOne, doubleIt)
fmt.Println(doubleAndAddOne(5)) // 11 (5 * 2 + 1)
```

## Objects

### keys
Retrieves all the keys of an object.

```go
person := map[string]interface{}{
    "name": "John",
    "age": 30,
    "city": "New York",
}

keys := objects.Keys(person)
fmt.Println(keys) // [name age city]
```

### allKeys
Retrieves all the keys of an object, including inherited ones.

```go
// In Go, this is similar to Keys for maps
allKeys := objects.AllKeys(person)
fmt.Println(allKeys) // [name age city]
```

### values
Retrieves all the values of an object.

```go
values := objects.Values(person)
fmt.Println(values) // [John 30 New York]
```

### mapObject
Maps an object to a new object with the same keys.

```go
uppercased := objects.MapObject(person, func(value interface{}, key string, _ interface{}) interface{} {
    if str, ok := value.(string); ok {
        return strings.ToUpper(str)
    }
    return value
})
fmt.Println(uppercased) // map[name:JOHN age:30 city:NEW YORK]
```

### pairs
Converts an object into a list of [key, value] pairs.

```go
pairs := objects.Pairs(person)
fmt.Println(pairs) // [[name John] [age 30] [city New York]]
```

### invert
Creates an object with the keys and values swapped.

```go
inverted := objects.Invert(map[string]string{
    "first": "John",
    "last": "Doe",
})
fmt.Println(inverted) // map[John:first Doe:last]
```

### create
Creates a new object with the given prototype.

```go
prototype := map[string]interface{}{
    "species": "human",
}

person := objects.Create(prototype, map[string]interface{}{
    "name": "John",
    "age": 30,
})
fmt.Println(person) // map[species:human name:John age:30]
```

### functions
Returns a sorted list of all function names available on the object.

```go
type API struct{}

func (api *API) Get() {}
func (api *API) Post() {}
func (api *API) Delete() {}

methods := objects.Functions(&API{})
fmt.Println(methods) // [Delete Get Post]
```

### findKey
Returns the first key where the predicate returns true.

```go
key, found := objects.FindKey(person, func(value interface{}, key string, _ interface{}) bool {
    if str, ok := value.(string); ok {
        return strings.Contains(str, "John")
    }
    return false
})
fmt.Println(key, found) // name true
```

### extend
Copies all properties from the source objects to the destination object.

```go
user := map[string]interface{}{
    "name": "John",
}

details := map[string]interface{}{
    "age": 30,
    "city": "New York",
}

objects.Extend(&user, details)
fmt.Println(user) // map[name:John age:30 city:New York]
```

### pick
Returns a copy of the object with only the whitelisted properties.

```go
nameAndAge := objects.Pick(person, "name", "age")
fmt.Println(nameAndAge) // map[name:John age:30]
```

### omit
Returns a copy of the object without the blacklisted properties.

```go
withoutAge := objects.Omit(person, "age")
fmt.Println(withoutAge) // map[name:John city:New York]
```

### defaults
Fills in undefined properties in object with values from the defaults objects.

```go
user := map[string]interface{}{
    "name": "John",
}

defaultValues := map[string]interface{}{
    "name": "Anonymous",
    "age": 25,
    "city": "Unknown",
}

objects.Defaults(&user, defaultValues)
fmt.Println(user) // map[name:John age:25 city:Unknown]
```

### clone
Creates a shallow copy of the object.

```go
userCopy := objects.Clone(user)
fmt.Println(userCopy) // map[name:John age:25 city:Unknown]
```

### isEqual
Performs a deep comparison between two objects.

```go
equal := objects.IsEqual(user, userCopy)
fmt.Println(equal) // true
```

## Utility

### noConflict
Returns a reference to the Underscore object.

```go
underscore := utility.NoConflict()
fmt.Println(underscore) // "underscore"
```

### identity
Returns the same value that is used as the argument.

```go
value := utility.Identity(42)
fmt.Println(value) // 42
```

### constant
Returns a function that returns value.

```go
answer := utility.Constant(42)
fmt.Println(answer()) // 42
```

### noop
Does nothing and returns nothing.

```go
utility.Noop() // Does nothing
```

### times
Invokes the iteratee n times, returning an array of the results.

```go
squares := utility.Times(5, func(n int) int {
    return n * n
})
fmt.Println(squares) // [0, 1, 4, 9, 16]
```

### random
Returns a random integer between min and max (inclusive).

```go
randomNum := utility.Random(1, 10)
fmt.Println(randomNum) // Random integer between 1 and 10
```

### uniqueId
Returns a unique ID with an optional prefix.

```go
id1 := utility.UniqueId()
id2 := utility.UniqueId("user_")
fmt.Println(id1) // e.g., "1627484945854"
fmt.Println(id2) // e.g., "user_1627484945855"
```

### escape
Escapes HTML special characters in a string.

```go
escaped := utility.Escape("<script>alert('XSS')</script>")
fmt.Println(escaped) // &lt;script&gt;alert(&#39;XSS&#39;)&lt;/script&gt;
```

### unescape
Unescapes HTML special characters in a string.

```go
unescaped := utility.Unescape("&lt;script&gt;alert(&#39;XSS&#39;)&lt;/script&gt;")
fmt.Println(unescaped) // <script>alert('XSS')</script>
```

### template
Compiles a template into a function that can be evaluated.

```go
compiled := utility.Template("Hello <%= name %>!")
result := compiled(map[string]interface{}{"name": "World"})
fmt.Println(result) // "Hello World!"
```

## License

MIT