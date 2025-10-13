package env

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	t.Run("Read all variables", func(t *testing.T) {
		envs, err := ReadDir("../testdata/env")
		want := make(Environment)
		want["BAR"] = Value{Value: "bar", NeedRemove: false}
		want["EMPTY"] = Value{Value: "", NeedRemove: false}
		want["FOO"] = Value{Value: "   foo\nwith new line", NeedRemove: false}
		want["HELLO"] = Value{Value: "\"hello\"", NeedRemove: false}
		want["UNSET"] = Value{Value: "", NeedRemove: true}

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
