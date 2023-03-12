package main

import (
	"os"
)

func add(path string) {
	stat, err := os.Stat(path)
	if os.IsNotExist(err) {
		print("no such file or directory ", path)
		return
	} else {
		check(err)
	}
	if stat.IsDir() {
		files, err := os.ReadDir(path)
		check(err)
		for _, file := range files {
			if stringInArray(file.Name(), filesToIgnore) {
				continue
			}
			if file.IsDir() {
				check(err)
				var list = listFiles(path + "/" + file.Name())
				for _, fileName := range list {
					addFileToStage(fileName)
				}
			} else {
				addFileToStage(path + "/" + file.Name())
			}
		}
	} else {
		addFileToStage(path)
	}
}

func commit(message string) {
	print("committing", message)
}
