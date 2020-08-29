package main

import (
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/mitchellh/go-homedir"
)

func TestIsDirectoryExists(t *testing.T) {
	// when the directory exists
	tmpDir := ".tmp"
	os.Mkdir(getHomeDirectory()+"/.tmp", 0755)
	answer := isDirectoryExists(tmpDir)
	tempDirToRemove := getHomeDirectory() + "/.tmp"
	if answer == false {
		t.Errorf("The isDirectoryExists() should return true if the directory exists, but it returned %v\n", answer)
	}
	defer exec.Command("rmdir", tempDirToRemove).Run()

	// when the directory doesn't exist
	answer = isDirectoryExists("jjjjjjjjjjjaksaaaaaaaaanvvvvvvvvvvvvvvjshdbcbfh")
	if answer == true {
		t.Errorf("The isDirectoryExists() should return false if the directory does not exists, but it returned %v\n", answer)
	}
}
func TestGetHomeDirectory(t *testing.T) {
	homeDirectory, _ := homedir.Dir()
	if homeDirectory != getHomeDirectory() {
		t.Errorf("getHomeDirectory() should return a home directory but that failed")
	}
}

func TestCreateProjectDirectory(t *testing.T) {
	workingDir = "."
	projectName = "tmpProject"
	currDir, _ := os.Getwd()
	createProjectDirectory()
	_, err := os.Stat(currDir + "/" + projectName)
	if os.IsNotExist(err) {
		t.Errorf("createProjectDirectory() should create a project directory but it failed")
	}

	workingDir = "jjjjjjjjjjjaksaaaaaaaaanvvvvvvvvvvvvvvjshdbcbfh"
	createProjectDirectory()
	_, err = os.Stat(currDir + "/" + projectName)
	if os.IsNotExist(err) {
		t.Errorf("createProjectDirectory() should create a project directory but it failed")
	}

	os.Mkdir(getHomeDirectory()+"/.tmp", 0755)
	workingDir = ".tmp"
	createProjectDirectory()
	if isDirectoryExists(".tmp"+"/"+projectName) == false {
		t.Errorf("createProjectDirectory() should create a project directory but it failed")
	}

	exec.Command("rmdir", projectName)
	strRm := "rm -r " + getHomeDirectory() + "/" + workingDir
	strRmArgs := strings.Split(strRm, " ")
	exec.Command(strRmArgs[0], strRmArgs[1:]...).Run()
}

func TestCreateGithubActionsDirectory(t *testing.T) {
	isGithub = "y"
	projectName = "tmpProject"
	workingDir = "."
	createProjectDirectory()
	currDir, _ := os.Getwd()
	wrkDr = currDir + "/" + projectName
	createGithubActionsDirectory()
	_, err = os.Stat(wrkDr + "/.github")
	if os.IsNotExist(err) {
		t.Errorf("createGithubActionsDirectory() should create github directory but it is failing")
	}

	strRmTwo := "rm -r " + "/" + wrkDr
	strRmArgs := strings.Split(strRmTwo, " ")
	exec.Command(strRmArgs[0], strRmArgs[1:]...).Run()
}

// func TestCreateRubocopFile(t *testing.T) {
// 	isRubocop = "y"
// 	projectName = "tmpProject"
// 	workingDir = "."
// 	createProjectDirectory()
// 	currDir, _ := os.Getwd()
// 	wrkDr = currDir + "/" + projectName
// 	createRubocopFile()
// 	_, err := os.Stat(wrkDr + "/.rubocop.yml")
// 	if os.IsNotExist(err) {
// 		t.Errorf("createRubocopFile() should create a rubocop file but it failed to do so!")
// 	}
// 	rmFilStr := "rm -r " + wrkDr
// 	strRmFilArgs := strings.Split(rmFilStr, " ")
// 	exec.Command(strRmFilArgs[0], strRmFilArgs[1:]...).Run()
// }

