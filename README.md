# libraidan
A general-purpose library from R. Aidan Campbell

[![Go Report Card](https://goreportcard.com/badge/github.com/raidancampbell/libraidan)](https://goreportcard.com/report/github.com/raidancampbell/libraidan)
[![Build Status](https://travis-ci.com/raidancampbell/libraidan.svg?branch=master)](https://travis-ci.com/raidancampbell/libraidan)

**Caveat Emptor**: Use this library at your own risk.  This is merely a collection of functions and abstractions that I've found useful

## Structure & Usage
This library prefixes each package name with `r` to prevent namespace collisions.  Currently this is a `pkg`-only library, meaning it's meant to be imported and used in `golang` code.  There is no command-line interface

Import the desired package
```
import "github.com/raidancampbell/libraidan/pkg/rstrings"
```

## Packages
#### [~~rcollections~~](https://godoc.org/github.com/raidancampbell/libraidan/pkg/rcollections) 

**Deprecated, use a package like [GoDS](https://github.com/emirpasic/gods) instead**

The ~~rcollections~~ package implements several common map and set operations, such as `Contains` or `GetWithDefault`, as well as the more advanced `map`, `reduce`, and `filter` operations.  These are provided one for each type, because I don't know enough golang to write it cleaner.  PRs welcome.

#### [`rmath`](https://godoc.org/github.com/raidancampbell/libraidan/pkg/rmath)

The `rmath` package contains math reducing operators, such as `sum`, `min`, and `max`, etc.

#### [`roper`](https://godoc.org/github.com/raidancampbell/libraidan/pkg/roper)

The `roper` package implements more advanced "operator-style" functions.  The main influence for this is Python's coalescing `or` operator: the first non-falsy operand is returned. 

#### [`rstrings`](https://godoc.org/github.com/raidancampbell/libraidan/pkg/rstrings)

The `rstrings` package implements more niche or abbreviated string functions.

#### [`rruntime`](https://godoc.org/github.com/raidancampbell/libraidan/pkg/rruntime)

The `rruntime` package contains useful runtime or meta functions.
