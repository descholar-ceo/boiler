package main

import (
	"os"
	"os/exec"
	"testing"
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
}

func TestIsGitHub(t *testing.T) {

}
