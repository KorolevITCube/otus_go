package executor

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/KorolevITCube/otus_go/hw08_envdir_tool/env" //nolint:all
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(commands []string, env env.Environment) (returnCode int) {
	if len(commands) == 0 {
		fmt.Fprintln(os.Stderr, "no command to execute")
		return -1
	}
	var cmd *exec.Cmd
	if len(commands) == 1 {
		cmd = exec.Command(commands[0]) //nolint:gosec // по заданию подразумевает запуск всего
	} else {
		cmd = exec.Command(commands[0], commands[1:]...) //nolint:gosec // по заданию подразумевает запуск всего
	}
	cmd.Env = generateEnvs(os.Environ(), env)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			return exitErr.ExitCode()
		}
		fmt.Fprintln(os.Stderr, err.Error())
		return -1
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

	res := make([]string, 0, len(temp))
	for k, v := range temp {
		res = append(res, fmt.Sprintf("%s=%s", k, v))
	}
	return res
}
