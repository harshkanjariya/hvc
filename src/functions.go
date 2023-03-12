package main

import (
	"os"
	"strings"
)

func addFileToStage(path string) {
	_, errStat := os.Stat(stageFileName)
	if os.IsNotExist(errStat) {
		_, err := os.Create(stageFileName)
		check(err)
	}

	content, err := os.ReadFile(path)
	check(err)
	if len(content) == 0 {
		return
	}

	var compressed = compressData(content)

	stagedContent, err2 := os.ReadFile(stageFileName)
	check(err2)

	if len(stagedContent) == 0 {
		finalBytes := append([]byte("file:"+path+"\n"), compressed...)
		err = os.WriteFile(stageFileName, finalBytes, 0644)
		check(err)
	} else {
		var tree = make(map[string]string)
		res1 := strings.Split(string(stagedContent), "file:")
		for _, data := range res1[1:] {
			content := strings.SplitN(data, "\n", 2)
			tree[content[0]] = content[1]
		}
		tree[path] = string(compressed)

		var bytes = []byte{}
		for key, value := range tree {
			bytes = append(bytes, []byte("file:"+key+"\n")...)
			bytes = append(bytes, []byte(value)...)
		}
		err = os.WriteFile(stageFileName, bytes, 0644)
		check(err)
	}
}

func printMenu() {
	println("Available commands:")
	println("\tadd <file/folder>\t\t: Add file/folder to commit")
	println("\tcommit <message>\t\t: Commit the changes to repo")
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
