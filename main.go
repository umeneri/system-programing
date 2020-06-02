package main

import "fmt"

type Talker interface {
	Talk() int
}

type Greeter struct {
	name string
}

func (g Greeter) Talk() int {
	fmt.Printf("Hello, my name is %s\n", g.name)
	return 1
}

func main() {
	var talker Talker
	var greeter = Greeter{name: "wozozo"}
	talker = &greeter
	talker.Talk()
}
