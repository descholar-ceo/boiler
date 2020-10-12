package languages

import (
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestRubyBoiler(t *testing.T) {
	isGithub = "y"
	isRubocop = "y"
	isTests = "y"
	testFramework = "rspec"
	workingDir = "."
	projectName = "tmpProject"
	RubyBoiler()
	os.Chdir("../")
	currDir, _ := os.Getwd()
	_, err := os.Stat(currDir)
	_, err1 := os.Stat(currDir + "/" + projectName + "/lib")
	_, err2 := os.Stat(currDir + "/" + projectName + "/bin")
	_, err3 := os.Stat(currDir + "/" + projectName + "/spec")
	_, err4 := os.Stat(currDir + "/" + projectName + "/Gemfile")
	_, err5 := os.Stat(currDir + "/" + projectName + "/.github")
	_, err6 := os.Stat(currDir + "/" + projectName + "/.rspec")
	_, err7 := os.Stat(currDir + "/" + projectName + "/.rubocop.yml")
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
	if os.IsNotExist(err5) {
		t.Error("rubyBoiler() should create .github folder in the project dir but it failed")
	}
	if os.IsNotExist(err6) {
		t.Error("rubyBoiler() should create .rspec file in the project dir but it failed")
	}
	if os.IsNotExist(err7) {
		t.Error("rubyBoiler() should create .rubocop.yml file in the project dir but it failed")
	}

	rmPrStr := "rm -r " + currDir + "/" + projectName
	argsRmPrStr := strings.Split(rmPrStr, " ")
	exec.Command(argsRmPrStr[0], argsRmPrStr[1:]...).Run()
}
