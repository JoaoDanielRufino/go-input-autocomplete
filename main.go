package main

import (
	"./input_autocomplete"
)

func main() {
	_, err := input_autocomplete.Read("Path: ")

	if err != nil {
		panic(err)
	}
}
