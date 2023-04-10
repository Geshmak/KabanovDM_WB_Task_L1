package main

import "fmt"

type Human struct {
	name      string
	age       int
	education bool
}

type Action struct {
	Human    //embedding
	location string
}

func NewHuman(nname string, nage int, neducation bool) Human {
	return Human{nname, nage, neducation}
}
func NewAction(nname string, nage int, neducation bool, nlocation string) Action {
	return Action{Human{nname, nage, neducation}, nlocation}
}
func (h Human) Introduce() {
	fmt.Printf("Hi my name is %s I am %d and it is %t that i have an education\n", h.name, h.age, h.education)
}
func main() {
	nh := NewHuman("Dkab", 18, false)
	fmt.Println(nh)
	na := NewAction("aDkab", 22, true, "Moscow")
	fmt.Println(na)

	nh.Introduce()
	na.Introduce()
}
