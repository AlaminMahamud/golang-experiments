package main

import (
	"fmt"
	"time"
)

func main() {
	c := time.Tick(1 * time.Nanosecond)
	for new := range c {
		fmt.Println(new)
	}

}
