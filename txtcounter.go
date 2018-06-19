package main

import (
	. "github.com/lxn/walk/declarative"
)

func main() {

	MainWindow{
		Title:   "마감하세요",
		MinSize: Size{600, 400},
		Layout:  VBox{},
		Children: []Widget{
			HSplitter{
				PushButton{
					Text: "작업파일 등록하기",
					OnClicked: func() {
						Fileupload()
					},
				},

				PushButton{
					Text: "마감현황 보기",
					OnClicked: func() {
					},
				},
			},
		},
	}.Run()
}
