package main

import "fmt"

func main() {
	fmt.Println("Welcome to the Bo!ler cli utility, We will initialize your basic project, \nbut to do so, you will help us with few answers to the following questions.")
	fmt.Println("\n\nWhich language would you like to initialize?")
	var language string
	fmt.Scan(&language)
	fmt.Printf("You enetered %v", language)
}
