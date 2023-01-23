# Introduction

Interface is an abstract type. It doesn't expose its representation or structure of internal values, or the set of operations they support. It reveals only some of their methods.

An interface type specifies a set of methods that a concrete type must posses to be considered an instance of that interface.

``` golang
// A writer interface specifies that for a concrete type to be a writer interface it musty posses a write method with a given signature.
    type Writer interface {
        Write(p []byte) (len(p), error)
    }
```

The assignabilty rule for interface is very simple. An expression may be assigned to an interface only if its type satisfies the interface.

``` golang
    var w io.Writer
    w = os.Stdout // Fine as *os.File has write method.
    w = time.Second // Not Ok as time.Second has no write method.
```

A different notion about go interfaces is that they are satisfied implicitly.
This is a major concern that differentiate interfaces in golang and other programming languages.

Two most common interfaces in golang is Writer and Reader interfaces.

Interfaces can be embedded to generate more complex interfaces like structs embedding.

```golang
    type ReadWriter interface {
        Reader
        Writer
    }
```

The order in which methods appear is immaterial. All that matters is the set of methods.

An interface with more methods gives places much demands on the type that implements it.
Now what about an interface without any method.

```golang
    type any interface {}

```

The empty interface places no demands on the type that implements it. thus we can assign any value to the empty interface.

``` golang
    type any interface {}
    any = true
    any = 12.34
    any = "hello"
    any = new(bytes.Buffer())
```

But now we can do nothing directly to the value it holds since an interface has no methods.We will need a way to get back the value from the interface. Using type Assertion.

## Interface Values

Conceptually a value of an interface has two components a concrete type and a value of that type. These  values are called interfaces dynamic types and dynamic value.

In Go, variables are always initialized to a well defined value. There is no exception to interfaces.
The zero value for an interface has both its types and value components set to nil.

``` golang

    var w io.Writer // dynamic type is set to nil and dynamic value to nil
    
    w = os.Stdout   // dynamic type is *os.File and dynamic value a pointer to stdout descriptor
    
    w = new(bytes.Buffer) // dynamic type is *bytes.Buffer and dynamic value a pointer to the newly allocated buffer
    
    w = nil // sets both its components to nil

```

An interface value is described as nil or non-nil based on it's dynamic type.

An interface value can hold large dynamic values for an instance `time.Time` which is an instance in time struct having few unexported fields.

Interface values may be comparable and because of this they may be used as keys in a map.

However if two interface values are compared and their dynamic type is not comparable a panic would result.

``` golang
    var x interface{} = []int{1, 3, 4}
    fmt.Println(x == x) // Panic due to comparing slices
```

In this aspect basic types are safe for comparison and other types like functions, slices and maps are not safe for comparison in an interface.

Thus when comparing interface values or aggregate types we must be aware of the potential for a panic.

You may check the dynamic type of an interface using the `%T` formmatter.

``` golang
    resp, _ := http.Get("http:///google.com/")
    fmt.Println("%T",resp.Body)
```

Internally fmt uses Reflection to know the dynamic type of the interface.

## Caveat: An Interface containing a nil Pointer is Non-nil

A nil interface value containing no value at all is not equal to an interface having a pointer that happens to be nil.
