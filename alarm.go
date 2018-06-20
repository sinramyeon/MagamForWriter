package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func Alarm(name, day, count string) {

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
}
