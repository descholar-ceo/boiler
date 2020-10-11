package main

import "testing"

// import (
// 	"os"
// 	"os/exec"
// 	"strings"
// 	"testing"
// )

func TestMain(t *testing.T) {
	language = 5
	main()
	if language != 0 {
		t.Errorf("If the provided language number is not valid, the program should return and language choice should be reset to zero value, but it is still on %v", language)
	}
}
