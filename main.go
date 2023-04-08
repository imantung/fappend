package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Missing filename")
	}

	if len(os.Args) < 3 {
		log.Fatal("Missing text to append")
	}

	filename := os.Args[1]
	text := strings.Join(os.Args[2:], " ") + "\n"

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	if _, err = f.WriteString(text); err != nil {
		log.Fatal(err)
	}
}
