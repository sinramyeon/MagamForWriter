package main

import (
	"docx"
	"encoding/json"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/lxn/walk"
)

const ConfFilePath = "C:\\temp\\conf.json"

type Configuration struct {
	Dday     string `json:"Dday"`
	Filename string `json:"Filename"`
}

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

func TxtFileOpen(filepath string) string {
	file, err := os.Open(filepath)
	if err != nil {
		return err.Error()
	}
	defer file.Close()

	fi, err := file.Stat()
	if err != nil {
		return err.Error()
	}

	var data = make([]byte, fi.Size())

	_, err = file.Read(data)
	if err != nil {
		return err.Error()
	}

	return string(data)
}

func DocFileOpen(filepath string) string {
	r, err := docx.ReadDocxFile(filepath)
	if err != nil {

		return err.Error()
	}
	docx := r.Editable()
	wholeText := docx.GetText()

	reg, err := regexp.Compile("<[^>]*>")
	if err != nil {
		return err.Error()
	}
	processedString := reg.ReplaceAllString(wholeText, "")
	r.Close()

	return processedString

}

func SaveFile(day, filepath string) error {

	configuration := Configuration{
		Dday:     day,
		Filename: filepath,
	}

	confJson, _ := json.Marshal(configuration)

	f, err := os.OpenFile(ConfFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	if _, err := f.Write(confJson); err != nil {
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}

	return err
}

func GetAlarmText() (string, string, string, string) {

	var newFile TxtFile
	txt := GetFile()
	filearray := strings.Split(txt, ";") //2018-06-20 C:\windows-version.txt;

	for i := range filearray {
		oneFile := strings.Split(filearray[i], " ")
		if len(oneFile) > 0 {

			newFile.dday = oneFile[0]
			newFile.name = oneFile[1]

			str := TxtFileOpen(newFile.name)
			count := CountAll(str)
			countWithoutBlank := CountRemoveBlank(str)

			return newFile.dday, newFile.name, strconv.Itoa(count), strconv.Itoa(countWithoutBlank)

		}
	}
	return "", "", "", ""
}

func GetFile() string {
	// 1. 파일 가져오기
	var file, err = ioutil.ReadFile(ConfFilePath)
	if err != nil {
		WalkError(err)
	}
	return string(file)
}

func WalkError(err error) {
	walk.MsgBox(
		nil,
		"Error",
		err.Error(),
		walk.MsgBoxOK|walk.MsgBoxIconError)
}

func GetDDay(day string) int {

	now := time.Now().Format("2006-01-02")

	walk.MsgBox(
		nil,
		"GetDDay",
		day,
		walk.MsgBoxOK|walk.MsgBoxIconError)

	t, err := time.Parse(now, day)

	if err != nil {

		walk.MsgBox(
			nil,
			"Error",
			err.Error(),
			walk.MsgBoxOK|walk.MsgBoxIconError)

		return 0

	}

	rightnow := time.Now()
	days := rightnow.Sub(t)

	return int(days.Hours() / 24)
}

func FloatToString(input_num float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(input_num, 'f', 6, 64)
}

func CountFile(name string) (string, string) {
	var str string

	if strings.Contains(name, "txt") {
		str = TxtFileOpen(name)
	}
	if strings.Contains(name, "doc") || strings.Contains(name, "docx") {
		str = DocFileOpen(name)
	}

	count := CountAll(str)
	countWithoutBlank := CountRemoveBlank(str)

	return strconv.Itoa(count), strconv.Itoa(countWithoutBlank)

}
