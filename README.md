# file_append

Append text to file
Altenative to `echo "sometext" >> file.txt` 
May useful where in case [i/o redirection](https://tldp.org/LDP/intro-linux/html/sect_05_01.html) not working like in [go generate](https://go.dev/blog/generate).

## Install

```
go install -u github.com/imantung/file_append
```

## Usage

```
$ file_append file.txt sometext
```