package main

import (
	"fmt"
	"github.com/gdamore/tcell"
	"log"
	"log/slog"
	"os"
)

func main() {
	file, err := os.OpenFile("logs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	logHandler := slog.NewTextHandler(file, &slog.HandlerOptions{Level: slog.LevelDebug})
	logger := slog.New(logHandler)
	backgroundStyle := tcell.StyleDefault.Foreground(tcell.ColorRed).Background(tcell.ColorWhite)

	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}

	if err = s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}

	s.SetStyle(backgroundStyle)
	if err = s.Beep(); err != nil {
		fmt.Println("err", err)
	}

	finish := func() {
		maybePanic := recover()
		s.Fini()
		if maybePanic != nil {
			panic(maybePanic)
		}
	}
	defer finish()

	fishStyle := tcell.StyleDefault.Foreground(tcell.ColorBlue).Background(tcell.ColorWhite)

	decorations := make([]*Decoration, 0)

	decorations = append(decorations, NewDecoration(sea, 0, 5, s, fishStyle))

	fishes := generateFishes(s, fishStyle, logger)

	for _, fish := range fishes {
		go func() {
			fish.Swim()
		}()
	}

	for _, decoration := range decorations {
		decoration.Draw()

	}

	quit := make(chan struct{})

	go func() {
		for {
			ev := s.PollEvent()

			switch ev := ev.(type) {
			case *tcell.EventKey:
				if ev.Key() == tcell.KeyCtrlC {
					quit <- struct{}{}
				}
			default:
			}
		}
	}()

	for {
		s.Show()

		select {
		case <-quit:
			return
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

func updateFishes(fishes []*Fish) []*Fish {
	return nil
}
