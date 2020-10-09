package languages

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/descholar-ceo/boiler/utils"
)

// RorBoiler is a boilerplate generator for ror
func RorBoiler() {
	workingDir := utils.AskWorkingDirectory()
	var wrkDr string
	projectName := utils.AskProjectName()
	isGithub := utils.AskGithub()
	isRubocop := utils.AskRubocop()
	database := utils.AskDatabase()

	// checking the Rails installation
	fmt.Println(`Make sure that rails and ruby are installed correctly on your system, and is working well, 
	if rails is not installed yet, then this time your project initialization will take some time. 
	If ruby is not installed yet, please refer to this link for a proper ruby installation: 
	https://www.theodinproject.com/courses/ruby-programming/lessons/installing-ruby-ruby-programming`)

	fmt.Println("\n\nChecking Rails installation on your computer...")
	mStr := "gem install rails"
	argsMStr := strings.Split(mStr, " ")
	exec.Command(argsMStr[0], argsMStr[1:]...).Run()

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
	exec.Command(args[0], args[1:]...).Run()
}
