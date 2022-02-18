package methods

import "github.com/ttacon/chalk"

var ColorMap = map[JudgeInfo]func(string)string{
	InWordButIncorrectPosition: chalk.Yellow.NewStyle().WithBackground(chalk.Black).Style,
	InCorrectCompletely: chalk.Red.NewStyle().WithBackground(chalk.Black).Style,
	Correct: chalk.Green.NewStyle().WithBackground(chalk.Black).Style,
}
