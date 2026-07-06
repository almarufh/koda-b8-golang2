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

func answer1() string {
	we := c{
		are: b{
			the: a{
				best: "Koda",
			},
		},
	}
	return we.are.the.best
}

type world struct {
	world string
}

func answer2() string {
	hello := world{
		world: "Hello World",
	}
	return hello.world
}

func main() {
	fmt.Println(answer1())
	fmt.Println(answer2())
}
