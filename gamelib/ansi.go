package gamelib

type AnsiMode byte

const (
	NoAnsi AnsiMode = iota
	Letters
	Numbers
	Rgb
	Hex1
	Hex2
	Name
)

type AnsiGround struct {
	text  string
	mode  AnsiMode
	clear bool
	data  string
	codes string
}

type AnsiData struct {
	bits    byte
	offBits byte
	fg      AnsiGround
	bg      AnsiGround
	reset   bool
}

type MarkupMode byte

const (
	NoMarkup MarkupMode = iota
	Color
	MXP
)

type Markup struct {
	mode   MarkupMode
	ansi   AnsiData
	idx    uint16
	parent *Markup
	html   string
}

type MarkupChar struct {
	info Markup
	char byte
}

type MarkupString struct {
	clean   string
	source  string
	infos   []Markup
	markups []MarkupChar
}
