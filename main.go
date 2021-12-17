package main

import "fmt"

type Creature struct {
	Species string
}

func main() {

	var creature string = "shark"
	var pointer *string = &creature

	fmt.Println("creature =", creature)
	fmt.Println("pointer =", pointer)
	fmt.Println("*pointer =", *pointer)
	*pointer = "jellyfish"
	fmt.Println("*pointer =", *pointer)
	fmt.Println("creature =", creature)

	var creature2 Creature = Creature{Species: "Shark"}

	fmt.Printf("1) %+v\n", creature2)
	changeCreature_PassByValue(creature2)
	fmt.Printf("3) %+v\n", creature2)
	changeCreature_PassByReference(&creature2)
	fmt.Printf("5) %+v\n", creature2)

	nilPointer()
}

func changeCreature_PassByValue(creature Creature) {
	creature.Species = "jellyfish"
	fmt.Printf("2) %+v\n", creature)
}

func changeCreature_PassByReference(creature *Creature) {
	if creature == nil {
		fmt.Println("creature is nil")
		return
	}
	creature.Species = "Dolphin"
	fmt.Printf("4) %+v\n", creature)
}

func nilPointer() {
	var creature *Creature
	fmt.Printf("1) %+v\n", creature)
	changeCreature_PassByReference(creature)
	fmt.Printf("3) %+v\n", creature)
}
