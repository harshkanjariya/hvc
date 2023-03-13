package main

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
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

func validateFolder(path string) {
	_, errStat := os.Stat(path)
	if os.IsNotExist(errStat) {
		err := os.MkdirAll(path, defaultFilePermission)
		check(err)
	}
}

func validateFile(path string) {
	_, errStat := os.Stat(path)
	if os.IsNotExist(errStat) {
		_, err := os.Create(path)
		check(err)
	}
}

func validateHVCTree() {
	validateFile(stageFileName)
	validateFile(commitFileName)
	validateFolder(objectsFolder)
	validateFolder(addFilesFolder)

	validateFile(headFileName)
	validateFolder(headsFolder)
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

func getConfig() map[string]string {
	ex, err2 := os.Executable()
	check(err2)
	exPath := filepath.Dir(ex)
	content, err := os.ReadFile(exPath + "/" + configFilePath)
	check(err)

	var conf = make(map[string]string)

	for _, line := range strings.Split(string(content), "\n") {
		if len(line) > 0 && strings.Contains(line, "=") {
			parts := strings.Split(line, "=")
			conf[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
		}
	}

	_, exists := conf["defaultBranch"]
	if !exists {
		conf["defaultBranch"] = "master"
	}

	return conf
}

func updateHead(content string) {
	head, err := os.ReadFile(headFileName)
	check(err)
	headStr := strings.TrimSpace(string(head))
	if strings.Contains(headStr, ":") {
		headStr = strings.Split(headStr, ":")[1]
	} else {
		// TODO
	}
	err = os.WriteFile(headsFolder+"/"+headStr, []byte(content), defaultFilePermission)
	check(err)
}
func getHead() string {
	head, err := os.ReadFile(headFileName)
	check(err)
	headStr := strings.TrimSpace(string(head))
	if strings.Contains(headStr, ":") {
		headStr = strings.Split(headStr, ":")[1]
	} else {
		return headStr
	}
	target, err2 := os.ReadFile(headsFolder + "/" + headStr)
	check(err2)
	return strings.TrimSpace(string(target))
}

func test() {
}
