package main

import (
	"fmt"
	"os"
	"simple.com/main/color"
	"strconv"
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
	var tree = getStagedContent()
	if len(tree) == 0 {
		print("no files added to commit.")
		return
	}

	var commits = getCommits()
	if len(commits) == 0 {
		commit := generateCommit(message, nil)

		addContent := ""
		var files []CommitFile
		for file, content := range tree {
			addContent += strconv.Itoa(countRune(content, '\n')+1) + "\n"
			addContent += content + "\n"
			files = append(files, CommitFile{
				path:   file,
				action: CommitFileActionAdd,
			})
		}
		commit.files = files
		addFilePath := addFilesFolder + "/" + commit.hash[:2]
		if _, err := os.Stat(addFilePath); err != nil {
			if os.IsNotExist(err) {
				err := os.MkdirAll(addFilePath, defaultFilePermission)
				check(err)
			} else {
				check(err)
			}
		}
		addFilePath += "/" + commit.hash[2:]
		err := os.WriteFile(addFilePath, []byte(addContent), defaultFilePermission)
		check(err)
		commits = append(commits, commit)
	}
	saveCommits(commits)
}

func status() {
	stagedContent := getStagedContent()
	println("File changes added to stage")
	for file, _ := range stagedContent {
		println("\t" + color.Green + file)
	}
}

func log() {
	commits := getCommits()
	fmt.Println(commits)
}
