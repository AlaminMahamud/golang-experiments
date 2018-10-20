package main

import "fmt"

func pass(ptr *string) {
	*ptr = "Mahamud"
	ptr = nil
}

func main1() {
	s := "Alamin"
	refOfs := &s
	pass(refOfs)

	fmt.Printf("%T => %v, %v\n", s, s, &s)
	fmt.Printf("%T => %v", refOfs, refOfs)
}

type Count int

// func (c Count) Incr() int {
// 	return int(c + 1)
// }

func main2() {
	var c Count
	c.Incr()
	fmt.Println(c)
	c.Incr()
	fmt.Println(c)
}

func main3() {
	var c Count = Count(1)
	fmt.Printf("%T => %v\n", c, c)
}

func (c *Count) Incr() int {
	*c = *c + 1
	return int(*c)
}

func main() {
	c := Count(5)
	c.Incr()
	fmt.Println(c)
	c.Incr()
	fmt.Println(c)
}

type Counter interface {
	Incr() int
}

func onApiHit(c Counter) {
	c.Incr()
}
