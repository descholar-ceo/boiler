package utils

import (
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/mitchellh/go-homedir"
)

func TestGetHomeDirectory(t *testing.T) {
	homeDirectory, _ := homedir.Dir()
	if homeDirectory != GetHomeDirectory() {
		t.Errorf("getHomeDirectory() should return a home directory but that failed")
	}
}
func TestIsDirectoryExists(t *testing.T) {
	// when the directory exists
	tmpDir := ".tmp"
	os.Mkdir(GetHomeDirectory()+"/.tmp", 0755)
	answer := IsDirectoryExists(tmpDir)
	tempDirToRemove := GetHomeDirectory() + "/.tmp"
	if answer == false {
		t.Errorf("The isDirectoryExists() should return true if the directory exists, but it returned %v\n", answer)
	}
	defer exec.Command("rmdir", tempDirToRemove).Run()

	// when the directory doesn't exist
	answer = IsDirectoryExists("jjjjjjjjjjjaksaaaaaaaaanvvvvvvvvvvvvvvjshdbcbfh")
	if answer == true {
		t.Errorf("The isDirectoryExists() should return false if the directory does not exists, but it returned %v\n", answer)
	}
}

func TestCreateProjectDirectory(t *testing.T) {
	workingDir := "."
	projectName := "tmpProject"
	currDir, _ := os.Getwd()
	CreateProjectDirectory(workingDir, projectName)
	_, err := os.Stat(currDir + "/" + projectName)
	if os.IsNotExist(err) {
		t.Errorf("createProjectDirectory() should create a project directory but it failed")
	}

	workingDir = "jjjjjjjjjjjaksaaaaaaaaanvvvvvvvvvvvvvvjshdbcbfh"
	CreateProjectDirectory(workingDir, projectName)
	_, err = os.Stat(currDir + "/" + projectName)
	if os.IsNotExist(err) {
		t.Errorf("createProjectDirectory() should create a project directory but it failed")
	}

	os.Mkdir(GetHomeDirectory()+"/.tmp", 0755)
	workingDir = ".tmp"
	CreateProjectDirectory(workingDir, projectName)
	if IsDirectoryExists(".tmp"+"/"+projectName) == false {
		t.Errorf("createProjectDirectory() should create a project directory but it failed")
	}

	exec.Command("rmdir", projectName)
	strRm := "rm -r " + GetHomeDirectory() + "/" + workingDir
	strRmArgs := strings.Split(strRm, " ")
	exec.Command(strRmArgs[0], strRmArgs[1:]...).Run()
}

func TestCreateGithubActionsDirectory(t *testing.T) {
	isGithub := "y"
	projectName := "tmpProject"
	workingDir := "."
	CreateProjectDirectory(workingDir, projectName)
	currDir, _ := os.Getwd()
	wrkDr := currDir + "/" + projectName
	CreateGithubActionsDirectory(isGithub, wrkDr, projectName)
	_, err := os.Stat(wrkDr + "/.github")
	if os.IsNotExist(err) {
		t.Errorf("createGithubActionsDirectory() should create github directory but it is failing")
	}

	strRmTwo := "rm -r " + "/" + wrkDr
	strRmArgs := strings.Split(strRmTwo, " ")
	exec.Command(strRmArgs[0], strRmArgs[1:]...).Run()
}

func TestAskWorkingDirectory(t *testing.T) {

}
