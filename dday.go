package main

import (
	"regexp"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func (txtFile *TxtFile) DdaySet() {
	var Day *walk.TextEdit
	var mainWindow *walk.MainWindow

	MainWindow{
		AssignTo: &mainWindow,
		Title:    "마감일을 입력하세요",
		MinSize:  Size{800, 20},
		Layout:   VBox{},
		Children: []Widget{
			HSplitter{
				Children: []Widget{
					Label{
						Text: "마감일을 2018-06-06 형태로 입력해 주세요",
					},
					TextEdit{AssignTo: &Day},

					PushButton{
						Text: "마감일 등록",
						OnClicked: func() {
							txtFile.dday = Day.Text()
							re := regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)

							if !re.MatchString(txtFile.dday) {
								walk.MsgBox(
									nil,
									"날짜 에러",
									"날짜 형식을 2018-07-05 에 맞춰 입력하세요.",
									walk.MsgBoxOK|walk.MsgBoxIconError)
								return
							}

							mainWindow.Close()
						},
					},
				},
			},
		},
	}.Run()

}
