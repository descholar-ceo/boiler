package main

import (
	"os"
	"testing"
)

func TestIsDirectoryExists(t *testing.T) {
	os.Mkdir(".tmp", 0755)
	answer := isDirectoryExists(".tmp")
	if !answer {

	}
}