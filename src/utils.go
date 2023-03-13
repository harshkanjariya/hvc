package main

import (
	"crypto/sha256"
	"encoding/base64"
	"os"
	"time"
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

func check(e error) {
	if e != nil {
		panic(e)
	}
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

func generateHash(str string) string {
	hasher := sha256.New()
	hasher.Write([]byte(str))
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return sha
}

func getCurrentUser() string {
	return "Harsh Kanjariya"
}
