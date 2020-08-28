package main

import (
	"fmt"
	"os"
	"testing"
)

func TestIsDirectoryExists(t *testing.T) {
	os.Mkdir(getHomeDirectory()+"/.tmp", 0755)
	answer := isDirectoryExists(getHomeDirectory() + "/descholar")
	fmt.Printf("the value of the answer is %v\n", getHomeDirectory()+"/descholar")
	if answer == false {
		t.Errorf("The isDirectoryExists() should return true if the directory exists, but it returned %v\n", answer)
	}
}
