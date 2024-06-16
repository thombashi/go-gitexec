# go-gitexec

Thin wrapper to execute git commands in Go.

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
