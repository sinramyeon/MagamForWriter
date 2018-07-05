package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

var isSpecialMode = walk.NewMutableCondition()

type MyMainWindow struct {
	*walk.MainWindow
}

func main() {

	mw := new(MyMainWindow)

	var openAction, showAboutBoxAction, fileUploadAction *walk.Action
	var recentMenu *walk.Menu

	if err := (MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    "마감 안내기",
		MenuItems: []MenuItem{
			Menu{
				Text: "파일 업로드",
				Items: []MenuItem{
					Action{
						AssignTo:    &openAction,
						Text:        "&File",
						Enabled:     Bind("enabledCB.Checked"),
						Visible:     Bind("!openHiddenCB.Checked"),
						OnTriggered: mw.fileUploadAction_Triggered,
					},
					Menu{
						AssignTo: &recentMenu,
						Text:     "최근 파일",
					},
					Separator{},
					Action{
						Text:        "종료",
						OnTriggered: func() { mw.Close() },
					},
				},
			},

			Menu{
				Text: "&Help",
				Items: []MenuItem{
					Action{
						AssignTo:    &showAboutBoxAction,
						Text:        "About",
						OnTriggered: mw.showAboutBoxAction_Triggered,
					},
				},
			},
		},

		ToolBar: ToolBar{
			ButtonStyle: ToolBarButtonImageBeforeText,
			Items: []MenuItem{
				Action{
					AssignTo:    &fileUploadAction,
					Text:        "파일 추가",
					OnTriggered: mw.fileUploadAction_Triggered,
				},
				Separator{},
				Menu{
					Text: "설정",
					Items: []MenuItem{
						Action{
							Text:        "디자인 바꾸기",
							OnTriggered: mw.colourAction_Triggered,
						},
					},
				},
			},
		},

		ContextMenuItems: []MenuItem{
			ActionRef{&showAboutBoxAction},
		},

		MinSize: Size{270, 150},
		Layout:  VBox{},
		Children: []Widget{
			PushButton{
				Text: "마감일 안내받기",
				OnClicked: func() {
					day, name, count, countWithoutBlank := GetAlarmText()
					Alarm(day, name, count, countWithoutBlank)
				},
			},
		},
	}.Create()); err != nil {
		walk.MsgBox(mw, "err", err.Error(), walk.MsgBoxIconInformation)
	}

	// txtFilesName := GetTextNameFromConf()

	// addRecentFileActions := func(texts []string) {

	// 	walk.MsgBox(
	// 		nil,
	// 		"addRecentFileActions",
	// 		"addRecentFileActions",
	// 		walk.MsgBoxIconInformation)

	// 	for _, text := range texts {
	// 		a := walk.NewAction()
	// 		a.SetText(text)
	// 		//a.Triggered().Attach(mw.openAction_Triggered)
	// 		recentMenu.Actions().Add(a)
	// 	}
	// }

	// addRecentFileActions(txtFilesName)

	mw.Run()
}

func (mw *MyMainWindow) showAboutBoxAction_Triggered() {
	walk.MsgBox(mw, "About", `글 쓰시는 분들의 마감을 도와드립니다.
			20180703 히어로
				github @hero0926
		`, walk.MsgBoxIconInformation)
}

func (mw *MyMainWindow) fileUploadAction_Triggered() {
	Fileupload()
}

func (mw *MyMainWindow) colourAction_Triggered() {
	ColourSetting()
}
