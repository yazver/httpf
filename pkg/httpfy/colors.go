package httpfy

import (
	"strings"

	"github.com/gotidy/httpf/pkg/termite/color"
)

type HTTP struct {
	HeaderKey color.Color
	Status100 color.Color // 1xx statuses
	Status200 color.Color // 2xx statuses
	Status300 color.Color // 3xx statuses
	Status400 color.Color // 4xx statuses
	Status500 color.Color // 5xx statuses
}

type TokensFormat struct {
	Identifier color.Color
	Symbol     color.Color
	Brackets   color.Color
	String     color.Color
	Number     color.Color
	Bool       color.Color
	Null       color.Color

	Indent  int
	NewLine string
}

type Colors struct {
	Tokens TokensFormat
	HTTP   HTTP
}

func (c Colors) Identifier(s string) string {
	return color.Colorize(s, c.Tokens.Identifier)
}

func (c Colors) String(s string) string {
	return color.Colorize(s, c.Tokens.String)
}

func (c Colors) Number(s string) string {
	return color.Colorize(s, c.Tokens.Number)
}

func (c Colors) Bool(s string) string {
	return color.Colorize(s, c.Tokens.Bool)
}

func (c Colors) Null(s string) string {
	return color.Colorize(s, c.Tokens.Null)
}

func (c Colors) Symbol(s string) string {
	return color.Colorize(s, c.Tokens.Symbol)
}

func (c Colors) Brackets(s string) string {
	return color.Colorize(s, c.Tokens.Brackets)
}

func (c Colors) Unknown(s string) string {
	return s
}

func (c Colors) Indent(level int) string {
	return strings.Repeat(" ", c.Tokens.Indent*level)
}

func (c Colors) NewLine() string {
	return c.Tokens.NewLine
}

func (c Colors) StatusColor(status string) color.Color {
	if len(status) == 0 {
		return color.Undefined
	}
	switch status[:1] {
	case "1":
		return c.HTTP.Status100
	case "2":
		return c.HTTP.Status200
	case "3":
		return c.HTTP.Status300
	case "4":
		return c.HTTP.Status400
	case "5":
		return c.HTTP.Status500
	default:
		return color.Undefined
	}
}

func (c Colors) Status(s, status string) string {
	if len(status) == 0 {
		return s
	}
	switch status[:1] {
	case "1":
		return color.Colorize(s, c.HTTP.Status100)
	case "2":
		return color.Colorize(s, c.HTTP.Status200)
	case "3":
		return color.Colorize(s, c.HTTP.Status300)
	case "4":
		return color.Colorize(s, c.HTTP.Status400)
	case "5":
		return color.Colorize(s, c.HTTP.Status500)
	default:
	}
	return s
}

func (c Colors) Header(s, status string) string {
	return color.Colorize(s, c.HTTP.HeaderKey)
}

var DefaultColors = Colors{
	Tokens: TokensFormat{
		Identifier: color.BrightBlue,
		String:     color.White,
		Number:     color.BrightMagenta,
		Symbol:     color.BrightRed,
		Brackets:   color.White,
		Bool:       color.BrightYellow,
		Null:       color.BrightBlack,
		Indent:     4,
		NewLine:    "\n",
	},
	HTTP: HTTP{
		HeaderKey: color.BrightGreen,
		Status100: color.White,
		Status200: color.Green,
		Status300: color.Yellow,
		Status400: color.Red,
		Status500: color.Blue,
	},
}

type MonoColors struct {
	indent  int
	newline string
}

func (c MonoColors) Identifier(s string) string {
	return s
}

func (c MonoColors) Symbol(s string) string {
	return s
}

func (c MonoColors) Brackets(s string) string {
	return s
}

func (c MonoColors) String(s string) string {
	return s
}

func (c MonoColors) Number(s string) string {
	return s
}

func (c MonoColors) Bool(s string) string {
	return s
}

func (c MonoColors) Null(s string) string {
	return s
}

func (c MonoColors) Unknown(s string) string {
	return s
}

func (c MonoColors) Indent(level int) string {
	return strings.Repeat(" ", c.indent*level)
}

func (c MonoColors) NewLine() string {
	return c.newline
}
