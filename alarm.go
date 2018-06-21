package main

import (
	"io/ioutil"
	"strings"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func getFile() string {
	// 1. 파일 가져오기
	var file, err = ioutil.ReadFile("C:\\temp\\magamDday.txt")

	if err != nil {
		walk.MsgBox(
			nil,
			"Error",
			err.Error(),
			walk.MsgBoxOK|walk.MsgBoxIconError)
	}

	return string(file)
}

func Alarm() {
	var day, name string
	txt := getFile()
	// 2. 글이름, 마감일 읽기
	filearray := strings.Split(txt, ";") //2018-06-20 C:\windows-version.txt;

	for i := range filearray {
		oneFile := strings.Split(filearray[i], " ")
		day, name = oneFile[0], oneFile[1]
	}
	// 3. 글이름, 마감일, 글자수 세기
	// * 글자수는 10분마다 새로 세야함
	count := CountAll(day)

	var mainWindow *walk.MainWindow

	MainWindow{
		AssignTo: &mainWindow,
		Title:    "마감 안내기",
		MinSize:  Size{100, 400},
		Layout:   VBox{},
		Children: []Widget{

			Label{
				Text: "글이름",
			},
			Label{
				Text: name,
			},
			Label{
				Text: "마감일",
			},
			Label{
				Text: day,
			},
			Label{
				Text: "글자수",
			},
			Label{
				Text: count,
			},
		},
	}.Run()

	// ticker := time.NewTicker(10 * time.Minute)
	// quit := make(chan struct{})
	// go func() {
	// 	for {
	// 		select {
	// 		case <-ticker.C:
	// 			// do stuff
	// 			txtfile, _ := ioutil.ReadFile(name)
	// 			count = CountChar(string(txtfile))
	// 		case <-quit:
	// 			ticker.Stop()
	// 			return
	// 		}
	// 	}
	// }()

}
