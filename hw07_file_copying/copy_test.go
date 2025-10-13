package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	tests := []struct {
		name   string
		offset int
		limit  int
		from   string
		to     string
		err    error
	}{
		{
			name:   "Limit more then size",
			offset: 1000,
			limit:  0,
			from:   "./testdata/input.txt",
			to:     "./tmp/out.txt",
			err:    nil,
		},
		{
			name:   "Offset more then size",
			offset: 100000,
			limit:  0,
			from:   "./testdata/input.txt",
			to:     "./tmp/out.txt",
			err:    ErrOffsetExceedsFileSize,
		},
	}

	err := os.Mkdir("tmp", 0o755)
	if err != nil {
		require.FailNow(t, fmt.Errorf("Cant create temporary dir: %w", err).Error())
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			gotResult := Copy(tc.from, tc.to, int64(tc.offset), int64(tc.limit))
			require.Equal(t, tc.err, gotResult)
		})
	}

	err = os.RemoveAll("tmp")
	if err != nil {
		require.FailNow(t, fmt.Errorf("Cant delete temporary dir: %w", err).Error())
	}
}
