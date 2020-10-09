package languages

import (
	"github.com/descholar-ceo/boiler/utils"
)

// RorBoiler is a boilerplate generator for ror
func RorBoiler() {
	workingDir := utils.AskWorkingDirectory()
	isGithub := utils.AskGithub()
	isRubocop := utils.AskRubocop()
	database := utils.AskDatabase()
}
