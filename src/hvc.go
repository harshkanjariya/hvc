package main

import (
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		printMenu()
		return
	}
	switch args[0] {
	case "add":
		if len(args) < 2 {
			printMenu()
		} else {
			add(args[1])
		}
		break
	case "commit":
		if len(args) < 2 {
			printMenu()
		} else {
			commit(args[1])
		}
		break
	default:
		printMenu()
	}
}
