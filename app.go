package main

import (
	"github.com/gdamore/tcell"
	"log/slog"
	"math/rand"
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

	//initialFishes := a.generateFishes()

	initialFishes := a.generateRandomFish(15)

	for _, fish := range initialFishes {
		go fish.Swim()
	}

}

func (a *App) InsertRandomFish() {
	swimForward := rand.Intn(2) == 0

	var model []string

	if swimForward {
		model = fishModelsForward[rand.Intn(len(fishModelsForward))]
	} else {
		model = fishModelsBackward[rand.Intn(len(fishModelsBackward))]
	}

	var x int
	if swimForward {
		x = 0 - len(model[0])
	} else {
		x = a.width + len(model[0])
	}
	y := 6 + rand.Intn(a.height-6)

	fish := NewFishWithRandomColor(model, swimForward, Speed(rand.Intn(4)), x, y, a.screen, a.seaBackgroundColor, a.log)

	go fish.Swim()
}

func (a *App) Run() error {
	go a.HandleShutdown()

	go func() {
		for {
			a.InsertRandomFish()

			time.Sleep(3 * time.Second)
		}
	}()

	for {
		a.screen.Show()

		select {
		case <-a.quit:
			return nil
		default:
			time.Sleep(16 * time.Millisecond)
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
