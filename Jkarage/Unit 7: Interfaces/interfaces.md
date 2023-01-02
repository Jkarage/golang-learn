# Introduction

Interface is an abstract type. It doesn't expose its representation or structure of internal values, or the set of operations they support. It reveals only some of their methods.

An interface type specifies a set of methods that a concrete type must posses to be considered an instance of that interface.

The assignabilty rule for interface is very simple. An expression may be assigned to an interface only if its type satisfies the interface.

var w io.Writer
w = os.Stdout // Fine as *os.File has write method.
w = time.Second // Not Ok as time.Second has no write method.

