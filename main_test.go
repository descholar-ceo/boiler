package main

import (
	"os"
	"os/exec"
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
	answer := isDirectoryExists(currDir + "/" + projectName)
	if answer == false {
		t.Errorf("createProjectDirectory() should create a project directory but it failed")
	}
}
