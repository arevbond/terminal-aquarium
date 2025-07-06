package main

import (
	"github.com/gdamore/tcell"
	"log"
)

var fish = "><_>"

func main() {
	backgroundStyle := tcell.StyleDefault.Foreground(tcell.ColorRed).Background(tcell.ColorBlue)

	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}

	if err = s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}

	s.SetStyle(backgroundStyle)

	var start int

	for {
		s.Show()

		ev := s.PollEvent()

		switch ev := ev.(type) {
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyCtrlC {
				return
			}

			if ev.Key() == tcell.KeyCtrlA {
				s.Clear()
				drawText(s, start, start+len(fish)+1, 1, len(fish)+1, tcell.StyleDefault, fish)
				start++
			}
		}

	}
}

func drawText(s tcell.Screen, x1, x2, y1, y2 int, style tcell.Style, text string) {
	row := y1
	col := x1
	for _, ch := range text {
		s.SetContent(col, row, ch, nil, style)
		col++

		if col >= x2 {
			row++
			col = x1
		}

		if row > y2 {
			break
		}
	}
}
