package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/lxn/walk"

	. "github.com/lxn/walk/declarative"
)

type Directory struct {
	name     string
	parent   *Directory
	children []*Directory
}

type TxtFile struct {
	name  string
	path  string
	count int
}

func NewDirectory(name string, parent *Directory) *Directory {
	return &Directory{name: name, parent: parent}
}

var _ walk.TreeItem = new(Directory)

func (d *Directory) Text() string {
	return d.name
}

func (d *Directory) Parent() walk.TreeItem {
	if d.parent == nil {
		// We can't simply return d.parent in this case, because the interface
		// value then would not be nil.
		return nil
	}

	return d.parent
}

func (d *Directory) ChildCount() int {
	if d.children == nil {
		// It seems this is the first time our child count is checked, so we
		// use the opportunity to populate our direct children.
		if err := d.ResetChildren(); err != nil {
			log.Print(err)
		}
	}

	return len(d.children)
}

func (d *Directory) ChildAt(index int) walk.TreeItem {
	return d.children[index]
}

func (d *Directory) Image() interface{} {
	return d.Path()
}

func (d *Directory) ResetChildren() error {
	d.children = nil

	dirPath := d.Path()

	if err := filepath.Walk(d.Path(), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			if info == nil {
				return filepath.SkipDir
			}
		}

		name := info.Name()

		if !info.IsDir() || path == dirPath || shouldExclude(name) {
			return nil
		}

		d.children = append(d.children, NewDirectory(name, d))

		return filepath.SkipDir
	}); err != nil {
		return err
	}

	return nil
}

func (d *Directory) Path() string {
	elems := []string{d.name}

	dir, _ := d.Parent().(*Directory)

	for dir != nil {
		elems = append([]string{dir.name}, elems...)
		dir, _ = dir.Parent().(*Directory)
	}

	return filepath.Join(elems...)
}

type DirectoryTreeModel struct {
	walk.TreeModelBase
	roots []*Directory
}

var _ walk.TreeModel = new(DirectoryTreeModel)

func NewDirectoryTreeModel() (*DirectoryTreeModel, error) {
	model := new(DirectoryTreeModel)

	drives, err := walk.DriveNames()
	if err != nil {
		return nil, err
	}

	for _, drive := range drives {
		switch drive {
		case "A:\\", "B:\\":
			continue
		}

		model.roots = append(model.roots, NewDirectory(drive, nil))
	}

	return model, nil
}

func (*DirectoryTreeModel) LazyPopulation() bool {
	// We don't want to eagerly populate our tree view with the whole file system.
	return true
}

func (m *DirectoryTreeModel) RootCount() int {
	return len(m.roots)
}

func (m *DirectoryTreeModel) RootAt(index int) walk.TreeItem {
	return m.roots[index]
}

type FileInfo struct {
	Name     string
	Size     int64
	Modified time.Time
}

type FileInfoModel struct {
	walk.SortedReflectTableModelBase
	dirPath string
	items   []*FileInfo
}

var _ walk.ReflectTableModel = new(FileInfoModel)

func NewFileInfoModel() *FileInfoModel {
	return new(FileInfoModel)
}

func (m *FileInfoModel) Items() interface{} {
	return m.items
}

func (m *FileInfoModel) SetDirPath(dirPath string) error {
	m.dirPath = dirPath
	m.items = nil

	if err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			if info == nil {
				return filepath.SkipDir
			}
		}

		name := info.Name()

		if path == dirPath || shouldExclude(name) {
			return nil
		}

		item := &FileInfo{
			Name:     name,
			Size:     info.Size(),
			Modified: info.ModTime(),
		}

		m.items = append(m.items, item)

		if info.IsDir() {
			return filepath.SkipDir
		}

		return nil
	}); err != nil {
		return err
	}

	m.PublishRowsReset()

	return nil
}

func (m *FileInfoModel) Image(row int) interface{} {
	return filepath.Join(m.dirPath, m.items[row].Name)
}

func shouldExclude(name string) bool {
	switch name {
	case "System Volume Information", "pagefile.sys", "swapfile.sys":
		return true
	}

	return false
}

