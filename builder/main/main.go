package main

import "fmt"

func main() {
	me := newHuman().named("Milena").withAge(24).withHeight(1.63)
	tarcisio := newHuman().named("Tarci").withAge(28).withHeight(1.5)
	fmt.Printf("%+v\n", me)
	fmt.Printf("%+v\n", tarcisio)

	youngGiant := giant().withAge(28)
	fmt.Printf("%+v\n", youngGiant)

}

type human struct {
	name   string
	age    int
	height float64
}

// Empty constructor
func newHuman() human {
	return human{}
}

// Full constructors allow you to pass all properties once
func completeNewHuman(
	name string,
	age int,
	height float64) human {
	return human{
		name:   name,
		age:    age,
		height: height,
	}
}

// Builders
func (h human) named(name string) human {
	h.name = name
	return h
}

func (h human) withAge(age int) human {
	h.age = age
	return h
}

func (h human) withHeight(height float64) human {
	h.height = height
	return h
}

// Pre-defined builders: when you know you're gonna build frequently the same set of properties
func giant() human {
	return newHuman().named("Onias").withHeight(2.0)
}
