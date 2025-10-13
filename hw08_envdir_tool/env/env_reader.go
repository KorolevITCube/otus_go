package env

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	envs := make(Environment)
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("error while opening directory: %w", err)
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			fi, err := os.Stat(fmt.Sprintf("%s/%s", dir, entry.Name()))
			if err != nil {
				return nil, fmt.Errorf("error while getting file stat: %w", err)
			}

			if fi.Size() == 0 {
				envs[entry.Name()] = EnvValue{
					Value:      "",
					NeedRemove: true,
				}
			} else {
				file, err := os.Open(fmt.Sprintf("%s/%s", dir, entry.Name()))
				if err != nil {
					return nil, fmt.Errorf("error while opening file: %w", err)
				}
				defer file.Close()

				scanner := bufio.NewScanner(file)
				if scanner.Scan() {
					firstLine := scanner.Text()
					firstLine = strings.TrimRight(firstLine, " \t")
					firstLine = string(bytes.ReplaceAll([]byte(firstLine), []byte{0x00}, []byte("\n")))
					envs[entry.Name()] = EnvValue{
						Value:      firstLine,
						NeedRemove: false,
					}
				}

				if err := scanner.Err(); err != nil {
					return nil, fmt.Errorf("error while reading file: %w", err)
				}
			}
		}
	}
	return envs, nil
}
