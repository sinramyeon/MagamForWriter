package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func Alarm(day, name, count string) {
	// ticker := time.NewTicker(10 * time.Minute)
	// quit := make(chan struct{})
	// go func() {
	// 	for {
	// 		select {
	// 		case <-ticker.C:
	// 			// do stuff
	// 			txtfile, _ := ioutil.ReadFile(name)
	// 			count = CountChar(string(txtfile))
	// 		case <-quit:
	// 			ticker.Stop()
	// 			return
	// 		}
	// 	}
	// }()

	walk.MsgBox(
		nil,
		"Test",
		"Alarm",
		walk.MsgBoxOK|walk.MsgBoxIconError)

	var mainWindow *walk.MainWindow

	MainWindow{
		AssignTo: &mainWindow,
		Title:    "마감 안내기",
		MinSize:  Size{100, 400},
		Layout:   VBox{},
		Children: []Widget{

			Label{
				Text: "글이름",
			},
			Label{
				Text: name,
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
		},
	}.Run()

}
