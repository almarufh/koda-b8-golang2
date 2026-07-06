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

func answer() string {
	we := c{
		are: b{
			the: a{
				best: "Koda",
			},
		},
	}
	return (we.are.the.best)
}

func main() {
	fmt.Println(answer())
}
