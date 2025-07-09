package main

import (
	"github.com/gdamore/tcell"
	"strings"
	"time"
)

type Decoration struct {
	model  []string
	screen tcell.Screen
	style  tcell.Style
}

func NewDecoration(model []string, screen tcell.Screen, style tcell.Style) *Decoration {
	return &Decoration{model: model, screen: screen, style: style}
}

func (d *Decoration) Draw(x, y int) {
	for col := 0; col < len(d.model); col++ {
		for row := 0; row < len(d.model[col]); row++ {
			d.screen.SetContent(x+row, y+col, rune(d.model[col][row]), nil, d.style)
		}
	}
}

func (a *App) DrawSea(x, y int) {
	style := tcell.StyleDefault.Foreground(tcell.ColorBlue).Background(a.seaBackgroundColor)

	sea := strings.Repeat(sea, x/2)
	for row := 0; row < len(sea); row++ {
		a.screen.SetContent(row, y, rune(sea[row]), nil, style)
	}
}

func (a *App) InitSeaWithResizeHandling() {
	a.DrawSea(a.width, 5)

	go func() {
		for {
			if a.ScreenResized() {
				a.DrawSea(a.width, 5)
				a.SetSkyColor(tcell.ColorWhite)
			}

			time.Sleep(1 * time.Second)
		}
	}()
}
