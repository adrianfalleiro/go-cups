# go-cups 

> Experimental Golang bindings for CUPS

## Info

This is an experimental library that provides limited bindings to CUPS. This is my first attempt at using CGO.

Currently only working on OS X and printing test pages.

## Usage

```shell
go get https://github.com/adrianfalleiro/go-cups
```

## Example

```go
package main

import (
    "fmt"

    cups "github.com/adrianfalleiro/go-cups"
)

func main() {
    printers := cups.NewDefaultConnection()
    for _, dest := range printers.Dests {
        fmt.Printf("%v (%v)\n", dest.Name, dest.Status())
    }

}
```
