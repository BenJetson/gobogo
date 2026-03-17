# gobogo

[![Go Report Card](https://goreportcard.com/badge/github.com/BenJetson/gobogo)](https://goreportcard.com/report/github.com/BenJetson/gobogo)
[![Go Reference](https://pkg.go.dev/badge/github.com/BenJetson/gobogo.svg)](https://pkg.go.dev/github.com/BenJetson/gobogo)

This package provides an implementation of
[bogosort](https://en.wikipedia.org/wiki/Bogosort) for the
[Go programming language](https://go.dev).

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
