# Loops

The for loop  is the only loop in golang.

```golang
    for initialization; condition; post {
        // zero or more statements
        }
```

Parentheses are not used around the three components of a for loop.
The braces are mandatory, however, and the opening brace must be on the same line as the post statement.

Any of these three parts may be ommitted.

If initialization and post are ommited the traditional while loop is created.

``` golang
    for condition {
        // zero or more statements
    }

```

If all three components are omited an infinite loop is created.

``` golang
    for {
        // In  an infinte loop
    }
```

Another form of for loop iterates in a range of values in a data type like string or slices.

``` golang

for i, v := range employees {
    // Do some stuff with a employees values
}

```
