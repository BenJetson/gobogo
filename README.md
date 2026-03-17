# gobogo

[![Go Report Card](https://goreportcard.com/badge/github.com/BenJetson/gobogo)](https://goreportcard.com/report/github.com/BenJetson/gobogo)
[![Go Reference](https://pkg.go.dev/badge/github.com/BenJetson/gobogo.svg)](https://pkg.go.dev/github.com/BenJetson/gobogo)
[![Test](https://github.com/BenJetson/gobogo/actions/workflows/test.yml/badge.svg)](https://github.com/BenJetson/gobogo/actions/workflows/test.yml)

This package provides an implementation of
[bogosort](https://en.wikipedia.org/wiki/Bogosort) for the
[Go programming language](https://go.dev) that usees the
[generics feature introduced in Go 1.18](https://go.dev/blog/intro-generics).

## Installation

```sh
go get -u github.com/BenJetson/gobogo
```

## Example

```go
package main

import (
	"fmt"

	bogo "github.com/BenJetson/gobogo"
)

func main() {
	s := []int{5, 1, 9, 8, 3, 6, 2, 4, 7}
	fmt.Println("input:", s) // unsorted.
	bogo.Sort(s)
	fmt.Println("output:", s) // (eventually) sorted!
}
```

You can [try it on the Go Playground](https://go.dev/play/p/5kRsU2Y0-br)!

## Usage

I wouldn't recommend using this package, but
[documentation](https://pkg.go.dev/github.com/BenJetson/gobogo) is available if
you would like to do so.

All usage is subject to the [terms of the MIT License](./LICENSE).
