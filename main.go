package main

import (
	"DesignPatterns/builder"
	"DesignPatterns/factory"
	"log"
)

func main() {
	//Buider exemples
	me := builder.NewHuman().Named("Milena").WithAge(24).WithHeight(1.63)
	tarcisio := builder.CompleteNewHuman("Tarci", 28, 1.5)
	log.Printf("%+v", me)
	log.Printf("%+v", tarcisio)

	youngGiant := builder.Giant().WithAge(28)
	log.Printf("%+v", youngGiant)

	//Factory exemples
	animalFarmSound := factory.Farm(45)
	log.Printf("%+v", animalFarmSound)
}
