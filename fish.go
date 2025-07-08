package main

import (
	"github.com/gdamore/tcell"
	"log/slog"
	"math/rand"
	"time"
)

func (a *App) generateFishes(fishStyle tcell.Style) []*Fish {
	fishes := make([]*Fish, 0)

	//fishes = append(fishes, NewFish(whaleBackward, false, HighSpeed, 100, 0, a.screen, tcell.StyleDefault.Foreground(tcell.ColorBlue).Background(tcell.ColorWhite), a.log))

	fishes = append(fishes, a.NewWhaleFish(Speed(2+rand.Intn(2))))

	fishes = append(fishes, NewFishWithRandomColor(fishForward, true, MediumSpeed, rand.Intn(10), 7, a.screen, a.seaBackgroundColor, a.log))
	fishes = append(fishes, NewFishWithRandomColor(fishForward2, true, MediumSpeed, rand.Intn(10), 23, a.screen, a.seaBackgroundColor, a.log))
	fishes = append(fishes, NewFishWithRandomColor(fishForward, true, HighSpeed, rand.Intn(10), 37, a.screen, a.seaBackgroundColor, a.log))

	fishes = append(fishes, NewFishWithRandomColor(fishBackward2, false, LowSpeed, 100+rand.Intn(10), 15, a.screen, a.seaBackgroundColor, a.log))
	fishes = append(fishes, NewFishWithRandomColor(fishBackward, false, MediumSpeed, 100+rand.Intn(10), 30, a.screen, a.seaBackgroundColor, a.log))

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
	endSwim chan struct{}
}

func NewFish(model []string, swimForward bool, speed Speed, x, y int, screen tcell.Screen, style tcell.Style, logger *slog.Logger) *Fish {
	return &Fish{
		model: model,
		curX:  x, curY: y,
		speed:       speed,
		style:       style,
		screen:      screen,
		swimForward: swimForward,
		logger:      logger,
		endSwim:     make(chan struct{})}
}

func (a *App) NewWhaleFish(speed Speed) *Fish {
	colors := []tcell.Color{tcell.ColorBlue}
	color := colors[rand.Intn(len(colors))]
	return NewFish(whaleBackward, false, speed, a.width, 0, a.screen, tcell.StyleDefault.Foreground(color).Background(tcell.ColorWhite), a.log)
}

func NewFishWithRandomColor(model []string, swimForward bool, speed Speed, x, y int, screen tcell.Screen, background tcell.Color, logger *slog.Logger) *Fish {
	color := fishColors[rand.Intn(len(fishColors))]

	style := tcell.StyleDefault.Foreground(color).Background(background)

	return &Fish{
		model: model,
		curX:  x, curY: y,
		speed:       speed,
		style:       style,
		screen:      screen,
		swimForward: swimForward,
		logger:      logger,
		endSwim:     make(chan struct{})}
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
