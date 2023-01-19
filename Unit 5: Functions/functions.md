# Introduction

A function lets us wrap a sequence of statements together as a unit that can be called fron elsewhere in a program.

## Function Declarations

``` golang
    func name(parameter-list) (result-list) {
        // body
    }
```

The parameter-list specifies names and types of function parameters.
The result-list specifies the types of the values that the function returns.

Two functions have the same types or signature if they have same sequence of parameters types and same sequence of result types.

Go has no concept of default parameter values, nor anyway to specify arguments by names, so the name of paramters and result doesn't matter to the caller except for the documentation.

Arguments are passed by value, so the function receives a copy of the argument and changes to the argument value has no effect to the original value.

However if the arguments contains a reference types like a pointer, slice,maps, function or channel then the caller may be affected by any modification done to the passed value indirectly.

If you encouneter a function declaration without a body, It is an indication that a function implementation is done in other language other than go.

``` golang
    func Sin(x float64) float64 // implemented in assembly language
```

## Recursive Functions

Recursive functions call themselves directly or indirectly.

## Function Values

Functions are first class values in go. Like other values function has types and may be assigned to variables and be passed or returned in a function.

``` golang
    var f func(n int) int
```

A function  value may be called like any other function.

``` golang
f(3)
```

The zero value of a function type is nil. Calling a nil function value causes a panic.

Function values may be compared with nil but not with other function values. Function values are not comparable. Thus function values may not be used as map keys.

## Anonnymous Functions

Named functions can be declared only at the package level. But we can use a function literal to denote a function value within an expresion.

A function literal is written like a function declaration, but without a func name.

Functions declared in this manner have access to the entire lexical block from the outer function declaration.

## Capturing Iterating variables

A closure has accesss to the outer lexical environment, Thus all variables accessed from it are passed as reference (not as value at a particular time) this leads to a bug that if a variable changes during the runtime the closure accessing it might get a value that was not intended to.

## Variadic Functions

Variadic function is the one that can be called with a varying number of arguments.

To declare variadic function the type of the final parameter must be preceded by an elipsis "..."

That indicates that a function may be called with a number of arguments for this type.

Implicitly the caller allocates an array to store the arguments, and passes the slice of the entire array to the function.

Invoking a variadic function when the arguments are already in a slice.

``` golang
    values := []{10, 20, 30}
    func (values...)
```

## Defered Function Calls

Syntactically, a defer statement is an ordinary function or method call prefixed by the keyword defer.

The function and argument expressions are evaluated when the statement is executed, but the actual call is deferred until the function that contains the defer statement has finished, whether normally, by executing a return statement or falling off the end, or abnormally, by panicking.
