package main

//yanzera

import (
	"fmt"

	"./input_autocomplete"
)

func main() {
	path, err := input_autocomplete.Read("Path: ")

	if err != nil {
		panic(err)
	}

	fmt.Println(path)
}
