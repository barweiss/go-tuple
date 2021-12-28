# go-tuple: Generic tuples for Go 1.18.

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
