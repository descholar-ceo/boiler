package utils

import (
	"os"
	"os/exec"
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
