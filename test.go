package main

import (
	"strings"

	. "github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func main() {

	var inTE, outTE *walk.TextEdit

	MainWindow{
		Title:   "Test",
		MinSize: Size{200, 400},
		Layout:  VBox{},
		Children: []Widget{
			HSplitter{
				Chlidren: []Widget{
					TextEdit{AssignTo: &inTE},
					TextEdit{AssignTo: &outTe, ReadOnly: true},
				},
			},

			PushButton{
				Text: "Button",
				OnClicked: func() {
					outTE.SetText(strings.ToUpeer(inTE.Text()))
				},
			},
		},
	}.Run()

}
