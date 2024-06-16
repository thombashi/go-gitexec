package gitexec_test

import (
	"fmt"
	"strings"

	"github.com/thombashi/go-gitexec"
)

func ExampleRunGit() {
	executor, err := gitexec.New(&gitexec.Params{})
	if err != nil {
		panic(err)
	}

	result, err := executor.RunGit("status")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(result.Stdout.String(), "\n")
	fmt.Printf("exit code: %d\n", result.ExitCode)
	fmt.Printf("stdout: %s\n", lines[0])
	// Output:
	// exit code: 0
	// stdout: On branch main
}

func ExampleRunGit_withEnv() {
	executor, err := gitexec.New(&gitexec.Params{
		Env: []string{
			"GIT_AUTHOR_NAME=John Doe",
			"GIT_AUTHOR_EMAIL=johndoe@example.com",
		},
	})
	if err != nil {
		panic(err)
	}

	result, err := executor.RunGit("var", "GIT_AUTHOR_IDENT")
	if err != nil {
		panic(err)
	}

	items := strings.Split(result.Stdout.String(), " ")
	fmt.Printf("exit code: %d\n", result.ExitCode)
	fmt.Printf("stdout: %s\n", items[:3])
	// Output:
	// exit code: 0
	// stdout: [John Doe <johndoe@example.com>]
}
