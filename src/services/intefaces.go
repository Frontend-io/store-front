package objects

import (
	"fmt"
	"log"
)

type (
	Animal interface {
		Says() string
		NumberOfLegs() int
	}

	Dog struct {
		Name  string
		Color string
		Age   int
	}

	Goat struct {
		Name    string
		Color   string
		Sex     string
		HasHorn bool
	}
)

func AnimalInteface() {
	newDog := Dog{
		Name:  "bingo",
		Color: "Brown",
		Age:   1,
	}

	newAnimal := PrintInfo(&newDog)
	log.Print(newAnimal)
}

func PrintInfo(s Animal) string {
	return fmt.Sprintf("This animals's says %s", s.Says())
}

func (d *Dog) Says() string {
	return "Woof"
}

func (d *Dog) NumberOfLegs() int {
	return 4
}
