package main

import (
	"os"
	"os/exec"
	"testing"
)

func TestIsDirectoryExists(t *testing.T) {
	tmpDir := getHomeDirectory() + "/.tmp"
	os.Mkdir(tmpDir, 0755)
	answer := isDirectoryExists(tmpDir)
	if answer == false {
		t.Errorf("The isDirectoryExists() should return true if the directory exists, but it returned %v\n", answer)
	}
	exec.Command("rmdir", tmpDir).Run()
}

func TestIsGitHub(t *testing.T) {}
