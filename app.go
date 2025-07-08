package main

import (
	"github.com/gdamore/tcell"
	"log/slog"
)

type App struct {
	log         *slog.Logger
	screen      tcell.Screen
	quit        chan struct{}
	endSwimFish chan [2]int
}

func NewApp(log *slog.Logger) *App {
	backgroundStyle := tcell.StyleDefault.Foreground(tcell.ColorRed).Background(tcell.ColorWhite)

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

	return &App{
		log:         log,
		screen:      s,
		quit:        make(chan struct{}),
		endSwimFish: make(chan [2]int),
	}
}

func (a *App) InitStartDecorationAndFishes() {
	fishStyle := tcell.StyleDefault.Foreground(tcell.ColorBlue).Background(tcell.ColorWhite)

	decorations := make([]*Decoration, 0)

	decorations = append(decorations, NewDecoration(sea, 0, 5, a.screen, fishStyle))

	initialFishes := a.generateFishes(fishStyle)

	for _, fish := range initialFishes {
		go func() {
			go fish.Swim()
			//<-fish.endSwim
		}()
	}

	for _, decoration := range decorations {
		decoration.Draw()

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
	x, y   int
	screen tcell.Screen
	style  tcell.Style
}

func NewDecoration(model []string, x, y int, screen tcell.Screen, style tcell.Style) *Decoration {
	return &Decoration{model: model, x: x, y: y, screen: screen, style: style}
}

func (d *Decoration) Draw() {
	for col := 0; col < len(d.model); col++ {
		for row := 0; row < len(d.model[col]); row++ {
			d.screen.SetContent(d.x+row, d.y+col, rune(d.model[col][row]), nil, d.style)
		}
	}
}
