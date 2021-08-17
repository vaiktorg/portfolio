package components

import (
	. "github.com/maxence-charriere/go-app/pkg/app"
	"github.com/vaiktorg/grimoire/documentstore"
)

// Directory
type Tree struct {
	Compo
	//tree *documentstore.Dir
	dir *Dir
}

//
//func (t *Tree) OnMount(Context) {
//	t.tree = documentstore.NewDir("static")
//}
//
//func (t *Tree) Render() UI {
//	return Ul().Class("tree").Body(
//		&Dir{Name: t.tree.Meta.Name},
//	)
//}
//
//func (t *Tree) GetStruct() *documentstore.Dir {
//	return t.tree
//}

type File struct {
	Compo
	Name string
	Path string
}

func (f *File) AddFile() {}

type Dir struct {
	Compo
	Name  string
	dirs  []*Dir
	files []*File
}

func (d *Dir) Render() UI {
	return Li().Body(
		Span().Class("caret").Text(d.Name),
		Ul().Class("nested").Body(
			Range(d.files).Slice(func(i int) UI {
				return d.files[i]
			}),
			Range(d.dirs).Slice(func(i int) UI {
				return d.dirs[i]
			}),
		),
	)
}

func (d *Dir) AddFile(file ...documentstore.File) {
	for _, f := range file {
		d.files = append(d.files, &File{Name: f.Meta.Name})
	}
}

func (d *Dir) AddDir(dirs ...documentstore.Dir) {
	for _, dir := range dirs {
		d.dirs = append(d.dirs, &Dir{Name: dir.Meta.Name})
	}
}
