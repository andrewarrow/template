package screen

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func setList(title string, l *widgets.List) {
	l.SelectedRowStyle.Fg = ui.ColorWhite
	l.SelectedRowStyle.Bg = ui.ColorMagenta
	l.TextStyle.Fg = ui.ColorWhite
	l.TextStyle.Bg = ui.ColorBlack
	l.Title = title
}

func add(l *widgets.List, s string) {
	l.Rows = append(l.Rows, s)
}

func prepend(l *widgets.List, s string) {
	l.Rows = append([]string{s}, l.Rows...)
}
