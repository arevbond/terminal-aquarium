package main

import (
	"github.com/gdamore/tcell"
	"log/slog"
	"math/rand"
	"time"
)

func generateFishes(s tcell.Screen, fishStyle tcell.Style, log *slog.Logger) []*Fish {
	fishes := make([]*Fish, 0)

	fishes = append(fishes, NewFish(whaleBackward, false, HighSpeed, 100, 0, s, fishStyle, log))

	fishes = append(fishes, NewFish(fishForward, true, MediumSpeed, rand.Intn(10), 7, s, fishStyle, log))
	fishes = append(fishes, NewFish(fishForward2, true, MediumSpeed, rand.Intn(10), 23, s, fishStyle, log))
	fishes = append(fishes, NewFish(fishForward, true, HighSpeed, rand.Intn(10), 37, s, fishStyle, log))

	fishes = append(fishes, NewFish(fishBackward2, false, LowSpeed, 100+rand.Intn(10), 15, s, fishStyle, log))
	fishes = append(fishes, NewFish(fishBackward, false, MediumSpeed, 100+rand.Intn(10), 30, s, fishStyle, log))

	return fishes
}

type Speed int

const (
	DisableSpeed Speed = iota
	LowSpeed
	MediumSpeed
	HighSpeed
)

type Fish struct {
	model       []string
	curX, curY  int
	style       tcell.Style
	screen      tcell.Screen
	swimForward bool
	speed       Speed
	logger      *slog.Logger
	// TODO: добавить сигнал о том. что рыба достигла границы
	//endSwim     chan struct{}
}

func NewFish(model []string, swimForward bool, speed Speed, x, y int, screen tcell.Screen, style tcell.Style, logger *slog.Logger) *Fish {
	return &Fish{model: model, curX: x, curY: y, speed: speed, style: style, screen: screen, swimForward: swimForward, logger: logger}
}

func (f *Fish) Draw() {
	for col := 0; col < len(f.model); col++ {
		for row := 0; row < len(f.model[col]); row++ {
			f.screen.SetContent(f.curX+row, f.curY+col, rune(f.model[col][row]), nil, f.style)
		}
	}
}

func (f *Fish) Move() {
	f.Clear()

	if f.swimForward {
		f.curX++
	} else {
		f.curX--
	}

	f.Draw()
}

func (f *Fish) Swim() {
	for {
		f.Move()
		x, _ := f.screen.Size()
		if f.curX == x || f.curX == 0 {
			f.logger.Info("fish has border", slog.Int("x", f.curX), slog.Int("y", f.curY))
		}
		f.SleepBySpeed()

	}
}

func (f *Fish) SleepBySpeed() {
	switch f.speed {
	case LowSpeed:
		time.Sleep(800 * time.Millisecond)
	case MediumSpeed:
		time.Sleep(400 * time.Millisecond)
	case HighSpeed:
		time.Sleep(100 * time.Millisecond)
	default:
	}
}

func (f *Fish) Clear() {
	clearUnicode := ' '

	for col := 0; col < len(f.model); col++ {
		for row := 0; row < len(f.model[col]); row++ {
			f.screen.SetContent(f.curX+row, f.curY+col, clearUnicode, nil, f.style)
		}
	}
}
