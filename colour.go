package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func ColourSetting() {
	MainWindow{
		Title:   "디자인 설정",
		MinSize: Size{400, 0},
		Background: GradientBrush{
			Vertexes: []walk.GradientVertex{
				{X: 0, Y: 0, Color: walk.RGB(255, 255, 127)},
				{X: 1, Y: 0, Color: walk.RGB(127, 191, 255)},
				{X: 0.5, Y: 0.5, Color: walk.RGB(255, 255, 255)},
				{X: 1, Y: 1, Color: walk.RGB(127, 255, 127)},
				{X: 0, Y: 1, Color: walk.RGB(255, 127, 127)},
			},
			Triangles: []walk.GradientTriangle{
				{0, 1, 2},
				{1, 3, 2},
				{3, 4, 2},
				{4, 0, 2},
			},
		},
		Layout: HBox{Margins: Margins{100, 100, 100, 100}},
		Children: []Widget{
			GradientComposite{
				Border:   true,
				Vertical: Bind("verticalCB.Checked"),
				Color1:   Bind("rgb(c1RedSld.Value, c1GreenSld.Value, c1BlueSld.Value)"),
				Color2:   Bind("rgb(c2RedSld.Value, c2GreenSld.Value, c2BlueSld.Value)"),
				Layout:   HBox{},
				Children: []Widget{
					GroupBox{
						Title:  "색 파라미터",
						Layout: VBox{},
						Children: []Widget{
							CheckBox{Name: "verticalCB", Text: "세로", Checked: true},
							GroupBox{
								Title:  "색 파레트 1",
								Layout: Grid{Columns: 2},
								Children: []Widget{
									Label{Text: "Red:"},
									Slider{Name: "c1RedSld", Tracking: true, MaxValue: 255, Value: 95},
									Label{Text: "Green:"},
									Slider{Name: "c1GreenSld", Tracking: true, MaxValue: 255, Value: 191},
									Label{Text: "Blue:"},
									Slider{Name: "c1BlueSld", Tracking: true, MaxValue: 255, Value: 255},
								},
							},
							GroupBox{
								Title:  "색 파레트 2",
								Layout: Grid{Columns: 2},
								Children: []Widget{
									Label{Text: "Red:"},
									Slider{Name: "c2RedSld", Tracking: true, MaxValue: 255, Value: 239},
									Label{Text: "Green:"},
									Slider{Name: "c2GreenSld", Tracking: true, MaxValue: 255, Value: 63},
									Label{Text: "Blue:"},
									Slider{Name: "c2BlueSld", Tracking: true, MaxValue: 255, Value: 0},
								},
							},

							PushButton{
								Text: "색 정하기",
								OnClicked: func() {
									walk.MsgBox(
										nil,
										"색",
										GradientComposite,
										walk.MsgBoxIconInformation)

								},
							},
						},
					},
				},
			},
		},
		Functions: map[string]func(args ...interface{}) (interface{}, error){
			"rgb": func(args ...interface{}) (interface{}, error) {
				return walk.RGB(byte(args[0].(float64)), byte(args[1].(float64)), byte(args[2].(float64))), nil
			},
		},
	}.Run()

}