func TestRubyBoiler(t *testing.T) {
	language = 1
	isGithub = "y"
	isRubocop = "y"
	isTests = "y"
	testFramework = 1
	workingDir = "."
	projectName = "tmpProject"
	main()
	os.Chdir("../")
	currDir, _ := os.Getwd()
	_, err := os.Stat(currDir)
	_, err1 := os.Stat(currDir + "/" + projectName + "/lib")
	_, err2 := os.Stat(currDir + "/" + projectName + "/bin")
	_, err3 := os.Stat(currDir + "/" + projectName + "/spec")
	_, err4 := os.Stat(currDir + "/" + projectName + "/Gemfile")
	_, err5 := os.Stat(currDir + "/" + projectName + "/.github")
	_, err6 := os.Stat(currDir + "/" + projectName + "/.rspec")
	_, err7 := os.Stat(currDir + "/" + projectName + "/.rubocop.yml")
	if os.IsNotExist(err) {
		t.Error("rubyBoiler() should create project in the current dir but it failed")
	}
	if os.IsNotExist(err1) {
		t.Error("rubyBoiler() should create lib folder in the project dir but it failed")
	}
	if os.IsNotExist(err2) {
		t.Error("rubyBoiler() should create bin folder in the current dir but it failed")
	}
	if os.IsNotExist(err3) {
		t.Error("rubyBoiler() should create spec folder in the project dir but it failed")
	}
	if os.IsNotExist(err4) {
		t.Error("rubyBoiler() should create gemfile in the the project dir but it failed")
	}
	if os.IsNotExist(err5) {
		t.Error("rubyBoiler() should create .github folder in the project dir but it failed")
	}
	if os.IsNotExist(err6) {
		t.Error("rubyBoiler() should create .rspec file in the project dir but it failed")
	}
	if os.IsNotExist(err7) {
		t.Error("rubyBoiler() should create .rubocop.yml file in the project dir but it failed")
	}

	rmPrStr := "rm -r " + currDir + "/" + projectName
	argsRmPrStr := strings.Split(rmPrStr, " ")
	exec.Command(argsRmPrStr[0], argsRmPrStr[1:]...).Run()
}

func TestRorBoiler(t *testing.T) {
	language = 2
	isGithub = "y"
	isRubocop = "y"
	workingDir = "."
	projectName = "tmpRorProject"
	main()
	os.Chdir("../")
	currDir, _ := os.Getwd()
	_, err := os.Stat(currDir + "/" + projectName)
	_, err1 := os.Stat(currDir + "/" + projectName + "/.github")
	_, err2 := os.Stat(currDir + "/" + projectName + "/.stylelintrc.json")
	_, err3 := os.Stat(currDir + "/" + projectName + "/.rubocop.yml")
	_, err4 := os.Stat(currDir + "/" + projectName + "/.github/PULL_REQUEST_TEMPLATE.md")
	_, err5 := os.Stat(currDir + "/" + projectName + "/.github/workflows")
	_, err6 := os.Stat(currDir + "/" + projectName + "/.github/workflows/linters.yml")
	if os.IsNotExist(err) {
		t.Error("rorBoiler() should create project in the current dir but it failed")
	}
	if os.IsNotExist(err1) {
		t.Error("rorBoiler() should create .github folder in the project dir but it failed")
	}
	if os.IsNotExist(err2) {
		t.Error("rorBoiler() should create .stylelintrc.json file in the project dir but it failed")
	}
	if os.IsNotExist(err3) {
		t.Error("rorBoiler() should create .rubocop.yml file in the project dir but it failed")
	}
	if os.IsNotExist(err4) {
		t.Error("rorBoiler() should create PULL_REQUEST_TEMPLATE.md file in the .github dir but it failed")
	}
	if os.IsNotExist(err5) {
		t.Error("rorBoiler() should create workflows folder in the .github dir but it failed")
	}
	if os.IsNotExist(err6) {
		t.Error("rorBoiler() should create linters.yml file in the .github/workflows dir but it failed")
	}
	rmPrStr := "rm -r " + currDir + "/" + projectName
	argsRmPrStr := strings.Split(rmPrStr, " ")
	exec.Command(argsRmPrStr[0], argsRmPrStr[1:]...).Run()
}
