# go-tuple: Generic tuples for Go 1.18.

[![Go](https://github.com/barweiss/go-tuple/actions/workflows/go.yml/badge.svg)](https://github.com/barweiss/go-tuple/actions/workflows/go.yml)

Go 1.18 tuple implementation.

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

Tuples can be used as slice or array items, map keys or values, and as channel payloads.

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
