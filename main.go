package main

import (
	"github.com/gdamore/tcell"
	"log"
	"log/slog"
	"os"
)

func main() {
	logger := setupLogger()

	app := NewApp(logger)

	finish := func() {
		maybePanic := recover()
		app.screen.Fini()
		if maybePanic != nil {
			panic(maybePanic)
		}
	}
	defer finish()

	fishStyle := tcell.StyleDefault.Foreground(tcell.ColorBlue).Background(tcell.ColorWhite)

	decorations := make([]*Decoration, 0)

	decorations = append(decorations, NewDecoration(sea, 0, 5, app.screen, fishStyle))

	initialFishes := app.generateFishes(fishStyle)

	for _, fish := range initialFishes {
		go func() {
			go fish.Swim()
			//<-fish.endSwim
		}()
	}

	for _, decoration := range decorations {
		decoration.Draw()

	}

	if err := app.Run(); err != nil {
		logger.Error("error while end app", slog.Any("error", err))
	}
}

func setupLogger() *slog.Logger {
	file, err := os.OpenFile("logs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	logHandler := slog.NewTextHandler(file, &slog.HandlerOptions{Level: slog.LevelDebug})
	logger := slog.New(logHandler)

	return logger
}

//func updateFishes(fishes []*Fish) []*Fish {
//	return nil
//}
//
//func addNewFish(x, y int, addNewFishCh chan<- [2]int) {
//	fish := generateRandomFish(x)
//	go func() {
//		go fish.Swim()
//		<-fish.endSwim
//
//		addNewFishCh <- [2]int{fish.curX, fish.curY}
//	}()
//}
//
//func generateRandomFish(y int) *Fish {
//	if y == 0 {
//		return NewWhale(1 + rand.Intn(3))
//	}
//}
