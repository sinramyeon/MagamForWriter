package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func Alarm(day, name, count, countWithoutBlank string) {

	var mainWindow *walk.MainWindow

	MainWindow{
		AssignTo: &mainWindow,
		Title:    "마감 안내기",
		MinSize:  Size{200, 50},
		Layout:   VBox{},
		Children: []Widget{

			Label{
				Text: "글이름",
			},
			Label{
				Text: GetFilename(name),
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
			Label{
				Text: "공백 제거 글자수",
			},
			Label{
				Text: countWithoutBlank,
			},
		},
	}.Run()

}
