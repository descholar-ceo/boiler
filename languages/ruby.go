package languages

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/descholar-ceo/boiler/utils"
)

// RubyBoiler function which will provide a bilerplate of the ruby project
func RubyBoiler() {

	// declaration and initialization of variables
	workingDir = utils.AskWorkingDirectory(nil)
	projectName = utils.AskProjectName()
	isGithub = utils.AskGithub()
	isRubocop = utils.AskRubocop(nil)
	testFramework, isTests = utils.AskTests()

	// informing a user about the ruby installation
	fmt.Println(`
	Make sure that ruby is well installed, and your bundler is working well. If it is not the case 
	please refer to this link for ruby installation guides: 
	https://www.theodinproject.com/courses/ruby-programming/lessons/installing-ruby-ruby-programming
	`)

	// create project dir
	wrkDr := utils.CreateProjectDirectory(workingDir, projectName)

	// initialize rubocop
	if isRubocop == "y" {
		fmt.Printf("\nInitializing rubocop in %s directory...\n", projectName)
		utils.Copy("./lib/.ruby/.rubocop.yml", wrkDr+"/.rubocop.yml")
	}

	if isGithub == "y" {
		// initialize github actions
		utils.CreateGithubActionsDirectory(isGithub, wrkDr, projectName)
		utils.Copy("./lib/.ruby/.github/workflows/linters.yml", wrkDr+"/.github/workflows/linters.yml")
		utils.Copy("./lib/.ruby/.github/workflows/tests.yml", wrkDr+"/.github/workflows/tests.yml")

		// create a PR template file
		fmt.Printf("\nCreating PR template file in %s directory...\n", projectName)
		utils.Copy("./lib/.defaults/.github/PULL_REQUEST_TEMPLATE.md", wrkDr+"/.github/PULL_REQUEST_TEMPLATE.md")
	}

	// create a readme file
	fmt.Printf("\nCreating README file in %s directory...\n", projectName)
	utils.Copy("./lib/.defaults/README.md", wrkDr+"/README.md")

	// create initial files
	fmt.Printf("\nCreating lib folder in %s directory...\n", projectName)
	os.Mkdir(wrkDr+"/lib", 0755)

	fmt.Printf("\nCreating bin folder in %s directory...\n", projectName)
	os.Mkdir(wrkDr+"/bin", 0755)

	fmt.Printf("\nAdding .gitkeep file in %s/lib directory...\n", projectName)
	os.Create(wrkDr + "/lib/.gitkeep")

	fmt.Printf("\nCreating main.rb file in %s directory...\n", projectName)
	os.Create(wrkDr + "/bin/main.rb")
	utils.WriteToFile(wrkDr+"/bin/main.rb", "puts 'Hello from Boiler! Welcome to the new world!'")

	// change working dir
	fmt.Println("\nChecking out your project dir...")
	os.Chdir(wrkDr)

	// initialize gemfile
	fmt.Printf("\nInitializing gem in %s directory...\n", projectName)
	exec.Command("bundle", "init").Run()

	if isTests == "y" && testFramework == "rspec" {
		// initialize rspec
		fmt.Printf("\nInitializing rspec in %s directory...\n", projectName)
		str := "gem install rspec"
		argsRspec := strings.Split(str, " ")
		exec.Command(argsRspec[0], argsRspec[1:]...).Run()
		utils.WriteToFile("Gemfile", "gem 'rspec', '~>3.0'")
		exec.Command("rspec", "--init").Run()
	}

	if isRubocop == "y" {
		// install rubocop in gems
		fmt.Println("\nWriting gems...")
		utils.WriteToFile("Gemfile", "gem 'rubocop', '~>0.81.0'")
	}

	// initialize git
	fmt.Printf("\nInitializing git in %s directory...\n", projectName)
	exec.Command("git", "init").Run()

	// installing bundler gems
	fmt.Printf("\nInstalling gems %s directory, this might take some minutes, please wait...\n", projectName)
	exec.Command("bundle", "install").Run()
}
