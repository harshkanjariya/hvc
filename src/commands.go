package main

import "os"

func add(path string) {
	stat, err := os.Stat(path)
	if os.IsNotExist(err) {
		print("no such file or directory ", path)
	} else if err != nil {
		print(err)
	} else {
		if stat.IsDir() {
			
		} else {
			addFileToStage(path)
		}
	}
}

func commit(message string) {
	print("committing", message)
}
