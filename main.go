package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to the Bo!ler cli utility, We will initialize your basic project, \nbut to do so, you will help us with few answers to the following questions.")
	fmt.Println("\n\nWhat is the project name you want to use?")
	var projectName string
	fmt.Scan(&projectName)
	fmt.Println("Choose a number which correspond to the language you will be using:\n1.Ruby")
	var language int
	fmt.Scan(&language)
	if language != 1 {
		for i := 0; i < 5; i++ {
			fmt.Println("Choose a number which correspond to the language you will be using:\n1.Ruby")
			fmt.Scan(&language)
			if language == 1 {
				break
			}
		}
		fmt.Println("The language you chose is not supported")
	}
}
