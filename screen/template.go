package screen

import (
	"log"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type Template struct {
	source      *widgets.List
	transfer    *widgets.List
	destination *widgets.List
}

func NewTemplate() *Template {
	c := Template{}
	c.source = widgets.NewList()
	c.transfer = widgets.NewList()
	c.destination = widgets.NewList()
	return &c
}

func (c *Template) Run() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	setList("Left", c.source)
	setList("", c.transfer)
	c.transfer.Border = false
	setList("Right", c.destination)

	add(c.source, "there")
	prepend(c.source, "hi")

	grid := ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, termHeight)

	grid.Set(
		ui.NewRow(1.0,
			ui.NewCol(0.33, c.source),
			ui.NewCol(0.33, c.transfer),
			ui.NewCol(0.33, c.destination),
		),
	)

	ui.Render(grid)
	uiEvents := ui.PollEvents()
	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				return
			case "<Resize>":
				payload := e.Payload.(ui.Resize)
				grid.SetRect(0, 0, payload.Width, payload.Height)
				ui.Clear()
			case "<Enter>":
				c.handleEnter()
			}
		}
		ui.Render(grid)
	}
}

func (c *Template) handleEnter() {
}