func Fileupload() {
	var mainWindow *walk.MainWindow
	var splitter *walk.Splitter
	var treeView *walk.TreeView
	var tableView *walk.TableView
	var webView *walk.WebView

	treeModel, err := NewDirectoryTreeModel()
	if err != nil {
		log.Fatal(err)
	}
	tableModel := NewFileInfoModel()

	if err := (MainWindow{
		AssignTo: &mainWindow,
		Title:    "작업할 파일을 고르세요(.txt만 지원)",
		MinSize:  Size{600, 400},
		Size:     Size{1024, 640},
		Layout:   HBox{MarginsZero: true},
		Children: []Widget{
			HSplitter{
				AssignTo: &splitter,
				Children: []Widget{
					TreeView{
						AssignTo: &treeView,
						Model:    treeModel,
						OnCurrentItemChanged: func() {
							dir := treeView.CurrentItem().(*Directory)

							if err := tableModel.SetDirPath(dir.Path()); err != nil {
								walk.MsgBox(
									mainWindow,
									"Error",
									err.Error(),
									walk.MsgBoxOK|walk.MsgBoxIconError)
							}
						},
					},
					TableView{
						AssignTo:      &tableView,
						StretchFactor: 2,
						Columns: []TableViewColumn{
							TableViewColumn{
								DataMember: "Name",
								Width:      192,
							},
							TableViewColumn{
								DataMember: "Size",
								Format:     "%d",
								Alignment:  AlignFar,
								Width:      64,
							},
							TableViewColumn{
								DataMember: "Modified",
								Format:     "2006-01-02 15:04:05",
								Width:      120,
							},
						},
						Model: tableModel,
						OnCurrentIndexChanged: func() {

							var url string
							if index := tableView.CurrentIndex(); index > -1 {
								name := tableModel.items[index].Name
								dir := treeView.CurrentItem().(*Directory)
								url = filepath.Join(dir.Path(), name)
							}

							webView.SetURL(url)
						},
					},

					WebView{
						AssignTo:      &webView,
						StretchFactor: 2,
					},
				},
			},

			PushButton{
				Text: "등록하기",
				OnClicked: func() {
					// 1. txt일때만 등록(나중에 doc도 서포트하자)
					if index := tableView.CurrentIndex(); index > -1 {

						if !strings.Contains(tableModel.items[index].Name, "txt") {

							walk.MsgBox(
								mainWindow,
								"파일 형식 오류",
								".txt파일만 지원합니다",
								walk.MsgBoxOK|walk.MsgBoxIconError)

						} else {
							// 2. 텍스트 파일 불러오기
							dir := treeView.CurrentItem().(*Directory)
							url := filepath.Join(dir.Path(), tableModel.items[index].Name)
							txt := txtFileOpen(url)
							/*
								TxtFile{
									name:  tableModel.items[index].Name,
									path:  url,
									count: CountChar(txt),
								}*/

							// 3. 마감일 정하기
							DdaySet()

							// 4. 알리미로 넘어가기
							walk.MsgBox(
								mainWindow,
								"테스트",
								txt,
								walk.MsgBoxOK|walk.MsgBoxIconError)
						}

					}
					//dir := treeView.CurrentItem().(*Directory)

					// 2. txt일때 >
					//    가. 글자수 세기
					//    나. 마감일 정하기
					// 3.  새 창에 알림 띄우기 : 글자수, 마감일
				},
			},
		},
	}.Create()); err != nil {
		log.Fatal(err)
	}

	splitter.SetFixed(treeView, true)
	splitter.SetFixed(tableView, true)

	mainWindow.Run()
}

func txtFileOpen(filepath string) string {
	file, err := os.Open(filepath)
	if err != nil {
		return err.Error()
	}
	defer file.Close()

	fi, err := file.Stat()
	if err != nil {
		return err.Error()
	}

	var data = make([]byte, fi.Size())

	_, err = file.Read(data)
	if err != nil {
		return err.Error()
	}

	return string(data)
}
