package utils

import (
	"io"
	"io/ioutil"
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
	in, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatal(err)
	}
	defer in.Close()

	_, err = io.WriteString(in, "directory")
	if err != nil {
		t.Fatal(err)
	}
	_, err = in.Seek(0, os.SEEK_SET)
	if err != nil {
		t.Fatal(err)
	}
	workingDir := AskWorkingDirectory(in)
	if workingDir != "directory" {
		t.Errorf("The AskWorkingDirectory is not working!")
	}
}
func TestAskRubocop(t *testing.T) {
	in, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatal(err)
	}
	defer in.Close()

	_, err = io.WriteString(in, "y")
	if err != nil {
		t.Fatal(err)
	}
	_, err = in.Seek(0, os.SEEK_SET)
	if err != nil {
		t.Fatal(err)
	}
	workingDir := AskRubocop(in)
	if workingDir != "y" {
		t.Errorf("The AskRubocop is not working!")
	}
}
func TestAskGithub(t *testing.T) {
	in, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatal(err)
	}
	defer in.Close()

	_, err = io.WriteString(in, "y")
	if err != nil {
		t.Fatal(err)
	}
	_, err = in.Seek(0, os.SEEK_SET)
	if err != nil {
		t.Fatal(err)
	}
	workingDir := AskGithub(in)
	if workingDir != "y" {
		t.Errorf("The AskGithub is not working!")
	}
}
func TestAskDatabase(t *testing.T) {
	in, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatal(err)
	}
	defer in.Close()

	_, err = io.WriteString(in, "2")
	if err != nil {
		t.Fatal(err)
	}
	_, err = in.Seek(0, os.SEEK_SET)
	if err != nil {
		t.Fatal(err)
	}
	workingDir := AskDatabase(in)
	if workingDir != "postgresql" {
		t.Errorf("The AskDatabase is not working!")
	}
}

func TestAskProjectName(t *testing.T) {
	in, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatal(err)
	}
	defer in.Close()

	_, err = io.WriteString(in, "new_project")
	if err != nil {
		t.Fatal(err)
	}
	_, err = in.Seek(0, os.SEEK_SET)
	if err != nil {
		t.Fatal(err)
	}
	workingDir := AskProjectName(in)
	if workingDir != "new_project" {
		t.Errorf("The AskProjectName is not working!")
	}
}
func TestAskTests(t *testing.T) {
	in, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatal(err)
	}
	defer in.Close()

	_, err = io.WriteString(in, "y\n"+"1\n")
	if err != nil {
		t.Fatal(err)
	}
	_, err = in.Seek(0, os.SEEK_SET)
	if err != nil {
		t.Fatal(err)
	}
	testFramework, isTests := AskTests(in)
	if testFramework != "rspec" || isTests != "y" {
		t.Errorf("The AskTests is not working, we should have testframework = %s and isTest = %s", testFramework, isTests)
	}
}
