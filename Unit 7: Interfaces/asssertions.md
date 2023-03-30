# Introduction

Type assertion is an operation applied to an interface value.
The syntax is like this x.(T), where x is an expression of an interface type and T is a type called the asserted type.
A type assertion checks that the dynamic type of its operand matches the asserted type.

There are two possibilities:
- First when the asserted type T is a concrete type
- Second when the asserted type is an interface as well

An example:

``` golang 
	var w io.Writer
	w = os.Stdout
	f := w.(*os.File) // Success: f == os.Stdout
	c := w.(*bytes.Buffer) // panic: interface holds *os.File not *bytes.Buffer
```

If the asserted type is an interface, then the type assertion checks wheather x's dynamic type satisifies T.

An example: 

``` golang
	var w io.Writer
	w = os.Stdout
	rw := w.(io.ReadWriter) // success as *os.file has both read and write Methods
	w = new(ByteCounter)
	w.(io.ReadWriter) // panic as *ByteCounter has no Read Method
```

No matter what type was asserted, if the operand is a nil interface value, Thet type assertion fails.
An example: 

``` golang
	var w io.ReadWriter 
	rw := w.(io.Writer) // Fails since w is nil
```

## Testing Dynamic Value of an Interface

Often we are not sure of the dynamic type of the interface value, we'd like to test whether it is some particular type.If the type assertions appears in an assignment in which two results are expected, such as the following declarations, the operation does not panic on failure but6 instead returns a second result; a boolean indicating success.

An example:
``` golang
	var w io.Writer = os.Stdout
	f, ok := w.(*os.File) // Success ok, f == os.Stdout
	b, ok := w.(*os.File) // !ok, b == nil
```

if the operation fails the second result is false and the first result is equal to the zero value of the asserted type.
Another example:
``` golang
	if f,ok := w.(io.ReadWriter); ok {
		// use of f...
	}
```

## A word of Advice
Interfaces are only needed when there are two or more concrete types that must be dealt with
in a uniform way.
