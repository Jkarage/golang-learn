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

But now we can do nothing directly to the value it holds since an interface has no methods. 