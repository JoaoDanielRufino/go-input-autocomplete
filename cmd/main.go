package main

import (
	"fmt"
	input_autocomplete "github.com/JoaoDanielRufino/go-input-autocomplete"
)

func main() {
	path, err := input_autocomplete.Read("Path: ")

	if err != nil {
		panic(err)
	}

	fmt.Println(path)
}