# Data Types

Go's types fall into four types basic types,aggregate types,reference types and interface types.

Basic types includes string, numbers and booleans.

Aggregate types includes Arrays and structs.

Reference types is a diverse group that includes  pointers, slices, map, functions and channels.What they have in common is that they all refer to a program variable or state  indirectly such that an operation applied to the variable is seen by all copies of that reference.

interface awaits us in unit 10.

## Basic Types

All values of basic types are comparable, meaning the two values of the same data type can be compared using == and !=.

### Numbers

Numbers include integers which is divided into signed(int) and unsigned ints(uint) and further divided into int8, int16, int32 and int64 for signed and uint8, uint16, uint32 and uint64 for unsigned ints.

You cant compare different integers types as they are different data types, The compiler will catch the mismatch type error.Any operation wouldn't work unless they have been converted to a similar data type.

Numbers also include floating point numbers and complex numbers.
Floating point numbers are in two sizes float32 and float64.

Go also provides two sizes of complex number complex32 and complex64.

### Booleans

A value of type bool has only two possible values `true` and `false`.
There is no implicit conversion from bool to numeric value like 0 or 1

### Strings

Strings are an immutable sequence of bytes.(A string is an `array` of bytes)

The len() function returns the number of bytes (not runes) in a string.
Attempting to access a byte outside a range results in a panic.

The i-th byte of a string may not necessarilly be the i-th character of a string because non-ascii characters may take more than two bytes.

The substring s[i:j] creates a new string consisting of characters from i upto j(not inclusive).

strings may be compared with with comparison operators like == and <;The comparison is done byte by byte.

Escape sequences can be used in double quoted strings only.

## Aggregate Types

These includes arrays and structs.

### Array

Array is a fixed length sequence of zero  or more elements of a particular type.

By defaults elements of a new array variable are set to zero value of the array elements type.

The sizeof an array is part of its type, so [3]int [4]int are different types.

If an array type is comparable then the array is comparable.

The fixed size property of an array, makes it less used in go, The slice data type overtaking the position.

Passing an array to a function passes it by a copy, Thus no modification can be attained. Unless the pointer to array is passed.

### Slices

Slice is a variable-length sequence whose elements are of the same data type.

Array and slices are intimately connected. A slice gives access to a subsequence (or perhaps all) of the elements of an array which is known as slice's underlying array.

A slice has three properties: a `pointer`,length and capacity.

Due to that passing a slice to a function, assures the function to modify the underlying array, Unlike in Arrays.

Slices are not comparable. Although the standard library provides a highly optimized bytes.Equal function for comparing [] bytes. But for other slices.

The only legal comparison in slice is to nil. To check if a slice is empty use len() function as there are non-nil slices that are empty.

append() deals with size realocation of slices.

### Maps

An unordered collection of key, value pairs. All of the keys in the given map are of the same type, and values are of the same type. But the key and values can be of different types.

The Key type of map must be comparable, so that it a map can test if a key given is equal to one already within it.

Map element is not a variable and can not be addressed.

Main reason we can not take the adress of map element is growing the map may cause rehashing of existing elements to another storage locations thus potential invalidating the address.

The zero value of a map is nil, That is a reference to no hash table at all.

### Structs

An aggregate type that groups together zero or more named values of arbitrary types as a single entity. Each value is called a `field`

Field order is significant to type identity.

The name of a struct type is exported if it starts wiith a Capital letter.
A struct may contain a mixture of exported and non exported fields.

If all fields of a struct are comparable structs are comparable.

comparable structs like any other comparable types can be used as keys in maps.
