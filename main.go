package main

import (
	"fmt"
	"github.com/gdamore/tcell"
	"log"
)

//var fish = "><_>"

var bigFish = `
o           //        +
 o        //////    ++
  .    @))))))))))+++
   .}:<<)))))))))))++
       <<\\)))))) +++
          \\   \\   ++
                      +`

var bigFish2 = `
     |\    o
    |  \    o
|\ /    .\ o
| |       (
|/ \     /
    |  /
     |/
`

var smallFishSlice = []string{
	"  _ ",
	"><_>",
}

var fishForward = []string{
	"     |\\    o",
	"    |  \\    o",
	"|\\ /    .\\ o",
	"| |       (",
	"|/ \\     /",
	"    |  /",
	"     |/",
}

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

	var start int

	quit := func() {
		maybePanic := recover()
		s.Fini()
		if maybePanic != nil {
			panic(maybePanic)
		}
	}
	defer quit()

	fish := fishForward

	isForward := true

	for {
		s.Show()

		ev := s.PollEvent()

		switch ev := ev.(type) {
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyCtrlC {
				return
			}
			s.Clear()

			drawFish(s, start, tcell.StyleDefault.Foreground(tcell.ColorBlue), fish)

			if isForward {
				start++
			} else {
				start--
			}

			if start == 50 {
				isForward = false
				fish = reverse(fish)
			}
		}
	}
}

func drawFish(s tcell.Screen, startX int, style tcell.Style, fish []string) {
	for col := 0; col < len(fish); col++ {
		for row := 0; row < len(fish[col]); row++ {
			s.SetContent(startX+row, col, rune(fish[col][row]), nil, style)
		}
	}
}

func reverse(fish []string) []string {
	result := make([]string, 0, len(fish))
	for _, str := range fish {
		newStr := ""
		for r := len(str) - 1; r >= 0; r-- {
			newStr += string(str[r])
		}
		result = append(result, newStr)
	}

	return result
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
