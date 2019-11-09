# libraidan
A general-purpose library from R. Aidan Campbell

[![Go Report Card](https://goreportcard.com/badge/github.com/raidancampbell/libraidan)](https://goreportcard.com/report/github.com/raidancampbell/libraidan)
[![Build Status](https://travis-ci.com/raidancampbell/libraidan.svg?branch=master)](https://travis-ci.com/raidancampbell/libraidan)

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

#### `rruntime`

The `rruntime` package contains useful runtime or meta functions.

## TODO:

 - [X] lint code
 - [X] implement CI linting with cool badges
 - [ ] revisit `rcollections` after reviewing golang's type system
 - [X] add gitignore and license files
 - [ ] tag initial release
 - [ ] implement benchmark tests for `rcollections`
 - [ ] implement `int` and `float64` maps
 - [ ] add `map` / `reduce` / `filter` function implementation constants