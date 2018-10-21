package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 1000; i++ {
		sum += i
		if sum > 1000000 {
			fmt.Println(i)
			sum = 0
		}
	}
	fmt.Println(sum)
}
