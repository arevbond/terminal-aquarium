package main

import (
	"github.com/gdamore/tcell"
	"log/slog"
	"strings"
	"time"
)

type App struct {
	log                *slog.Logger
	screen             tcell.Screen
	quit               chan struct{}
	endSwimFish        chan [2]int
	seaBackgroundColor tcell.Color
	width, height      int
}

func NewApp(log *slog.Logger) *App {
	backgroundColor := tcell.ColorAqua

	backgroundStyle := tcell.StyleDefault.Foreground(tcell.ColorRed).Background(backgroundColor)

	s, err := tcell.NewScreen()
	if err != nil {
		log.Error("%+v", err)
		return nil
	}

	if err = s.Init(); err != nil {
		log.Error("%+v", err)
		return nil
	}

	s.SetStyle(backgroundStyle)

	w, h := s.Size()

	return &App{
		log:                log,
		screen:             s,
		quit:               make(chan struct{}),
		endSwimFish:        make(chan [2]int),
		seaBackgroundColor: backgroundColor,
		width:              w,
		height:             h,
	}
}

func (a *App) InitStartDecorationAndFishes() {
	a.SetSkyColor(tcell.ColorWhite)

	a.InitSeaWithResizeHandling()

	fishStyle := tcell.StyleDefault.Foreground(tcell.ColorBlue).Background(tcell.ColorWhite)

	initialFishes := a.generateFishes(fishStyle)

	for _, fish := range initialFishes {
		go func() {
			go fish.Swim()
			//<-fish.endSwim
		}()
	}

}

func (a *App) Run() error {
	go a.HandleShutdown()

	for {
		a.screen.Show()

		select {
		case <-a.quit:
			return nil
		default:
		}
	}
}

func (a *App) HandleShutdown() {
	for {
		ev := a.screen.PollEvent()

		switch ev := ev.(type) {
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyCtrlC {
				a.quit <- struct{}{}
			}
		default:
		}
	}
}

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
			}

			time.Sleep(1 * time.Second)
		}
	}()
}

func (a *App) ScreenResized() bool {
	w, h := a.screen.Size()

	if a.width != w || a.height != h {
		a.width = w
		a.height = h

		return true
	}

	return false
}

func (a *App) SetSkyColor(color tcell.Color) {
	for col := 0; col <= a.width; col++ {
		for row := 0; row < 5; row++ {
			a.screen.SetContent(col, row, ' ', nil, tcell.StyleDefault.Foreground(color).Background(color))
		}
	}
}
