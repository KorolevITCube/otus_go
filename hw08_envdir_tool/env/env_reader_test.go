package env

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	t.Run("Read all variables", func(t *testing.T) {
		envs, err := ReadDir("./testdata/env")
		want := make(Environment)
		want["BAR"] = EnvValue{Value: "bar", NeedRemove: false}
		want["EMPTY"] = EnvValue{Value: "", NeedRemove: false}
		want["FOO"] = EnvValue{Value: "   foo\nwith new line", NeedRemove: false}
		want["HELLO"] = EnvValue{Value: "\"hello\"", NeedRemove: false}
		want["UNSET"] = EnvValue{Value: "", NeedRemove: true}

		require.Nil(t, err)
		require.Equal(t, want, envs)
	})

	t.Run("Not a directory", func(t *testing.T) {
		_, err := ReadDir("./testdata/env/BAR")
		require.NotNil(t, err)
	})

	t.Run("No such directory", func(t *testing.T) {
		_, err := ReadDir("./testdata/env/BAR/BAR_BAR?")
		require.NotNil(t, err)
	})
}
