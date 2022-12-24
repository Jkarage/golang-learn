# introduction

Go is a compiled language just like c, c++ and Java.

The gotoolchain `go` converts the source and its dependencies into instructions in the native machine language of a computer.

Go natively handles unicode so it can process texts in all the world languages.

Go code is organized into packages which are similar to libraries or modules in other languages.

Package main is special, it defines a standalone executable program, not a library. Within package main function main is also special it is where execution of the program begins.

The imported packages should be exactly what the program requires, no unused imports or missing imports are allowed.

Go doesn't require semicolons at the end of statements(go takes care of it),except when two or more appear on the same line.

`gofmt` rewrites code into standard format. In some code editors you can configure gofmt to run every time a file is saved.

`goimports` is another tool that manages the insertion and removal of import declarations as needed.

Functions and other package-level entities may be declared in any order.

Naming in go functions, variables, constants, types and packages follow a simple rule.

A name begins with a letter(unicode letter) or an underscore and may be foollwed by a number of letters, digits or underscores.

case matters `jose` and `Jose` are two different names.

predeclared names like `new` can be redeclared in your program as they are not keywords. But beware of the confusion.

Stylistically go programmers use camel case naming convention.
