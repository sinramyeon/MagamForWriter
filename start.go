// package main

// import (
// 	"github.com/lxn/walk"
// 	. "github.com/lxn/walk/declarative"
// )

// type MyMainWindow struct {
// 	*walk.MainWindow
// }

// func main() {
// 	mw := new(MyMainWindow)

// 	var openAction, showAboutBoxAction *walk.Action
// 	var recentMenu *walk.Menu

// 	if err := (MainWindow{
// 		AssignTo: &mw.MainWindow,
// 		Title:    "마감 안내기",
// 		MenuItems: []MenuItem{
// 			Menu{
// 				Text: "&File",
// 				Items: []MenuItem{
// 					Action{
// 						AssignTo:    &openAction,
// 						Text:        "&Open",
// 						Enabled:     Bind("enabledCB.Checked"),
// 						Visible:     Bind("!openHiddenCB.Checked"),
// 						Shortcut:    Shortcut{walk.ModControl, walk.KeyO},
// 						OnTriggered: mw.openAction_Triggered,
// 					},
// 					Menu{
// 						AssignTo: &recentMenu,
// 						Text:     "Recent",
// 					},
// 					Separator{},
// 					Action{
// 						Text: "E&xit",
// 						OnTriggered: func() {
// 							//mw.Close()
// 						},
// 					},
// 				},
// 			},
// 			Menu{
// 				Text: "&Help",
// 				Items: []MenuItem{
// 					Action{
// 						AssignTo:    &showAboutBoxAction,
// 						Text:        "About",
// 						OnTriggered: mw.showAboutBoxAction_Triggered,
// 					},
// 				},
// 			},
// 		},
// 		ToolBar: ToolBar{
// 			ButtonStyle: ToolBarButtonImageBeforeText,
// 			Items: []MenuItem{
// 				ActionRef{&openAction},
// 				Menu{
// 					Text: "New A",
// 					Items: []MenuItem{
// 						Action{
// 							Text:        "A",
// 							OnTriggered: mw.newAction_Triggered,
// 						},
// 						Action{
// 							Text:        "B",
// 							OnTriggered: mw.newAction_Triggered,
// 						},
// 						Action{
// 							Text:        "C",
// 							OnTriggered: mw.newAction_Triggered,
// 						},
// 					},
// 					OnTriggered: mw.newAction_Triggered,
// 				},
// 				Separator{},
// 				Menu{
// 					Text: "View",
// 					Items: []MenuItem{
// 						Action{
// 							Text:        "X",
// 							OnTriggered: mw.changeViewAction_Triggered,
// 						},
// 						Action{
// 							Text:        "Y",
// 							OnTriggered: mw.changeViewAction_Triggered,
// 						},
// 						Action{
// 							Text:        "Z",
// 							OnTriggered: mw.changeViewAction_Triggered,
// 						},
// 					},
// 				},
// 			},
// 		},
// 		ContextMenuItems: []MenuItem{
// 			ActionRef{&showAboutBoxAction},
// 		},
// 		MinSize: Size{300, 200},
// 		Layout:  VBox{},
// 		Children: []Widget{
// 			CheckBox{
// 				Name:    "enabledCB",
// 				Text:    "Open / Special Enabled",
// 				Checked: true,
// 			},
// 			CheckBox{
// 				Name:    "openHiddenCB",
// 				Text:    "Open Hidden",
// 				Checked: true,
// 			},
// 		},
// 	}.Create()); err != nil {

// 		walk.MsgBox(
// 			nil,
// 			"Error",
// 			err.Error(),
// 			walk.MsgBoxOK|walk.MsgBoxIconError)
// 	}

// 	// var mainWindow *walk.MainWindow

// 	// MainWindow{
// 	// 	AssignTo: &mainWindow,
// 	// 	Title:    "마감 안내기",
// 	// 	MinSize:  Size{120, 50},
// 	// 	Layout:   VBox{},
// 	// 	Children: []Widget{

// 	// 		PushButton{
// 	// 			Text: "작업 등록하기",
// 	// 			OnClicked: func() {
// 	// 				mainWindow.Close()
// 	// 				Fileupload()
// 	// 			},
// 	// 		},
// 	// 		PushButton{
// 	// 			Text: "마감일 안내받기",
// 	// 			OnClicked: func() {
// 	// 				day, name, count, countWithoutBlank := GetAlarmText()
// 	// 				Alarm(day, name, count, countWithoutBlank)
// 	// 				mainWindow.Close()
// 	// 			},
// 	// 		},
// 	// 	},
// 	// }.Run()
// }

// func (mw *MyMainWindow) openAction_Triggered() {
// 	walk.MsgBox(mw, "Open", "Pretend to open a file...", walk.MsgBoxIconInformation)
// }

// func (mw *MyMainWindow) newAction_Triggered() {
// 	walk.MsgBox(mw, "New", "Newing something up... or not.", walk.MsgBoxIconInformation)
// }

