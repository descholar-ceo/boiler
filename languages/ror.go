package languages

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/descholar-ceo/boiler/utils"
)

// RorBoiler is a boilerplate generator for ror
func RorBoiler() {
	workingDir := utils.AskWorkingDirectory()
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
}
