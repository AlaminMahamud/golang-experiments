package main

import (
	"fmt"
	"time"
)

func expensiveCall() { time.Sleep(2 * time.Second) }

func main() {
	t0 := time.Now()
	expensiveCall()
	t1 := time.Now()
	fmt.Println("The call took %v to run", t1.Sub(t0))
}
