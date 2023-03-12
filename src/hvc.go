package main

import (
	"os"
)

func main() {
	args := os.Args[1:]
	switch args[0] {
	case "add":
		add(args[1])
		break
	case "commit":
		commit(args[1])
		break
	}
}
