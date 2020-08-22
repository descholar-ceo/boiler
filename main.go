package main

import (
	"fmt"
)

var (
	projectName   string
	language      int
	isRubocop     string
	isTests       string
	testFramework int
	isGithub      string
)

func main() {
	fmt.Println("Welcome to the Bo!ler cli utility, We will initialize your basic project, \nbut to do so, you will help us with few answers to the following questions.")

	// project name
	fmt.Println("\n\nWhat is the project name you want to use?")
	// var projectName string
	fmt.Scan(&projectName)

	// choose a language
	fmt.Println("Choose a number which correspond to the language you will be using:\n1.Ruby")
	// var language int
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
		return
	}
	fmt.Println("Will you use Rubocop as a linter? Enter y for yes or any other key for no")
	// var isRubocop string
	fmt.Scan(&isRubocop)

	// will you run tests?
	fmt.Println("Will you write some unit tests for your project? Enter y for yes or any other key for no")
	// var isTests string
	fmt.Scan(&isTests)
	if isTests == "y" || isTests == "Y" {
		fmt.Println("Choose a number which corresponds to the testing framework you will be using:\n1.RSpec")
		// var testFramework int
		fmt.Scan(&testFramework)
		if testFramework != 1 {
			for i := 0; i < 5; i++ {
				fmt.Println("Choose a number which corresponds to the testing framework you will be using:\n1.RSpec")
				fmt.Scan(&testFramework)
				if testFramework == 1 {
					break
				}
			}
			fmt.Println("The the testing framework you chose is not supported")
		}
	}

	// will you use github?
	fmt.Println("Will you use github as a collaboration tool? Enter y for yes or any other key for no")
	// var isGithub string
	fmt.Scan(&isGithub)

	fmt.Println("\n\n\nThe following are your preferences, we will setup your project depending on your preferences.")
	fmt.Printf("Project name : %v\n", projectName)
	fmt.Printf("Language name : %v\n", language)
	fmt.Printf("Will you write unit test? : %v\n", isTests)
	fmt.Printf("Testing framework : %v\n", testFramework)
	fmt.Printf("Will you use github? : %v\n", isGithub)

}
