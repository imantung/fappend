# file_append

Append text to file
Altenative to `echo "sometext" >> file.txt` 
May useful where in case [i/o redirection](https://tldp.org/LDP/intro-linux/html/sect_05_01.html) not working like in [go generate](https://go.dev/blog/generate).

## Install

```
go install github.com/imantung/file_append
```

You may need to add `GOPATH/bin` to `PATH` ([learn more](https://stackoverflow.com/questions/36083542/error-command-not-found-after-installing-go-eval))

## Usage

```
$ file_append file.txt sometext
```

Example for go generate

```go
package main

import "fmt"

// Generated file path relative to go file. Using $PROJ to create file in project directory

//go:generate rm -f $PROJ/.envrc
//go:generate file_append $PROJ/.envrc export APP_ADDRESS=:8089
//go:generate file_append $PROJ/.envrc export APP_READ_TIMEOUT=5s
//go:generate file_append $PROJ/.envrc export APP_WRITE_TIMEOUT=10s
//go:generate file_append $PROJ/.envrc export APP_DEBUG=true

func main() {
    fmt.Println("hello world")
}

```

```
$ go install github.com/imantung/file_append
$ PROJ=$(pwd) go generate ./...
```