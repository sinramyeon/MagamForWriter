package main

import (
	"log"

	"github.com/lxn/walk"
)

func Alarm(day, name, count, countWithoutBlank string) {

	mw, err := walk.NewMainWindow()
	if err != nil {
		log.Fatal(err)
	}
	icon, err := walk.Resources.Icon("./file.ico")
	if err != nil {
		log.Fatal(err)
	}
	ni, err := walk.NewNotifyIcon()

	if err != nil {
		log.Fatal(err)
	} // 에러처리 할 미들웨어를 만들던지하자 이짓거리하지말고

	defer ni.Dispose()

	if err := ni.SetIcon(icon); err != nil {

		walk.MsgBox(
			nil,
			"Error",
			err.Error(),
			walk.MsgBoxOK|walk.MsgBoxIconError)
	}

	if err := ni.SetToolTip("메뉴를 선택하세요."); err != nil {

		walk.MsgBox(
			nil,
			"Error",
			err.Error(),
			walk.MsgBoxOK|walk.MsgBoxIconError)
	}

	ni.MouseDown().Attach(func(x, y int, button walk.MouseButton) {
		if button != walk.LeftButton {
			return
		}

		if err := ni.ShowCustom(
			GetFilename(name),
			day+"일 까지 완성할 글이 "+count+"자 기록되었습니다."); err != nil {

			log.Fatal(err)
		}
	})

	exitAction := walk.NewAction()
	if err := exitAction.SetText("E&xit"); err != nil {
		log.Fatal(err)
	}
	exitAction.Triggered().Attach(func() { walk.App().Exit(0) })
	if err := ni.ContextMenu().Actions().Add(exitAction); err != nil {
		log.Fatal(err)
	}

	if err := ni.SetVisible(true); err != nil {
		log.Fatal(err)
	}

	if err := ni.ShowInfo("마감 알리미", "다시 보려면 클릭하세요"); err != nil {
		log.Fatal(err)
	}

	mw.Run()

	// var mainWindow *walk.MainWindow

	// MainWindow{
	// 	AssignTo: &mainWindow,
	// 	Title:    "마감 안내기",
	// 	MinSize:  Size{200, 50},
	// 	Layout:   VBox{},
	// 	Children: []Widget{

	// 		Label{
	// 			Text: "글이름",
	// 		},
	// 		Label{
	// 			Text: GetFilename(name),
	// 		},
	// 		Label{
	// 			Text: "마감일",
	// 		},
	// 		Label{
	// 			Text: day,
	// 		},
	// 		Label{
	// 			Text: "글자수",
	// 		},
	// 		Label{
	// 			Text: count,
	// 		},
	// 		Label{
	// 			Text: "공백 제거 글자수",
	// 		},
	// 		Label{
	// 			Text: countWithoutBlank,
	// 		},
	// 	},
	// }.Run()

}
