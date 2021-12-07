package command

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInitCommand(t *testing.T) {
	cwd, err := os.Getwd()
	require.NoError(t, err, "failed to get working dir")

	err = os.Chdir(t.TempDir())
	require.NoError(t, err, "failed to change dir")
	t.Cleanup(func() { os.Chdir(cwd) })

	err = Init(context.Background(), "valist/sdk", false)
	require.NoError(t, err, "failed to init valist project")

	err = Init(context.Background(), "valist/sdk", false)
	assert.ErrorIs(t, err, ErrProjectExist)
}