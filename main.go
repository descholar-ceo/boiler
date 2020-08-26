package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/mitchellh/go-homedir"
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
)

func main() {
	fmt.Println("Welcome to the Bo!ler cli utility, We will initialize your basic project, \nbut to do so, you will help us with few answers to the following questions.")

	// working directory
	fmt.Println("Enter the working directory:")
	fmt.Scan(&workingDir)

	// project name
	fmt.Println("\n\nWhat is the project name you want to use?")
	fmt.Scan(&projectName)

	// choose a language
	fmt.Println("Choose a number which correspond to the language or framework you will be using:\n1.Ruby\n2.Ruby on Rails (RoR)")
	fmt.Scan(&language)

	// different language boilers
	switch language {
	case 1: //ruby is chosen
		rubyBoiler()
	case 2:
		rorBoiler()
	default: // the chosen language is not yet supported
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

func askRubocop() {
	fmt.Println("\nWill you use Rubocop as a linter? Enter y for yes or any other key for no")
	fmt.Scan(&isRubocop)
}
func askGithub() {
	fmt.Println("\nWill you use github as a collaboration tool? Enter y for yes or any other key for no")
	fmt.Scan(&isGithub)
}

// rubyBoiler function
func rubyBoiler() {

	// will you use rubocop?
	askRubocop()

	// will you run tests?
	fmt.Println("\nWill you write some unit tests for your project? Enter y for yes or any other key for no")
	fmt.Scan(&isTests)
	if isTests == "y" || isTests == "Y" {
		fmt.Println("\nChoose a number which corresponds to the testing framework you will be using:\n1.RSpec")
		fmt.Scan(&testFramework)
		if testFramework != 1 {
			for i := 0; i < 5; i++ {
				fmt.Println("\nChoose a number which corresponds to the testing framework you will be using:\n1.RSpec")
				fmt.Scan(&testFramework)
				if testFramework == 1 {
					break
				}
			}
			fmt.Println("\nThe testing framework you chose is not supported")
		}
	}

	// will you use github?
	askGithub()

	fmt.Println("\n\n\nThe following are your preferences, we will setup your project depending on your preferences.")
	fmt.Printf("\n\nWorking dir : %v\n", workingDir)
	fmt.Printf("Project name : %v\n", projectName)
	fmt.Printf("Language name : %v\n", language)
	fmt.Printf("Will you write unit test? : %v\n", isTests)
	fmt.Printf("Testing framework : %v\n", testFramework)
	fmt.Printf("Will you use github? : %v\n\n", isGithub)

	// create project dir
	createProjectDirectory()

	// initialize rubocop
	createRubocopFile()

	if isGithub == "y" {
		// initialize github actions
		createGithubActionsDirectory()
		copy("./lib/.ruby/.github/workflows/linters.yml", wrkDr+"/.github/workflows/linters.yml")
		copy("./lib/.ruby/.github/workflows/tests.yml", wrkDr+"/.github/workflows/tests.yml")
	}

	// create initial files
	fmt.Printf("\nstep 06/15 => Creating lib folder in %s directory...\n", projectName)
	os.Mkdir(wrkDr+"/lib", 0755)

	fmt.Printf("\nstep 07/15 => Creating bin folder in %s directory...\n", projectName)
	os.Mkdir(wrkDr+"/bin", 0755)

	fmt.Printf("\nstep 08/15 => Adding .gitkeep file in %s/lib directory...\n", projectName)
	os.Create(wrkDr + "/lib/.gitkeep")

	fmt.Printf("\nstep 09/15 => Creating main.rb file in %s directory...\n", projectName)
	os.Create(wrkDr + "/bin/main.rb")
	writeToFile(wrkDr+"/bin/main.rb", "puts 'Hello from Boiler! Welcome to the new world!'")

	// change working dir
	fmt.Println("\nstep 10/15 => Checking out your project dir...")
	os.Chdir(wrkDr)

	// initialize gemfile
	fmt.Printf("\nstep 11/15 => Initializing gem in %s directory...\n", projectName)
	exec.Command("bundle", "init").Run()

	if isTests == "y" {
		// initialize rspec
		fmt.Printf("\nstep 12/15 => Initializing rspec in %s directory...\n", projectName)
		writeToFile("Gemfile", "gem 'rspec', '~>3.0'")
		exec.Command("rspec", "--init").Run()
	}

	if isRubocop == "y" {
		// install rubocop in gems
		fmt.Println("\nstep 13/15 => Writing gems...")
		writeToFile("Gemfile", "gem 'rubocop', '~>0.81.0'")
	}

	// initialize git
	fmt.Printf("\nstep 14/15 => Initializing git in %s directory...\n", projectName)
	exec.Command("git", "init").Run()

	// installing bundler gems
	fmt.Printf("\nstep 15/15 => Installing gems %s directory, this might take some minutes, please wait...\n", projectName)
	exec.Command("bundle", "install").Run()
}

// rorBoiler
func rorBoiler() {

	askGithub()
	askRubocop()
}

func writeToFile(file, stringToWrite string) {
	mFile, _ := os.OpenFile(file, os.O_APPEND|os.O_WRONLY, 0644)
	fmt.Fprintln(mFile, stringToWrite)
	mFile.Close()
}

func isDirectoryExists(directory string) bool {
	_, err := os.Stat(getHomeDirectory() + "/" + directory)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func getHomeDirectory() string {
	homeDirectory, _ := homedir.Dir()
	return homeDirectory
}

func createProjectDirectory() {
	if workingDir == "." {
		// create project in current directory
		tmpWrkDr, _ = os.Getwd()
		wrkDr = tmpWrkDr + "/" + projectName
	} else {
		// checking if a directory exists
		if isDirectoryExists(workingDir) {
			wrkDr = getHomeDirectory() + "/" + workingDir + "/" + projectName
		} else {
			fmt.Println("The directory you entered does not exists, your project will be created in the current directory")
			tmpWrkDr, _ = os.Getwd()
			wrkDr = tmpWrkDr + "/" + projectName
		}
	}

	// create a project directory
	fmt.Printf("\nstep 01/15 => Creating directory to %s...\n", projectName)
	os.Mkdir(wrkDr, 0755)
}

func createGithubActionsDirectory() {
	if isGithub == "y" {
		fmt.Printf("\nstep 03/15 => Initializing github actions in %s directory...\n", projectName)
		os.Mkdir(wrkDr+"/.github", 0755)
		os.Mkdir(wrkDr+"/.github/workflows", 0755)

		// create a readme file
		fmt.Printf("\nstep 04/15 => Creating README file in %s directory...\n", projectName)
		copy("./lib/.defaults/README.md", wrkDr+"/README.md")

		// create a PR template file
		fmt.Printf("\nstep 05/15 => Creating PR template file in %s directory...\n", projectName)
		copy("./lib/.defaults/.github/PULL_REQUEST_TEMPLATE.md", wrkDr+"/.github/PULL_REQUEST_TEMPLATE.md")
	}
}

func createRubocopFile() {
	if isRubocop == "y" {
		fmt.Printf("\nstep 02/15 => Initializing rubocop in %s directory...\n", projectName)
		copy("./lib/.ruby/.rubocop.yml", wrkDr+"/.rubocop.yml")
	}
}

func copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}
