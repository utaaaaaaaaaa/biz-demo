package main

import "fmt"

type gasEngine struct {
	mpg       int
	gallon    int
	ownerInfo owner
}

func (engine gasEngine) milesLeft() int {
	return engine.mpg * engine.gallon
}

type owner struct {
	name string
}

func main() {

	var engine gasEngine = gasEngine{4, 4, owner{"zoro"}}
	fmt.Println(engine)
	fmt.Println(engine.milesLeft())
}
