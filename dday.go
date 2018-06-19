package main

import (
	"os"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func DdaySet() {
	var Day *walk.TextEdit
	MainWindow{
		Title:   "마감일을 입력하세요",
		MinSize: Size{600, 600},
		Layout:  VBox{},
		Children: []Widget{
			HSplitter{
				Children: []Widget{
					Label{
						Text: "마감일을 2018-06-06 형태로 입력해 주세요/",
					},
					TextEdit{AssignTo: &Day},

					PushButton{
						Text: "마감일 등록",
						OnClicked: func() {
							var file, _ = os.OpenFile("C:\\temp\\test.txt", os.O_RDWR|os.O_CREATE, 0644)
							defer file.Close()
							file.WriteString(Day.Text())

							file.Sync()
						},
					},
				},
			},
		},
	}.Run()

}
