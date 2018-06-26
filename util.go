package main

import "strings"

func GetFilename(filePath string) string {
	i := strings.Index(filePath, "\\")
	if i > -1 {
		fileName := filePath[i+1:]
		if strings.ContainsAny(fileName, "\\") {
			GetFilename(fileName)
		}
		return fileName
	} else {
		return filePath
	}

}
