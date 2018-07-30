package main

import (
	"bufio"
	"os"
	"strings"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

var isSpecialMode = walk.NewMutableCondition()

type MyMainWindow struct {
	*walk.MainWindow
}

func main() {

	mw := new(MyMainWindow)

	var inTe *walk.TextEdit

	var openAction, showAboutBoxAction *walk.Action
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
						Text:        "파일 추가",
						Enabled:     Bind("enabledCB.Checked"),
						Visible:     Bind("!openHiddenCB.Checked"),
						OnTriggered: mw.fileUploadAction_Triggered,
					},
					Menu{
						AssignTo: &recentMenu,
						Text:     "최근 파일",
						//OnTriggered: mw.recentFileAction_Triggered,
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

		ContextMenuItems: []MenuItem{
			ActionRef{&showAboutBoxAction},
		},

		MinSize: Size{270, 150},
		Layout:  VBox{},
		Children: []Widget{

			TextEdit{
				AssignTo: &inTe, ReadOnly: true},

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

	addRecentFileActions := func(texts ...string) {

		for _, text := range texts {
			a := walk.NewAction()
			a.SetText(text)
			a.Triggered().Attach(func() {

				day, name, count, countNoBlank := SplitTextDay(a.Text())

				inTe.SetText(day + name + count + countNoBlank)

			})
			recentMenu.Actions().Add(a)

		}
	}

	f, _ := os.Open("C:\\temp\\magamDday.txt")
	scanner := bufio.NewScanner(f)

	defer f.Close()

	for scanner.Scan() {
		line := scanner.Text()

		// Split the line on commas.
		parts := strings.Split(line, ";")

		// Loop over the parts from the string.
		for i := range parts {
			if len(parts[i]) > 1 {
				addRecentFileActions(parts[i])
			}
		}

	}

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
