package main

import (
	"log"

	"github.com/lxn/walk"
)

func Alarm(day, name, count, countWithoutBlank string) {

	mw, err := walk.NewMainWindow()
	if err != nil {
		WalkError(err)
	}

	icon, err := walk.Resources.Icon("./file.ico")
	if err != nil {
		WalkError(err)
	}

	ni, err := walk.NewNotifyIcon()
	if err != nil {
		WalkError(err)
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
	if err := exitAction.SetText("종료"); err != nil {
		log.Fatal(err)
	}
	exitAction.Triggered().Attach(func() { walk.App().Exit(0) })
	if err := ni.ContextMenu().Actions().Add(exitAction); err != nil {
		log.Fatal(err)
	}

	if err := ni.SetVisible(true); err != nil {
		log.Fatal(err)
	}

	if err := ni.ShowInfo("마감 알리미", day+"일 까지 완성할 글이 "+count+"자 기록되었습니다."); err != nil {
		log.Fatal(err)
	}

	mw.Run()

	/*

		// We need either a walk.MainWindow or a walk.Dialog for their message loop.
		// We will not make it visible in this example, though.
		mw, err := walk.NewMainWindow()
		if err != nil {
			log.Fatal(err)
		}

		// We load our icon from a file.
		icon, err := walk.Resources.Icon("../img/stop.ico")
		if err != nil {
			log.Fatal(err)
		}

		// Create the notify icon and make sure we clean it up on exit.
		ni, err := walk.NewNotifyIcon()
		if err != nil {
			log.Fatal(err)
		}
		defer ni.Dispose()

		// Set the icon and a tool tip text.
		if err := ni.SetIcon(icon); err != nil {
			log.Fatal(err)
		}
		if err := ni.SetToolTip("Click for info or use the context menu to exit."); err != nil {
			log.Fatal(err)
		}

		// When the left mouse button is pressed, bring up our balloon.
		ni.MouseDown().Attach(func(x, y int, button walk.MouseButton) {
			if button != walk.LeftButton {
				return
			}

			if err := ni.ShowCustom(
				"Walk NotifyIcon Example",
				"There are multiple ShowX methods sporting different icons."); err != nil {

				log.Fatal(err)
			}
		})

		// We put an exit action into the context menu.
		exitAction := walk.NewAction()
		if err := exitAction.SetText("E&xit"); err != nil {
			log.Fatal(err)
		}
		exitAction.Triggered().Attach(func() { walk.App().Exit(0) })
		if err := ni.ContextMenu().Actions().Add(exitAction); err != nil {
			log.Fatal(err)
		}

		// The notify icon is hidden initially, so we have to make it visible.
		if err := ni.SetVisible(true); err != nil {
			log.Fatal(err)
		}

		// Now that the icon is visible, we can bring up an info balloon.
		if err := ni.ShowInfo("Walk NotifyIcon Example", "Click the icon to show again."); err != nil {
			log.Fatal(err)
		}

		// Run the message loop.
		mw.Run()

		이예제를 복붙을해도 안뜸
	*/

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
