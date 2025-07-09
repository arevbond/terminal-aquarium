package main

import "github.com/gdamore/tcell"

var fishColors = []tcell.Color{
	tcell.ColorOrange,
	//tcell.ColorYellow,
	tcell.ColorRed,
	//tcell.ColorPink,
	//tcell.ColorWhite,
	tcell.ColorGold,
	tcell.ColorCoral,
	tcell.ColorTomato,
	tcell.ColorHotPink,
	tcell.ColorLightSalmon,
}

var fishModelsForward = [][]string{fishForward, fishForward2, fishForward3}
var fishModelsBackward = [][]string{fishBackward, fishBackward2}

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

var fishForward3 = []string{
	"     .-\"\"L_",
	";`, /   ( o\\ ",
	"\\  ;    `, / ",
	";_/\"`.__.-\"",
}
