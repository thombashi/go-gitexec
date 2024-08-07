package gitexec

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os/exec"
	"strings"

	"github.com/cli/safeexec"
)

// GitExecutor is an interface to execute Git commands.
type GitExecutor interface {
	// RunGit executes a Git command with the specified arguments.
	RunGit(args ...string) (*CmdResult, error)

	// RunGitContext executes a Git command with the specified context and arguments.
	RunGitContext(ctx context.Context, args ...string) (*CmdResult, error)

	// WithLogger sets the logger to use for logging.
	WithLogger(logger *slog.Logger) GitExecutor

	// GetLogger returns the logger instance.
	GetLogger() *slog.Logger
}

type gitExecutorImpl struct {
	logger  *slog.Logger
	gitPath string
	env     []string
}

// CmdResult represents the result of a Git command execution.
type CmdResult struct {
	// Stdout contains the standard output of the Git command (if any).
	Stdout *bytes.Buffer

	// Stderr contains the standard error of the Git command (if any).
	Stderr *bytes.Buffer

	// ExitCode is the exit code of the Git command.
	ExitCode int
}

// Params is the parameters for creating a new GitExecutor instance.
type Params struct {
	// Logger is the logger to use for logging.
	Logger *slog.Logger

	// GitPath is the path to the Git executable. If empty, the Git executable is searched in the PATH.
	GitPath string

	// Env specifies the environment of the process.
	// Each entry is of the form "key=value".
	// If Env is nil, the new process uses the current process's
	// environment.
	// If Env contains duplicate environment keys, only the last
	// value in the slice for each duplicate key is used.
	// As a special case on Windows, SYSTEMROOT is always added if
	// missing and not explicitly set to the empty string.
	Env []string
}

// New creates a new GitExecutor instance.
func New(params *Params) (GitExecutor, error) {
	var err error
	gitPath := params.GitPath

	if gitPath == "" {
		gitPath, err = safeexec.LookPath("git")
		if err != nil {
			return nil, fmt.Errorf("failed to find git: %w", err)
		}
	}

	return &gitExecutorImpl{
		logger:  params.Logger,
		gitPath: gitPath,
		env:     params.Env,
	}, nil
}

// RunGit executes a Git command with the specified arguments.
func (e gitExecutorImpl) RunGit(args ...string) (*CmdResult, error) {
	return e.RunGitContext(context.Background(), args...)
}

// RunGitContext executes a Git command with the specified context and arguments.
func (e gitExecutorImpl) RunGitContext(ctx context.Context, args ...string) (*CmdResult, error) {
	if e.logger != nil {
		e.logger.Debug("execute", slog.String("command", fmt.Sprintf("git %s", strings.Join(args, " "))))
	}

	var stdout, stderr bytes.Buffer
	gitCmd := exec.CommandContext(ctx, e.gitPath, args...)
	gitCmd.Stdout = &stdout
	gitCmd.Stderr = &stderr
	gitCmd.Env = e.env

	err := gitCmd.Run()
	result := &CmdResult{
		Stdout: &stdout,
		Stderr: &stderr,
	}
	if err != nil {
		var execError *exec.ExitError
		if errors.As(err, &execError) {
			result.ExitCode = execError.ExitCode()
		} else {
			return nil, fmt.Errorf("failed to execute git command: %w", err)
		}
	}

	return result, err
}

// WithLogger sets the logger to use for logging.
func (e gitExecutorImpl) WithLogger(logger *slog.Logger) GitExecutor {
	e.logger = logger
	return e
}

// GetLogger returns the logger instance.
func (e gitExecutorImpl) GetLogger() *slog.Logger {
	return e.logger
}
