package main

import (
	. "github.com/lxn/walk/declarative"
)

func main() {

	MainWindow{
		Title:   "마감 안내기",
		MinSize: Size{100, 400},
		Layout:  VBox{},
		Children: []Widget{

			PushButton{
				Text: "작업 등록하기",
				OnClicked: func() {
					Fileupload()
				},
			},
			PushButton{
				Text: "마감일 안내받기",
				OnClicked: func() {

				},
			},
		},
	}.Run()
}
