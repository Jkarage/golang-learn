# Introduction

Generics in Go are available since version 1.18, With generics you can create function with types as parameters instead of writing functions with specific

You can now write a function with a type parameter

``` golang
    func Last[T any] (s []) {
        return s[len(s)-1]
    }
```

Type parameters are declared using square brackets, They describe types that are allowed in a given function.

## Constraints

A constraint is an interface that describes a type parameter. Only types that satisfy the interface can be used as parameter in a generic function. The constraints always appears in a square brackets after the type parameter name.

any is an alias for interface{}

### Builtin Constraints

Apart from `any` another builtin constraint is `comparable` a constraints that describes any type whose values can be compared.

As of Go 1.18 an interface has a new syntax, Now it is possible to define an interface with a type.

``` golang
    type Integer interface {
        int
    }
```

Using the union operator `|` we can define an interface with more than one type.

``` golang
    type Number interface {
        int | float64
    }
```

When you see `~` special character before a type, it means the constraints allows this type as well as the other types whose underline type is as this defined in the constraints.

``` golang
    type Int interface {
        ~int | int8 | int16 | int32
    }
```
