package languages

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/descholar-ceo/boiler/utils"
)

func mQuestions() (string, string, string, string, string) {
	workingDir = utils.AskWorkingDirectory(nil)
	projectName = utils.AskProjectName(nil)
	isGithub = utils.AskGithub(nil)
	isRubocop = utils.AskRubocop(nil)
	database = utils.AskDatabase(nil)
	return workingDir, projectName, isGithub, isRubocop, database
}

// RorBoiler is a boilerplate generator for ror
func RorBoiler() {
	var wrkDr string
	workingDir, projectName, isGithub, isRubocop, database = mQuestions()

	// moving to the project dir
	if workingDir == "." {
		wrkDr, _ = os.Getwd()
	} else {
		if utils.IsDirectoryExists(workingDir) {
			wrkDr = utils.GetHomeDirectory() + "/" + workingDir
		} else {
			fmt.Printf("\n%s does not exist, your ruby on rails project will be created in the current directory\n",
				workingDir)
			wrkDr, _ = os.Getwd()
		}
	}

	// create a project with rails
	fmt.Println("\nChecking out your working directory")
	os.Chdir(wrkDr)

	// generating ror project with rails
	fmt.Println(`
Generating your Rails project using Rails installed on your machine, This might take several 
minutes depending on the internet connection you have, please bear with us, and wait...`)
	railsStr := "rails new " + projectName + " --database=" + strings.Trim(database, "\"")
	args := strings.Split(railsStr, " ")
	if err := exec.Command(args[0], args[1:]...).Run(); err != nil {
		log.Fatal(err)
	}

	// changing current dir
	fmt.Println("\nChecking out your project workspace...")
	os.Chdir(projectName)

	// tempalting the README file
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
