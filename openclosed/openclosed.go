package openclosed

import "fmt"

// software entities should be open for extension, but closed
// for modification

type A struct {
	year int
}

func (receiver A) Greet() {
	fmt.Println("Hello GolangUK", receiver.year)
}

type B struct {
	A
}

func (receiver B) Greet() {
	fmt.Println("Welcome to GolangUK", receiver.year)
}

func TIP() {
	var a A
	a.year = 2016

	var b B
	b.year = 2017

	a.Greet()
	b.Greet()
}

type Cat struct {
	Name string
}

func (cat Cat) Legs() int {
	return 4
}

func (cat Cat) PrintLegs() {
	fmt.Printf("I have %d legs\n", cat.Legs())
}

type OctoCat struct {
	Cat
}

func (octoCat OctoCat) Legs() int {
	return 5
}

func CatTip() {
	var octo OctoCat
	fmt.Println(octo.Legs())
	octo.PrintLegs()
}
