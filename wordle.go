package main

import (
	"github.com/ttacon/chalk"
	"github.com/wordle/conf"
	"github.com/wordle/wordler"
)

var green = chalk.Green.NewStyle().WithBackground(chalk.Black).Style

func main() {
	worlder := wordler.NewWordlerWithConf(conf.GetDefaultConfig())
	worlder.Start()
}
