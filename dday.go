package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func DdayCounter() {

	var TextCounter, Day *walk.TextEdit

	MainWindow{
		Title:   "다가오는 마감",
		MinSize: Size{400, 400},
		Layout:  VBox{},
		Children: []Widget{
			HSplitter{
				Children: []Widget{
					TextEdit{AssignTo: &TextCounter},
					TextEdit{AssignTo: &Day, ReadOnly: true},
				},
			},
		},
	}.Run()
}
