# go-hypem [![GoDoc](https://godoc.org/github.com/9uuso/go-hypem?status.svg)](https://godoc.org/github.com/9uuso/go-hypem)
Library for converting HypeMachine.com's mediaid's into streamable links.

## Installation

`go get github.com/9uuso/go-hypem`

## Example

```
package main

import (
	"fmt"

	"github/9uuso/go-hypem"
)

func main() {
	url, err := hypem.Stream("2a59t")
	if err != nil {
		panic(err)
	}

	// http://api.soundcloud.com/tracks/196680458/stream?consumer_key=nH8p0jYOkoVEZgJukRlG6w
	fmt.Println(url)
}
```
