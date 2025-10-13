package executor

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/KorolevITCube/otus_go/hw08_envdir_tool/env"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(commands []string, env env.Environment) (returnCode int) {
	if len(commands) == 0 {
		fmt.Fprintln(os.Stderr, "no command to execute")
		return -1
	}
	var cmd *exec.Cmd
	if len(commands) == 1 {
		cmd = exec.Command(commands[0])
	} else {
		cmd = exec.Command(commands[0], commands[1:]...)
	}
	cmd.Env = generateEnvs(os.Environ(), env)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			return exitErr.ExitCode()
		} else {
			fmt.Fprintln(os.Stderr, err.Error())
			return -1
		}
	}
	return 0
}

func generateEnvs(rootEnvs []string, localEnvs env.Environment) []string {
	temp := make(map[string]string)
	for _, v := range rootEnvs {
		parsed := strings.Split(v, "=")
		temp[parsed[0]] = parsed[1]
	}

	for k, v := range localEnvs {
		if v.NeedRemove {
			delete(temp, k)
		} else {
			temp[k] = v.Value
		}
	}

	var res []string
	for k, v := range temp {
		res = append(res, fmt.Sprintf("%s=%s", k, v))
	}
	return res
}
