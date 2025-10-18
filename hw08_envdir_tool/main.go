package main

import (
	"fmt"
	"os"

	"github.com/KorolevITCube/otus_go/hw08_envdir_tool/env"      //nolint:all
	"github.com/KorolevITCube/otus_go/hw08_envdir_tool/executor" //nolint:all
)

func main() {
	if len(os.Args) > 2 {
		envs, err := env.ReadDir(os.Args[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(-1)
		}
		exitCode := executor.RunCmd(os.Args[2:], envs)
		os.Exit(exitCode)
	}
	fmt.Fprintln(os.Stderr, "no arguments provided")
	os.Exit(-1)
}
