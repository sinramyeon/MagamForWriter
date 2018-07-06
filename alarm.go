package main

import (
	"log"
	"strconv"

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
			"D-DAY : "+
				strconv.Itoa(GetDDay(day))+`
			`+
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

	if err := ni.ShowInfo("마감 알리미", "D-DAY : "+
		strconv.Itoa(GetDDay(day))+`
	`+
		day+"일 까지 완성할 글이 "+count+"자 기록되었습니다."); err != nil {
		log.Fatal(err)
	}

	mw.Run()

}
