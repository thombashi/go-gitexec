# go-gitexec

Thin wrapper to execute git commands in Go.

[![Go Reference](https://pkg.go.dev/badge/github.com/thombashi/go-gitexec.svg)](https://pkg.go.dev/github.com/thombashi/go-gitexec)
[![Go Report Card](https://goreportcard.com/badge/github.com/thombashi/go-gitexec)](https://goreportcard.com/report/github.com/thombashi/go-gitexec)
[![CI](https://github.com/thombashi/go-gitexec/actions/workflows/ci.yaml/badge.svg)](https://github.com/thombashi/go-gitexec/actions/workflows/ci.yaml)
[![CodeQL](https://github.com/thombashi/go-gitexec/actions/workflows/github-code-scanning/codeql/badge.svg)](https://github.com/thombashi/go-gitexec/actions/workflows/github-code-scanning/codeql)


## Usage

### Basic Usage

```go
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

	fmt.Printf(result.Stdout.String())
}
```

### Usage: with environment variables

```go
import (
    "fmt"
    "os"
    "strings"

    "github.com/thombashi/go-gitexec"
)

func main() {
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

	fmt.Printf(result.Stdout.String())
}
```
