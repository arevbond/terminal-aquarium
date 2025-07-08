package main

import (
	"github.com/gdamore/tcell"
	"log/slog"
)

var fishForward = []string{
	"     |\\     ",
	"    |  \\     ",
	"|\\ /    .\\  ",
	"| |       (",
	"|/ \\     /",
	"    |  /",
	"     |/",
}

var fishBackward = []string{
	"             ,",
	"           .:/",
	"       ,,///;,   ,;/",
	"      o:::::::;;///",
	"     >::::::::;;\\\\\\",
	"       ''\\\\\\\\\\'\" ';\\",
	"          ';\\",
}

var whaleBackward = []string{
	"       .",
	"      \":\"",
	"    ___:____     |\"\\/\"|",
	"  ,'        `.    \\  /",
	"  |  O        \\___/  |",
	//"~^~^~^~^~^~^~^~^~^~^~^~^~",
}

func NewWhale(speed Speed, screen tcell.Screen, logger *slog.Logger) *Fish {
	style := tcell.StyleDefault.Foreground(tcell.ColorBlue).Background(tcell.ColorWhite)
	x, _ := screen.Size()
	return NewFish(whaleBackward, false, speed, x, 0, screen, style, logger)
}

var sea = []string{
	"~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^",
}

var wave = []string{
	"   ,(   ,(   ,(   ,(   ,(   ,(   ,(   ,(",
	"`-'  `-'  `-'  `-'  `-'  `-'  `-'  `-'  `",
}

var fishBackward2 = []string{
	"   _.-=-._     .-, ",
	" .'       \"-.,' / ",
	"(          _.  <",
	" `=.____.=\"  `._\\",
}

var fishForward2 = []string{
	"        |\\",
	" \\`._.-' `--.",
	" ) o o =[#]#]",
	" ) _o      -3",
	" /.' `-.,---' ",
	"       '",
}
