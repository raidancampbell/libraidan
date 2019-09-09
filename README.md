# libraidan
A general-purpose library from R. Aidan Campbell

**Caveat Emptor**: Use this library at your own risk.  This is merely a collection of functions and abstractions that I've found useful

## Structure & Usage
This library prefixes each package name with `r` to prevent namespace collisions.  Currently this is a `pkg`-only library, meaning it's meant to be imported and used in `golang` code.  There is no command-line interface

## Packages
#### `rcollections` 

The `rcollections` package implements several common map and set operations, such as `Contains` or `GetWithDefault`, as well as the more advanced `map`, `reduce`, and `filter` operations.  These are provided one for each type, because I don't know enough golang to write it cleaner.  PRs welcome.

#### `roper`

The `roper` package implements more advanced "operator-style" functions.  The main influence for this is Python's truthy `or` operator: the first non-falsy operand is returned. 

#### `rstrings`

The `rstrings` package implements more niche or abbreviated string functions.