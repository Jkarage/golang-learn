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

