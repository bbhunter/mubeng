package main

import (
	"github.com/mubeng/mubeng/internal/runner"
	"github.com/projectdiscovery/gologger"
)

func main() {
	opt := runner.Options()

	if err := runner.New(opt); err != nil {
		gologger.Fatal().Msgf("Error! %s", err)
	}
}