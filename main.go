package main

import (
	"fmt"
	"lockfer/key_generator"
)

func main() {
	value := key_generator.GetKey()
	id := key_generator.CreateIdentifier()
	fmt.Printf("Valeur: %x\n", value)
	fmt.Printf("Id: %x\n", id)
}
