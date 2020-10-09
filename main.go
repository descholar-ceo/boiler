package main

import (
	"fmt"

	"github.com/descholar-ceo/boiler/languages"
)

var (
	tmpWrkDr      string
	workingDir    string
	wrkDr         string
	err           error
	projectName   string
	language      int
	isRubocop     string
	isTests       string
	testFramework int
	isGithub      string
	database      string
)

func main() {
	fmt.Println("\nWelcome to the Bo!ler cli utility, We will initialize your basic project, \nbut to do so, you will help us with few answers to the following questions.")
	// working dir
	// workingDir = utils.AskWorkingDirectory()

	// project name
	// projectName = utils.AskProjectName()

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
			fmt.Println("\nChoose a number which correspond to the language you will be using:\n1.Ruby")
			fmt.Scan(&language)
			if language == 1 {
				break
			}
		}
		language = 0
		fmt.Println("\nThe language you chose is not supported")
		return
	}

	// Displaying last commands
	fmt.Println("\n\nYour project has been initialized successfully")
	fmt.Println("The remaining task is to go on github and create a repository and copy its url")
	fmt.Printf("Come back in the root directory of %s\n", projectName)
	fmt.Println("\nRun the following commands respectifuly")
	fmt.Println("1. git remote add .")
	fmt.Println("2. git commit -m \"Initial commit\"")
	fmt.Println("3. git remote add origin [Paste the url you copied from github]")
	fmt.Println("4. git push -u origin master")
	fmt.Print("\n\nCongratulations and good luck for your new project\n\n\n")
}

// rorBoiler
func rorBoiler() {

	// isGithub := utils.AskGithub()
	// isRubocop := utils.AskRubocop()
	// database := utils.AskDatabase()

}
