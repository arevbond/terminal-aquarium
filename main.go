package main

import (
	"fmt"
	"github.com/gdamore/tcell"
	"log"
	"math/rand"
	"time"
)

func main() {
	backgroundStyle := tcell.StyleDefault.Foreground(tcell.ColorRed).Background(tcell.ColorWhite)

	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}

	x, y := s.Size()
	fmt.Printf("w: %d, h: %d\n", x, y)

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

	fishes := make([]*Fish, 0)

	fishes = append(fishes, NewFish(whaleBackward, false, 100, 0, s, fishStyle))

	fishes = append(fishes, NewFish(fishForward, true, 0, 7, s, fishStyle))
	fishes = append(fishes, NewFish(fishForward, true, 5, 23, s, fishStyle))
	fishes = append(fishes, NewFish(fishForward, true, 3, 37, s, fishStyle))

	fishes = append(fishes, NewFish(fishBackward2, false, 100, 15, s, fishStyle))
	fishes = append(fishes, NewFish(fishBackward2, false, 100, 30, s, fishStyle))

	for _, fish := range fishes {
		go func() {
			for {
				fish.Move()

				time.Sleep(time.Duration(100+rand.Intn(500)) * time.Millisecond)
			}
		}()
	}

	for _, decoration := range decorations {
		go func() {
			for {
				decoration.Draw()

				time.Sleep(time.Duration(100+rand.Intn(1000)) * time.Millisecond)
			}
		}()
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

type Fish struct {
	model       []string
	x, y        int
	style       tcell.Style
	screen      tcell.Screen
	swimForward bool
}

func NewFish(model []string, swimForward bool, x, y int, screen tcell.Screen, style tcell.Style) *Fish {
	return &Fish{model: model, x: x, y: y, style: style, screen: screen, swimForward: swimForward}
}

func (f *Fish) Draw() {
	for col := 0; col < len(f.model); col++ {
		for row := 0; row < len(f.model[col]); row++ {
			f.screen.SetContent(f.x+row, f.y+col, rune(f.model[col][row]), nil, f.style)
		}
	}
}

func (f *Fish) Move() {
	f.Clear()

	if f.swimForward {
		f.x++
	} else {
		f.x--
	}

	f.Draw()

}

func (f *Fish) Clear() {
	clearUnicode := ' '

	for col := 0; col < len(f.model); col++ {
		for row := 0; row < len(f.model[col]); row++ {
			f.screen.SetContent(f.x+row, f.y+col, clearUnicode, nil, f.style)
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
