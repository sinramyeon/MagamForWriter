package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func main() {

	var mainWindow *walk.MainWindow

	MainWindow{
		AssignTo: &mainWindow,
		Title:    "마감 안내기",
		MinSize:  Size{100, 400},
		Layout:   VBox{},
		Children: []Widget{

			PushButton{
				Text: "작업 등록하기",
				OnClicked: func() {
					mainWindow.Close()
					Fileupload()
				},
			},
			PushButton{
				Text: "마감일 안내받기",
				OnClicked: func() {
					//mainWindow.Close()
					day, name, count := GetAlarmText()
					Alarm(day, name, count)
				},
			},
		},
	}.Run()
}
