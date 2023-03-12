package main

var filesToIgnore = []string{".hvc", ".hvcignore", ".hvcmodules"}

const rootPath = "./.hvc"

const stageFileName = rootPath + "/staged"

const commitFileName = rootPath + "/comments"

const objectsFolder = rootPath + "/objects"

const addFilesFolder = objectsFolder + "/files"

const defaultFilePermission = 0777
