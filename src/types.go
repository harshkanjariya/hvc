package main

type Person struct {
	firstName string
	lastName  string
	email     string
}

type CommitFileAction string

const (
	CommitFileActionAdd    CommitFileAction = "a"
	CommitFileActionChange CommitFileAction = "c"
	CommitFileActionRename CommitFileAction = "r"
	CommitFileActionDelete CommitFileAction = "d"
)

type CommitFile struct {
	path   string
	action CommitFileAction
}

type Commit struct {
	hash      string
	message   string
	committer string
	timestamp string
	parent    string
	files     []CommitFile
}
