# Introduction

A method is a function associated with a particular type.

## Method Declarations

A method is declared with a variant of a function declaration with an extra parameter before the function name. The parameter attaches the function to that type.

The extra parameter is called methods receiver, a alegacy from old object oriented languages that described calling methods as sending a message to an object.

In Go there is no special name like this or self for the receiver.

Methods may be defined in any named types as long as the underline type is not a pointer or an interface.

## Methods with a pointer receiver

Because calling a function makes a copy of arguments values.If we want to modify a variable or a  variable is so large and we dont want to copy it, we must pass the address of the variable using a pointer.

The same goes to the method that wants to update the receiver variable.

For convention if a method uses pointer receiver, Then all other methods should use pointer receiver even ones that dont strictly need it.

Named types and pointers to them are the only types that may appear in a
receiver declaration. Furthermore, to avoid ambiguities, method declarations are not permitted on named types that are themselves pointer types.

``` golang
    type P *int
    func (P) f() { body } // compile error: invalid receiver type
```

For simplicity in Go even if a method has apointer receiver, We can just use the normal variable to access it, and go will use the pointer implicitly.
But the variable must be addressable.

```golang
package main

type intSet struct{}

func main() {
intSet{}.String() // Cannot do this (Not addressable)
}

func (*intSet) String() string

```

We can not use a type literal to access a method, Because there is no way to access the address of the type literal.

Point{1, 3}.scaleBy(2) // compile error: can't take address of Point literal.
