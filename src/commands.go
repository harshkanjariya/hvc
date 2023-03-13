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
	var newCommit Commit
	if len(commits) == 0 {
		newCommit = generateCommit(message, nil)

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
		newCommit.files = files
		addFilePath := addFilesFolder + "/" + newCommit.hash[:2]
		if _, err := os.Stat(addFilePath); err != nil {
			if os.IsNotExist(err) {
				err := os.MkdirAll(addFilePath, defaultFilePermission)
				check(err)
			} else {
				check(err)
			}
		}
		addFilePath += "/" + newCommit.hash[2:]
		err := os.WriteFile(addFilePath, []byte(addContent), defaultFilePermission)
		check(err)
		commits = append(commits, newCommit)
	} else {
		parentCommitHash := getHead()
		println(parentCommitHash)
	}
	saveCommits(commits)
	updateHead(newCommit.hash)
	err := os.WriteFile(stageFileName, []byte(""), defaultFilePermission)
	check(err)
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

func initialize() {
	stats, err := os.Stat(headFileName)
	if err == nil && stats.Size() > 0 {
		println("hvc already exists")
		return
	}
	validateHVCTree()
	config := getConfig()
	err = os.WriteFile(headFileName, []byte("h:"+config["defaultBranch"]), defaultFilePermission)
	check(err)
	_, err2 := os.Create(headsFolder + "/" + config["defaultBranch"])
	check(err2)
}
