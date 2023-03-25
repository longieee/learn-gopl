# learn-gopl

My implementation of exercises in book The Go Programming Language

# Note to myself

When working on VSCode in local, if I open the root folder of this repo, then the compiler will throw some errors, saying that it cannot find the packgages, e.g.

```go
import (
 "chap-2/tempconv" // This line will be shown by the linter that go cannot find the package
    [...]
```

To solve this problem, when working on a specific chapter (e.g. Chapter 2):

- I open the folder `chap-2` on VSCode.
- Create `go.work` file at `chap-2/go.work`
- Use `go work use` command to use the packages that will be imported, example: `go work use tempconv`
- Inside folder for packages that will be imported to other modules, use `go mod init` inside that folder. E.g. `cd tempconv` and then `go mod init chap-2/tempconv`
