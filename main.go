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

type academy struct {
	academy string
}

type tech struct {
	tech academy
}

type man struct {
	man []tech
}

type str struct {
	str [][][]man
}

func answer3() string {
	obj := str{
		str: [][][]man{
			3: {
				1: {
					2: man{
						man: []tech{
							{
								tech: academy{
									academy: "Tech Academy",
								},
							},
						},
					},
				},
			},
		},
	}

	return obj.str[3][1][2].man[0].tech.academy
}

type fruit struct {
	is string
}

type favorite struct {
	fruit fruit
}

type Person struct {
	favourite []favorite
}

func answer4() string {
	my := []Person{
		{
			favourite: []favorite{
				{},
				{},
				{},
				{fruit: fruit{
					is: "Apple",
				}},
			},
		},
	}
	return (my[0].favourite[3].fruit.is)
}

func main() {
	fmt.Println(answer1())
	fmt.Println(answer2())
	fmt.Println(answer3())
	fmt.Println(answer4())
	// fmt.Println(answer5())
}
