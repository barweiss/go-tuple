# go-tuple: Generic tuples for Go 1.18+.

[![Go](https://github.com/barweiss/go-tuple/actions/workflows/go.yml/badge.svg)](https://github.com/barweiss/go-tuple/actions/workflows/go.yml)
[![Coverage Status](https://coveralls.io/repos/github/barweiss/go-tuple/badge.svg)](https://coveralls.io/github/barweiss/go-tuple)
[![Go Report Card](https://goreportcard.com/badge/github.com/barweiss/go-tuple)](https://goreportcard.com/report/github.com/barweiss/go-tuple)
[![Go Reference](https://pkg.go.dev/badge/github.com/barweiss/go-tuple.svg)](https://pkg.go.dev/github.com/barweiss/go-tuple)
[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go)

Go 1.18+ tuple implementation.

Use tuples to store 1 or more values without needing to write a custom struct.

```go
tup := tuple.New2(5, "hi!")
fmt.Println(tup.V1) // Outputs 5.
fmt.Println(tup.V2) // Outputs "hi!".
```

Tuples come in various sizes, from 1 to 9 elements.

```go
longerTuple := tuple.New5("this", "is", "one", "long", "tuple")
```

Tuples can be used as slice or array items, map keys or values, and as channel payloads:

```go
// Map holding tuples.
tupInMap := make(map[tuple.T2[string, string]]Person)
tupInMap[tuple.New2("John", "Doe")] = Person{
	FirstName: "John",
	LastName: "Doe",
	// ...
}

// Channel holding tuples.
tupInChan := make(chan tuple.T2[string, error])
go func() {
	defer close(tupInChan)
	tupInChan <- tuple.New2(os.Getwd())
}()
fmt.Print(<-tupInChan)
```

# Features

## Create tuples from function calls

```go
func vals() (int, string) {
    return 5, "hi!"
}

func main() {
    tup := tuple.New2(vals())
    fmt.Println(tup.V1)
    fmt.Println(tup.V2)
}
```

## Forward tuples as function arguments

```go
func main() {
    tup := tuple.New2(5, "hi!")
    printValues(tup.Values())
}

func printValues(a int, b string) {
    fmt.Println(a)
    fmt.Println(b)
}
```

## Access tuple values

```go
tup := tuple.New2(5, "hi!")
a, b := tup.Values()
```

## JSON Marshalling

Tuples are marshalled and unmarshalled as JSON arrays.

```go
type User struct {
	Name string `json:"name"`
	Age  int    `json:"age,omitempty"`
}

type MyJSON struct {
	Users []tuple.T2[string, User] `json:"users"`
}

func main() {
	data := MyJSON{
		Users: []tuple.T2[string, User]{
			tuple.New2("foo", User{Name: "foo", Age: 42}),
			tuple.New2("bar", User{Name: "bar", Age: 21}),
			tuple.New2("baz", User{Name: "baz"}),
		},
	}

	marshalled, _ := json.Marshal(data)
	fmt.Printf("%s\n", string(marshalled))
	// Outputs: {"users":[["foo",{"name":"foo","age":42}],["bar",{"name":"bar","age":21}],["baz",{"name":"baz"}]]}
}
```

## Comparison

Tuples are compared from the first element to the last.
For example, the tuple `[1 2 3]` is greater than `[1 2 4]` but less than `[2 2 2]`.

```go
fmt.Println(tuple.Equal3(tuple.New3(1, 2, 3), tuple.New3(3, 3, 3))) // false.
fmt.Println(tuple.LessThan3(tuple.New3(1, 2, 3), tuple.New3(3, 2, 1))) // true.

tups := []tuple.T3{
    tuple.New3("foo", 2, -23),
    tuple.New3("foo", 72, 15),
    tuple.New3("bar", -4, 43),
}
sort.Slice(tups, func (i, j int) {
    return tuple.LessThan3(tups[i], tups[j])
})

fmt.Println(tups) // [["bar", -4, 43], ["foo", 2, -23], ["foo", 72, 15]].
```

---

**NOTE**

In order to compare tuples, all tuple elements must match `constraints.Ordered`.

See [Custom comparison](#custom-comparison) in order to see how to compare tuples
with arbitrary element values.

---

### Comparison result

```go
// Compare* functions return an OrderedComparisonResult value.
result := tuple.Compare3(tuple.New3(1, 2, 3), tuple.New3(3, 2, 1))

// OrderedComparisonResult values are wrapped integers.
fmt.Println(result) // -1

// OrderedComparisonResult expose various method to see the result
// in a more readable way.
fmt.Println(result.GreaterThan()) // false
```

### Custom comparison

The package provides the `CompareC` comparison functions varation in order to compare tuples of complex
comparable types.

For a type to be comparable, it must match the `Comparable` or `Equalable` constraints.

```go
type Comparable[T any] interface {
	CompareTo(guest T) OrderedComparisonResult
}

type Equalable[T any] interface {
	Equal(guest T) bool
}
```

```go
type person struct {
	name string
	age  int
}

func (p person) CompareTo(guest person) tuple.OrderedComparisonResult {
	if p.name < guest.name {
		return -1
	}
	if p.name > guest.name {
		return 1
	}
	return 0
}

func main() {
	tup1 := tuple.New2(person{name: "foo", age: 20}, person{name: "bar", age: 24})
	tup2 := tuple.New2(person{name: "bar", age: 20}, person{name: "baz", age: 24})

	fmt.Println(tuple.LessThan2C(tup1, tup2)) // true.
}
```

In order to call the complex types variation of the comparable functions, **all** tuple types must match the `Comparable` constraint.

While this is not ideal, this a known inconvenience given the current type parameters capabilities in Go.
Some solutions have been porposed for this issue ([lesser](https://github.com/lelysses/lesser), for example, beatifully articulates the issue),
but they still demand features that are not yet implemented by the language.

Once the language will introduce more convenient ways for generic comparisons, this package will adopt it.

## Formatting

Tuples implement the `Stringer` and `GoStringer` interfaces.

```go
fmt.Printf("%s\n", tuple.New2("hello", "world"))
// Output:
// ["hello" "world"]

fmt.Printf("%#v\n", tuple.New2("hello", "world"))
// Output:
// tuple.T2[string, string]{V1: "hello", V2: "world"}
```

# Notes

The tuple code and test code are generated by the `scripts/gen/main.go` script.

Generation works by reading `tuple.tpl` and `tuple_test.tpl` using Go's `text/template` engine.
`tuple.tpl` and `tuple_test.tpl` contain the templated content of a generic tuple class, with variable number of elements.

# Contributing

Please feel free to contribute to this project by opening issues or creating pull-requests.
However, keep in mind that generic type features for Go are still in their early stages, so there might
not be support from the language to your requested feature.

Also keep in mind when contributing to keep the compilation time and performance of this package fast.

Feel free to contact me at [barw.code@gmail.com](mailto:barw.code@gmail.com) for questions or suggestions!
