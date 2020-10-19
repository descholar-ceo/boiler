package utils

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/mitchellh/go-homedir"
)

// RE-USABLE FUNCTIONS
func generateStdin(in *os.File) *os.File {
	if in == nil {
		in = os.Stdin
	}
	return in
}

func askBasicQuestions(in *os.File, variable string, textToAsk string) string {
	in = generateStdin(in)
	fmt.Println(textToAsk)
	if _, err := fmt.Fscanf(in, "%s", &variable); err != nil {
		log.Fatal(`|=====================Error message===================|
							 ||    Empty response, please enter a value!          ||
							 |=====================================================|`)
	}
	return variable
}

// AskWorkingDirectory function
func AskWorkingDirectory(in *os.File) string {
	var workingDir string
	return askBasicQuestions(in, workingDir, "\nEnter the working directory (Enter a dot (.) for the current directory):")
}

// AskRubocop is a function which asks a user if they will use rubocop and then returns the answer of the user
func AskRubocop(in *os.File) string {
	var isRubocop string
	return askBasicQuestions(in, isRubocop, "\nWill you use Rubocop as a linter? Enter y for yes or any other key for no")
}

// AskGithub is a function which asks a user if they will use github
func AskGithub(in *os.File) string {
	var isGithub string
	return askBasicQuestions(in, isGithub, "\nWill you use github as a collaboration tool? Enter y for yes or any other key for no")
}

// AskDatabase is a function which asks a user which database they will use and returns the answer
func AskDatabase(in *os.File) string {
	in = generateStdin(in)
	var database string
	var tmpDb int
	fmt.Println("\nSelect Enter the number corresponding to the database you want to use: ")
	fmt.Println("\n1.sqlite3\n2.postgresql\n3.mysql\n4.oracle\n5.frontbase\n6.db2(ibm_db)\n7.sqlserver\n8.jdbcmysql\n9.jdbcpostgresql\n10.jdbcsqlite3\n11.jdbc")
	if _, err := fmt.Fscanf(in, "%d", &tmpDb); err != nil {
		log.Fatal(err)
	}
	switch tmpDb {
	case 1:
		database = "sqlite3"
	case 2:
		database = "postgresql"
	case 3:
		database = "mysql"
	case 4:
		database = "oracle"
	case 5:
		database = "frontbase"
	case 6:
		database = "ibm_db"
	case 7:
		database = "sqlserver"
	case 8:
		database = "jdbcmysql"
	case 9:
		database = "jdbcpostgresql"
	case 10:
		database = "jdbcsqlite3"
	case 11:
		database = "jdbc"
	default:
		fmt.Println(`The database you choose is not supported by rails yet!
		So, I will create for you a rails app with a default database which is sqlite3`)
		database = "sqlite3"
	}

	return database
}

// AskProjectName function
func AskProjectName(in *os.File) string {
	var projectName string
	return askBasicQuestions(in, projectName, "\n\nWhat is the project name you want to use?")
}

// Copy function to copy files
func Copy(src, dst string) {
	source, _ := os.Open(src)

	defer source.Close()

	destination, _ := os.Create(dst)

	defer destination.Close()
	if _, err := io.Copy(destination, source); err != nil {
		log.Fatal(err)
	}
}

// WriteToFile is a function which used to write to file
func WriteToFile(file, stringToWrite string) {
	mFile, _ := os.OpenFile(file, os.O_APPEND|os.O_WRONLY, 0644)
	fmt.Fprintln(mFile, stringToWrite)
	mFile.Close()
}

// GetHomeDirectory is used to get home directory
func GetHomeDirectory() string {
	homeDirectory, _ := homedir.Dir()
	return homeDirectory
}

// IsDirectoryExists function is used to prove if a passed string stands for a directory
func IsDirectoryExists(directory string) bool {
	_, err := os.Stat(GetHomeDirectory() + "/" + directory)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

// CreateGithubActionsDirectory is a function which creates github actions directory
func CreateGithubActionsDirectory(isGithub, wrkDr, projectName string) {
	if isGithub == "y" {
		fmt.Printf("\nInitializing github actions in %s directory...\n", projectName)
		os.Mkdir(wrkDr+"/.github", 0755)
		os.Mkdir(wrkDr+"/.github/workflows", 0755)
	}
}

// CreateProjectDirectory function is to create projec
func CreateProjectDirectory(workingDir, projectName string) string {
	var tmpWrkDr, wrkDr string
	if workingDir == "." {
		// create project in current directory
		tmpWrkDr, _ = os.Getwd()
		wrkDr = tmpWrkDr + "/" + projectName
	} else {
		// checking if a directory exists
		if IsDirectoryExists(workingDir) {
			wrkDr = GetHomeDirectory() + "/" + workingDir + "/" + projectName
		} else {
			fmt.Println("The directory you entered does not exists, your project will be created in the current directory")
			tmpWrkDr, _ = os.Getwd()
			wrkDr = tmpWrkDr + "/" + projectName
		}
	}

	// create a project directory
	fmt.Printf("\nCreating directory to %s...\n", projectName)
	os.Mkdir(wrkDr, 0755)

	return wrkDr
}

// AskTests is a function which asks if a user will use tests
func AskTests(in *os.File) (string, string) {
	in = generateStdin(in)
	var isTests string
	var testFrameworkNumber int
	var testFramework string
	fmt.Println("\n\nWill you write some unit tests for your project? Enter y for yes or any other key for no")
	if _, err := fmt.Fscanf(in, "%s", &isTests); err != nil {
		log.Fatal(err)
	}
	if isTests == "y" || isTests == "Y" {
		fmt.Println("\nChoose a number which corresponds to the testing framework you will be using:\n1.RSpec")
		if _, err := fmt.Fscanf(in, "%d", &testFrameworkNumber); err != nil {
			log.Fatal(err)
		}
		if testFrameworkNumber != 1 {
			for i := 0; i < 5; i++ {
				fmt.Println("\nChoose a number which corresponds to the testing framework you will be using:\n1.RSpec")
				if _, err := fmt.Fscanf(in, "%d", &testFrameworkNumber); err != nil {
					log.Fatal(err)
				}
				if testFrameworkNumber == 1 {
					break
				}
			}
			fmt.Println("\nThe testing framework you chose is not supported")
		}
	}
	switch testFrameworkNumber {
	case 1:
		testFramework = "rspec"
	}
	return testFramework, isTests
}

// DisplayLastCommands func shows a user next steps
func DisplayLastCommands() string {
	return `

==============================================

Your project has been initialized successfully
The remaining task is to go on github and create a repository and copy its url
Come back in the root directory of this project and 
Run the following commands respectifully

1. git remote add .
2. git commit -m "Initial commit"
3. git remote add origin [Paste the url you copied from github]"
4. git push -u origin master"

Congratulations and good luck for your new project`
}
