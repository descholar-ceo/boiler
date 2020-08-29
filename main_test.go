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

func TestCreateRubocopFile(t *testing.T) {
	isRubocop = "y"
	projectName = "tmpProject"
	workingDir = "."
	createProjectDirectory()
	currDir, _ := os.Getwd()
	wrkDr = currDir + "/" + projectName
	createRubocopFile()
	_, err := os.Stat(wrkDr + "/.rubocop.yml")
	if os.IsNotExist(err) {
		t.Errorf("createRubocopFile() should create a rubocop file but it failed to do so!")
	}

	rmFilStr := "rm -r " + wrkDr
	strRmFilArgs := strings.Split(rmFilStr, " ")
	exec.Command(strRmFilArgs[0], strRmFilArgs[1:]...).Run()
}

func TestRubyBoiler(t *testing.T) {
	isGithub = "y"
	isRubocop = "y"
	isTests = "y"
	testFramework = 1
	workingDir = "."
	projectName = "tmpProject"
	rubyBoiler()
	currDir, _ := os.Getwd()
	_, err := os.Stat(currDir)
	if os.IsNotExist(err) {
		t.Error("rubyBoiler() should create project in the current dir but it failed")
	}
}