// func (mw *MyMainWindow) changeViewAction_Triggered() {
// 	walk.MsgBox(mw, "Change View", "By now you may have guessed it. Nothing changed.", walk.MsgBoxIconInformation)
// }

// func (mw *MyMainWindow) showAboutBoxAction_Triggered() {
// 	walk.MsgBox(mw, "About", "Walk Actions Example", walk.MsgBoxIconInformation)
// }

// func (mw *MyMainWindow) specialAction_Triggered() {
// 	walk.MsgBox(mw, "Special", "Nothing to see here.", walk.MsgBoxIconInformation)
// }

// Copyright 2013 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"log"
)

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

var isSpecialMode = walk.NewMutableCondition()

type MyMainWindow struct {
	*walk.MainWindow
}

func main() {
	MustRegisterCondition("isSpecialMode", isSpecialMode)

	mw := new(MyMainWindow)

	var openAction, showAboutBoxAction, fileUploadAction *walk.Action
	var recentMenu *walk.Menu
	var toggleSpecialModePB *walk.PushButton

	if err := (MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    "마감 안내기",
		MenuItems: []MenuItem{
			Menu{
				Text: "&File",
				Items: []MenuItem{
					Action{
						AssignTo:    &openAction,
						Text:        "&Open",
						Enabled:     Bind("enabledCB.Checked"),
						Visible:     Bind("!openHiddenCB.Checked"),
						Shortcut:    Shortcut{walk.ModControl, walk.KeyO},
						OnTriggered: mw.openAction_Triggered,
					},
					Menu{
						AssignTo: &recentMenu,
						Text:     "Recent",
					},
					Separator{},
					Action{
						Text:        "E&xit",
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
				ActionRef{&openAction},
				Action{
					AssignTo:    &fileUploadAction,
					Text:        "파일 추가",
					OnTriggered: mw.fileUploadAction_Triggered,
				},
				Separator{},
				Menu{
					Text: "View",
					Items: []MenuItem{
						Action{
							Text:        "X",
							OnTriggered: mw.changeViewAction_Triggered,
						},
						Action{
							Text:        "Y",
							OnTriggered: mw.changeViewAction_Triggered,
						},
						Action{
							Text:        "Z",
							OnTriggered: mw.changeViewAction_Triggered,
						},
					},
				},
				Separator{},
				Action{
					Text:        "Special",
					Enabled:     Bind("isSpecialMode && enabledCB.Checked"),
					OnTriggered: mw.specialAction_Triggered,
				},
			},
		},
		ContextMenuItems: []MenuItem{
			ActionRef{&showAboutBoxAction},
		},
		MinSize: Size{300, 200},
		Layout:  VBox{},
		Children: []Widget{
			CheckBox{
				Name:    "enabledCB",
				Text:    "Open / Special Enabled",
				Checked: true,
			},
			CheckBox{
				Name:    "openHiddenCB",
				Text:    "Open Hidden",
				Checked: true,
			},
			PushButton{
				AssignTo: &toggleSpecialModePB,
				Text:     "Enable Special Mode",
				OnClicked: func() {
					isSpecialMode.SetSatisfied(!isSpecialMode.Satisfied())

					if isSpecialMode.Satisfied() {
						toggleSpecialModePB.SetText("Disable Special Mode")
					} else {
						toggleSpecialModePB.SetText("Enable Special Mode")
					}
				},
			},
		},
	}.Create()); err != nil {
		log.Fatal(err)
	}

	addRecentFileActions := func(texts ...string) {
		for _, text := range texts {
			a := walk.NewAction()
			a.SetText(text)
			a.Triggered().Attach(mw.openAction_Triggered)
			recentMenu.Actions().Add(a)
		}
	}

	addRecentFileActions("Foo", "Bar", "Baz")

	mw.Run()
}

func (mw *MyMainWindow) openAction_Triggered() {
	walk.MsgBox(mw, "Open", "Pretend to open a file...", walk.MsgBoxIconInformation)
}

func (mw *MyMainWindow) newAction_Triggered() {
	walk.MsgBox(mw, "New", "Newing something up... or not.", walk.MsgBoxIconInformation)
}

func (mw *MyMainWindow) changeViewAction_Triggered() {
	walk.MsgBox(mw, "Change View", "By now you may have guessed it. Nothing changed.", walk.MsgBoxIconInformation)
}

func (mw *MyMainWindow) showAboutBoxAction_Triggered() {
	walk.MsgBox(mw, "About", "글 쓰시는 분들의 마감을 도와드립니다. 20180702 히어로 제작", walk.MsgBoxIconInformation)
}

func (mw *MyMainWindow) fileUploadAction_Triggered() {
	Fileupload()
}

func (mw *MyMainWindow) specialAction_Triggered() {
	walk.MsgBox(mw, "Special", "Nothing to see here.", walk.MsgBoxIconInformation)
}
