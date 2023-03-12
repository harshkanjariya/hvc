package main

import (
	"os"
)

func stringInArray(s string, arr []string) bool {
	for _, item := range arr {
		if item == s {
			return true
		}
	}
	return false
}

func listFiles(path string) []string {
	stat, err := os.Stat(path)
	var list []string
	check(err)
	if stringInArray(path, filesToIgnore) {
		return list
	}
	if stat.IsDir() {
		files, err := os.ReadDir(path)
		check(err)
		for _, file := range files {
			var newList = listFiles(path + "/" + file.Name())
			list = append(list, newList...)
		}
	} else {
		list = append(list, path)
	}
	return list
}
