package executor

import (
	"testing"

	"github.com/KorolevITCube/otus_go/hw08_envdir_tool/env"

	"github.com/stretchr/testify/require"
)

func TestRunCmd(t *testing.T) {
	t.Run("Test RunCmd Success", func(t *testing.T) {
		env := env.Environment{"TEST_VAR": env.EnvValue{Value: "123", NeedRemove: false}}
		returnCode := RunCmd([]string{"sh", "-c", "echo $TEST_VAR"}, env)
		require.Equal(t, 0, returnCode)
	})

	t.Run("Test RunCmd NonexistentCommand", func(t *testing.T) {
		env := env.Environment{}
		returnCode := RunCmd([]string{"nonexistentcmd"}, env)
		require.NotEqual(t, 0, returnCode)
	})

	t.Run("Test RunCmd EmptyCmd", func(t *testing.T) {
		env := env.Environment{}
		returnCode := RunCmd([]string{}, env)
		require.NotEqual(t, 0, returnCode)
	})

}
