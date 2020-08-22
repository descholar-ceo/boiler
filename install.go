package main

import (
	"fmt"
	"os/exec"
)

func main() {
	fmt.Println("Welcome to the Bo!ler cli utility, We will initialize your basic project, \nbut to do so, you will help us with few answers to the following questions.")
	fmt.Println("\n\nWhich language would you like to initialize?")
	var language string
	fmt.Scan(&language)
	if language == "ruby" {
		var projectName string
		fmt.Println("Enter the name of your project: ")
		fmt.Scan(&projectName)
		exec.Command("mkdir", projectName).Run()
	} else {
		fmt.Println("Currently we support ruby")
	}
}
