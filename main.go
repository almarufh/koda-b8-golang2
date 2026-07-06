package main

import "fmt"

type a struct {
	best string
}

type b struct {
	the a
}

type c struct {
	are b
}

func main() {
	we := c{
		are: b{
			the: a{
				best: "Koda",
			},
		},
	}
	fmt.Println(we.are.the.best)
}
