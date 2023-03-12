package main

import (
	"crypto/sha256"
	"encoding/base64"
	"os"
	"strconv"
	"strings"
	"time"
)

func addFileToStage(path string) {
	if path[:2] == "./" {
		path = path[2:]
	}
	content, err := os.ReadFile(path)
	check(err)
	if len(content) == 0 {
		return
	}

	var compressed = compressData(content)

	var stagedTree = getStagedContent()

	if len(stagedTree) == 0 {
		finalBytes := append([]byte("file:"+path+"\n"), compressed...)
		err = os.WriteFile(stageFileName, finalBytes, defaultFilePermission)
		check(err)
	} else {
		stagedTree[path] = string(compressed)

		var bytes []byte
		for key, value := range stagedTree {
			bytes = append(bytes, []byte("file:"+key+"\n")...)
			bytes = append(bytes, []byte(value)...)
		}
		err = os.WriteFile(stageFileName, bytes, defaultFilePermission)
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

func validateHVCTree() {
	_, errStat := os.Stat(stageFileName)
	if os.IsNotExist(errStat) {
		_, err := os.Create(stageFileName)
		check(err)
	}
	_, errStat = os.Stat(commitFileName)
	if os.IsNotExist(errStat) {
		_, err := os.Create(commitFileName)
		check(err)
	}
	_, errStat = os.Stat(objectsFolder)
	if os.IsNotExist(errStat) {
		err := os.MkdirAll(objectsFolder, defaultFilePermission)
		check(err)
	}
}

func getStagedContent() map[string]string {
	stagedContent, err2 := os.ReadFile(stageFileName)
	check(err2)

	var tree = make(map[string]string)
	if len(stagedContent) < 0 {
		return tree
	}
	res1 := strings.Split(string(stagedContent), "file:")
	for _, data := range res1[1:] {
		content := strings.SplitN(data, "\n", 2)
		tree[content[0]] = content[1]
	}
	return tree
}

func getCommits() []Commit {
	fileContent, err := os.ReadFile(commitFileName)
	check(err)

	var commits []Commit
	if len(fileContent) == 0 {
		return commits
	}

	var content = decompressData(fileContent)
	if len(content) == 0 {
		return commits
	}
	lines := strings.Split(string(content), "\n")

	var currentCommit *Commit = nil
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		if strings.Contains(line, ",") {
			if currentCommit != nil {
				commits = append(commits, *currentCommit)
			}
			parts := strings.SplitN(line, ",", 4)
			currentCommit = &Commit{
				hash:      parts[0],
				parent:    parts[1],
				timestamp: parts[2],
				message:   parts[3],
				files:     []CommitFile{},
			}
		} else {
			parts := strings.SplitN(line, ":", 2)
			currentCommit.files = append(currentCommit.files, CommitFile{
				path:   parts[0],
				action: CommitFileAction(parts[1]),
			})
		}
	}
	if currentCommit != nil {
		commits = append(commits, *currentCommit)
	}
	return commits
}

func saveCommits(commits []Commit) {
	content := ""
	const separator1 = ","
	const separator2 = ":"
	for _, commit := range commits {
		content += commit.hash + separator1 +
			commit.parent + separator1 +
			commit.timestamp + separator1 +
			commit.message + "\n"
		for _, file := range commit.files {
			content += file.path + separator2 + string(file.action) + "\n"
		}
	}
	err := os.WriteFile(commitFileName, compressData([]byte(content)), defaultFilePermission)
	check(err)
}

func generateCommit(message string, parentHash *string) Commit {
	var timestamp = currentMillis()
	var committer = getCurrentUser()
	hash := generateHash(committer + message + strconv.FormatInt(timestamp, 10))

	c := Commit{
		hash:      hash,
		message:   message,
		committer: committer,
		timestamp: strconv.FormatInt(timestamp, 10),
	}
	if parentHash != nil {
		c.parent = *parentHash
	}
	return c
}

func generateHash(str string) string {
	hasher := sha256.New()
	hasher.Write([]byte(str))
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return sha
}

func getCurrentUser() string {
	return "Harsh Kanjariya"
}

func currentMillis() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func countRune(s string, r rune) int {
	count := 0
	for _, c := range s {
		if c == r {
			count++
		}
	}
	return count
}
