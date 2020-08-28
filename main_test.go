package main

import (
	"os"
	"testing"
)

func TestIsDirectoryExists(t *testing.T) {
	os.Mkdir(getHomeDirectory()+"/.tmp", 0755)
	answer := isDirectoryExists(getHomeDirectory() + "/.tmp")
	if answer == false {
		t.Errorf("The isDirectoryExists() should return true if the directory exists, but it returned %v\n", answer)
	}
}
