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
