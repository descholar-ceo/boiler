package main

import (
	"os"
	"os/exec"
	"strings"
	"testing"
)




func TestRubyBoiler(t *testing.T) {
	language = 1
	isGithub = "y"
	isRubocop = "y"
	isTests = "y"
	testFramework = 1
	workingDir = "."
	projectName = "tmpProject"
	main()
	os.Chdir("../")
	currDir, _ := os.Getwd()
	_, err := os.Stat(currDir)
	_, err1 := os.Stat(currDir + "/" + projectName + "/lib")
	_, err2 := os.Stat(currDir + "/" + projectName + "/bin")
	_, err3 := os.Stat(currDir + "/" + projectName + "/spec")
	_, err4 := os.Stat(currDir + "/" + projectName + "/Gemfile")
	// _, err5 := os.Stat(currDir + "/" + projectName + "/.github")
	_, err6 := os.Stat(currDir + "/" + projectName + "/.rspec")
	// _, err7 := os.Stat(currDir + "/" + projectName + "/.rubocop.yml")
	if os.IsNotExist(err) {
		t.Error("rubyBoiler() should create project in the current dir but it failed")
	}
	if os.IsNotExist(err1) {
		t.Error("rubyBoiler() should create lib folder in the project dir but it failed")
	}
	if os.IsNotExist(err2) {
		t.Error("rubyBoiler() should create bin folder in the current dir but it failed")
	}
	if os.IsNotExist(err3) {
		t.Error("rubyBoiler() should create spec folder in the project dir but it failed")
	}
	if os.IsNotExist(err4) {
		t.Error("rubyBoiler() should create gemfile in the the project dir but it failed")
	}
	// if os.IsNotExist(err5) {
	// 	t.Error("rubyBoiler() should create .github folder in the project dir but it failed")
	// }
	if os.IsNotExist(err6) {
		t.Error("rubyBoiler() should create .rspec file in the project dir but it failed")
	}
	// if os.IsNotExist(err7) {
	// 	t.Error("rubyBoiler() should create .rubocop.yml file in the project dir but it failed")
	// }

	rmPrStr := "rm -r " + currDir + "/" + projectName
	argsRmPrStr := strings.Split(rmPrStr, " ")
	exec.Command(argsRmPrStr[0], argsRmPrStr[1:]...).Run()
}

func TestRorBoiler(t *testing.T) {
	language = 2
	isGithub = "y"
	isRubocop = "y"
	workingDir = "."
	projectName = "tmpRorProject"
	main()
	os.Chdir("../")
	currDir, _ := os.Getwd()
	_, err := os.Stat(currDir + "/" + projectName)
	_, err1 := os.Stat(currDir + "/" + projectName + "/.github")
	_, err2 := os.Stat(currDir + "/" + projectName + "/.stylelintrc.json")
	_, err3 := os.Stat(currDir + "/" + projectName + "/.rubocop.yml")
	_, err4 := os.Stat(currDir + "/" + projectName + "/.github/PULL_REQUEST_TEMPLATE.md")
	_, err5 := os.Stat(currDir + "/" + projectName + "/.github/workflows")
	_, err6 := os.Stat(currDir + "/" + projectName + "/.github/workflows/linters.yml")
	if os.IsNotExist(err) {
		t.Error("rorBoiler() should create project in the current dir but it failed")
	}
	if os.IsNotExist(err1) {
		t.Error("rorBoiler() should create .github folder in the project dir but it failed")
	}
	if os.IsNotExist(err2) {
		t.Error("rorBoiler() should create .stylelintrc.json file in the project dir but it failed")
	}
	if os.IsNotExist(err3) {
		t.Error("rorBoiler() should create .rubocop.yml file in the project dir but it failed")
	}
	if os.IsNotExist(err4) {
		t.Error("rorBoiler() should create PULL_REQUEST_TEMPLATE.md file in the .github dir but it failed")
	}
	if os.IsNotExist(err5) {
		t.Error("rorBoiler() should create workflows folder in the .github dir but it failed")
	}
	if os.IsNotExist(err6) {
		t.Error("rorBoiler() should create linters.yml file in the .github/workflows dir but it failed")
	}
	rmPrStr := "rm -r " + currDir + "/" + projectName
	argsRmPrStr := strings.Split(rmPrStr, " ")
	exec.Command(argsRmPrStr[0], argsRmPrStr[1:]...).Run()
}

func TestMain(t *testing.T) {
	language = 5
	main()
	if language != 0 {
		t.Errorf("If the provided language number is not valid, the program should return and language choice should be reset to zero value, but it is still on %v", language)
	}
}
