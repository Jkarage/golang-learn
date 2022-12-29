# Introduction

A function lets us wrap a sequence of statements together as a unit that can be called fron elsewhere in a program.

## Function Declarations

func name(parameter-list) (result-list) {
    body
}

The parameter-list specifies names and types of function parameters.
The result-list specifies the types of the values that the function returns.

Two functions have the same types or signature if they have same sequence of parameters types and same sequence of result types.

Go has no concept of default parameter values, nor anyway to specify arguments by names, so the name of paramters and result doesn't matter to the caller except for the documentation.

Arguments are passed by value, so the function receives a copy of the argument and changes to the argument value has no effect to the original value.

However if the arguments contains a reference types like a pointer, slice,maps, function or channel then the caller may be affected by any modification done to the passed value indirectly.

If you encouneter a function declaration without a body, It is an indication that a function implementation is done in other language other than go.

func Sin(x float64) float64 // implemented in assembly language

## Recursive Functions

Recursive functions call themselves directly or indirectly.

## Function Values

Functions are first class values in go. Like other values function has types and may be assigned to variables and be passed or returned in a function.

var f func(n int) int

A function  value may be called like any other function.

f(3)

The zero value of a function type is nil. Calling a nil function value causes a panic.

Function values may be compared with nil but not with other function values. Function values are not comparable. Thus function values may not be used as map keys.

## Anonnymous Functions

Named functions can be declared only at the package level. But we can use a function literal to denote a function value within an expresion.

A function literal is written like a function declaration, but without a func name.

Functions declared in this manner have access to the entire lexical block from the outer function declaration.

