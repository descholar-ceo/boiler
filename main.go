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
	isRubocop     int
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
	fmt.Println("Choose a number which correspond to the language you will be using:\n1.Ruby")
	fmt.Scan(&language)
	if language == 1 {
		// language is ruby
		rubyBoiler()
	} else {

		// language is not ruby
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

func rubyBoiler() {

	fmt.Println("\nWill you use Rubocop as a linter? Enter y for yes or any other key for no")
	fmt.Scan(&isRubocop)

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
			fmt.Println("\nThe the testing framework you chose is not supported")
		}
	}

	// will you use github?
	fmt.Println("\nWill you use github as a collaboration tool? Enter y for yes or any other key for no")
	fmt.Scan(&isGithub)

	fmt.Println("\n\n\nThe following are your preferences, we will setup your project depending on your preferences.")
	fmt.Printf("\n\nWorking dir : %v\n", workingDir)
	fmt.Printf("Project name : %v\n", projectName)
	fmt.Printf("Language name : %v\n", language)
	fmt.Printf("Will you write unit test? : %v\n", isTests)
	fmt.Printf("Testing framework : %v\n", testFramework)
	fmt.Printf("Will you use github? : %v\n\n", isGithub)

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
	fmt.Printf("\nCreating directory to %s...\n", projectName)
	os.Mkdir(wrkDr, 0755)

	// initialize rubocop
	if isRubocop == 1 {
		fmt.Printf("\nInitializing rubocop in %s directory...\n", projectName)
		copy("./lib/.ruby/.rubocop.yml", wrkDr+"/.rubocop.yml")
	}

	// initialize github actions
	fmt.Printf("\nInitializing github actions in %s directory...\n", projectName)
	os.Mkdir(wrkDr+"/.github", 0755)
	os.Mkdir(wrkDr+"/.github/workflows", 0755)
	copy("./lib/.ruby/.github/workflows/linters.yml", wrkDr+"/.github/workflows/linters.yml")
	copy("./lib/.ruby/.github/workflows/tests.yml", wrkDr+"/.github/workflows/tests.yml")

	// create a readme file
	fmt.Printf("\nCreating README file in %s directory...\n", projectName)
	copy("./lib/.ruby/README.md", wrkDr+"/README.md")

	// create a PR template file
	fmt.Printf("\nCreating PR template file in %s directory...\n", projectName)
	copy("./lib/.ruby/.github/PULL_REQUEST_TEMPLATE.md", wrkDr+"/.github/PULL_REQUEST_TEMPLATE.md")

	// change working dir
	os.Chdir(wrkDr)

	// initialize gemfile
	defer fmt.Printf("\nInitializing gem in %s directory...\n", projectName)
	defer exec.Command("bundle", "init").Run()

	// initialize rspec
	defer fmt.Printf("\nInitializing rspec in %s directory...\n", projectName)
	defer exec.Command("rspec", "--init").Run()

	// initialize git
	defer fmt.Printf("\nInitializing git in %s directory...\n", projectName)
	defer exec.Command("git", "init").Run()
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
