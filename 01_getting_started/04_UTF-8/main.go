package main

import "fmt"

func main() {
	for i := 60; i < 122; i++ {
		fmt.Printf("%d => %b => %o => %x => %c\n", i, i, i, i, i)
	}
}
