package main

import (
	"fmt"

	"github.com/descholar-ceo/boiler/utils"

	"github.com/descholar-ceo/boiler/languages"
)

var (
	projectName string
	language    int
)

func main() {
	fmt.Println("\nWelcome to the Bo!ler cli utility, We will initialize your basic project, \nbut to do so, you will help us with few answers to the following questions.")

	// choose a language
	fmt.Println("\nChoose a number which correspond to the language or framework you will be using:\n1.Ruby\n2.Ruby on Rails (RoR)")
	fmt.Scan(&language)

	// different language boilers
	switch language {
	case 1: //ruby is chosen
		languages.RubyBoiler()
	case 2:
		languages.RorBoiler()
	default: // the chosen language is not yet supported
		for i := 0; i < 5; i++ {
			fmt.Println("\nChoose a number which correspond to the language you will be using:\n1.Ruby\n2.Ruby on Rails")
			fmt.Scan(&language)
			if language == 1 {
				break
			}
		}
		language = 0
		fmt.Println("\nThe language you chose is not supported")
		return
	}

	utils.DisplayLastCommands()
}
