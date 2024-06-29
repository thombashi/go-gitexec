package gitexec

import (
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRunGit(t *testing.T) {
	a := assert.New(t)
	r := require.New(t)

	executor, err := New(&Params{})
	r.NoError(err)

	// Test case 1: RunGit with valid arguments
	result, err := executor.RunGit("status")
	r.NoError(err)
	a.NotNil(result)
	a.Equal(0, result.ExitCode)
	a.NotEmpty(result.Stdout.String())
	a.Empty(result.Stderr.String())

	// Test case 2: RunGit with invalid arguments
	_, err = executor.RunGit("invalid-command")
	a.Error(err)
}

func TestRunGitContext(t *testing.T) {
	a := assert.New(t)
	r := require.New(t)
	ctx := context.Background()

	executor, err := New(&Params{})
	r.NoError(err)

	// Test case 1: RunGit with valid arguments
	result, err := executor.RunGitContext(ctx, "status")
	r.NoError(err)
	a.NotNil(result)
	a.Equal(0, result.ExitCode)
	a.NotEmpty(result.Stdout.String())
	a.Empty(result.Stderr.String())

	// Test case 2: RunGit with invalid arguments
	_, err = executor.RunGit("invalid-command")
	a.Error(err)
}

func TestRunGitWithEnv(t *testing.T) {
	a := assert.New(t)
	r := require.New(t)

	executor, err := New(&Params{
		Env: []string{
			"GIT_AUTHOR_NAME=John Doe",
			"GIT_AUTHOR_EMAIL=johndoe@example.com",
		},
	})
	r.NoError(err)

	result, err := executor.RunGit("var", "GIT_AUTHOR_IDENT")
	r.NoError(err)
	a.NotNil(result)
	a.Equal(0, result.ExitCode)
	a.True(strings.HasPrefix(result.Stdout.String(), "John Doe <johndoe@example.com>"))
	a.Empty(result.Stderr.String())
}
