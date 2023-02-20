package main

import (
	"github.com/GLVSKiriti/studybuddy/cmd"
	"github.com/GLVSKiriti/studybuddy/data"
)

func main() {
	data.OpenDatabase()
	cmd.Execute()
}
