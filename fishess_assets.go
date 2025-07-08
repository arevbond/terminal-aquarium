package main

import "github.com/gdamore/tcell"

var fishColors = []tcell.Color{tcell.ColorBlue, tcell.ColorBlue, tcell.ColorRed, tcell.ColorOrange, tcell.ColorOrange}

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

const sea = "~^"

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
