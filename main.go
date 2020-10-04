package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	ruby "github.com/descholar-ceo/boiler/languages"

	"github.com/descholar-ceo/boiler/utils"
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
	workingDir = utils.AskWorkingDirectory()

	// project name
	fmt.Println("\n\nWhat is the project name you want to use?")
	fmt.Scan(&projectName)

	// choose a language
	fmt.Println("\nChoose a number which correspond to the language or framework you will be using:\n1.Ruby\n2.Ruby on Rails (RoR)")
	fmt.Scan(&language)

	// different language boilers
	switch language {
	case 1: //ruby is chosen
		ruby.RubyBoiler()
	case 2:
		rorBoiler()
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
	fmt.Println("Make sure that rails and ruby are installed correctly on your system, and is working well, if rails is not installed yet, then this time your project initialization will take some time")
	fmt.Println("If ruby is not installed yet, please refer to this link for a proper ruby installation: https://www.theodinproject.com/courses/ruby-programming/lessons/installing-ruby-ruby-programming")
	fmt.Println("\n\nChecking Rails installation on your computer...")
	mStr := "gem install rails"
	argsMStr := strings.Split(mStr, " ")
	exec.Command(argsMStr[0], argsMStr[1:]...).Run()

	isGithub := utils.AskGithub()
	isRubocop := utils.AskRubocop()
	database := utils.AskDatabase()

	// moving to the project dir
	if workingDir == "." {
		wrkDr, _ = os.Getwd()
	} else {
		if utils.IsDirectoryExists(workingDir) {
			wrkDr = utils.GetHomeDirectory() + "/" + workingDir
		} else {
			fmt.Printf("\n%s does not exist, your ruby on rails project will be created in the current directory\n", workingDir)
			wrkDr, _ = os.Getwd()
		}
	}

	// create a project with rails
	fmt.Println("\nChecking out your working directory")
	os.Chdir(wrkDr)

	fmt.Println("\nGenerating your Rails project using Rails installed on your machine, This might take several minutes depending on the internet connection you have, please bear with us, and wait...")
	railsStr := "rails new " + projectName + " --database=" + strings.Trim(database, "\"")
	args := strings.Split(railsStr, " ")
	exec.Command(args[0], args[1:]...).Run()

	fmt.Println("\nChecking out your project workspace...")
	os.Chdir(projectName)

	fmt.Println("\nTemplating your README file")
	utils.Copy(utils.GetHomeDirectory()+"/.boiler/boiler/lib/.defaults/README.md", "README.md")

	if isGithub == "y" {
		fmt.Println("\nSetting up your github directory...")
		os.Mkdir(".github", 0755)
		os.Mkdir(".github/workflows", 0755)
		utils.Copy(utils.GetHomeDirectory()+"/.boiler/boiler/lib/.defaults/.github/PULL_REQUEST_TEMPLATE.md", ".github/PULL_REQUEST_TEMPLATE.md")
	}

	if isRubocop == "y" {
		fmt.Println("\nCreating Rubocop YAML file...")
		utils.Copy(utils.GetHomeDirectory()+"/.boiler/boiler/lib/.ror/.rubocop.yml", ".rubocop.yml")

		if isGithub == "y" {
			utils.Copy(utils.GetHomeDirectory()+"/.boiler/boiler/lib/.ror/.github/workflows/linters.yml", ".github/workflows/linters.yml")
		}
		fmt.Println("\nCreating the stylelint file for your stylelings...")
		utils.Copy(utils.GetHomeDirectory()+"/.boiler/boiler/lib/.ror/.stylelintrc.json", ".stylelintrc.json")

		fmt.Println("\nInstalling custom linter dependecies...")
		stylelintStr := "yarn add --dev stylelint stylelint-scss stylelint-config-standard"
		styleArgs := strings.Split(stylelintStr, " ")
		exec.Command(styleArgs[0], styleArgs[1:]...).Run()

		fmt.Println("\nThe linters (Rubocop and stylelint) have been installed successfully!")
		fmt.Println("To use Rubocop for checking errors: rubocop")
		fmt.Println("To use Stylelint for checking errors: npx stylelint \"**/*.{css,scss}\"")
	}

}
