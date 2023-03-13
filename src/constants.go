package main

var filesToIgnore = []string{".hvc", ".hvcignore", ".hvcmodules"}

const binaryContentSeparator = "<SEPARATOR>"

const rootPath = "./.hvc"

const configFilePath = ".hvcrc"

const headFileName = rootPath + "/HEAD"

const headsFolder = rootPath + "/heads"

const stageFileName = rootPath + "/staged"

const commitFileName = rootPath + "/comments"

const objectsFolder = rootPath + "/objects"

const addFilesFolder = objectsFolder + "/files"

const defaultFilePermission = 0777
