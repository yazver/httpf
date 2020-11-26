package httpfy

type Color string

const (
	Black   Color = "\u001b[30m"
	Red     Color = "\u001b[31m"
	Green   Color = "\u001b[32m"
	Yellow  Color = "\u001b[33m"
	Blue    Color = "\u001b[34m"
	Magenta Color = "\u001b[35m"
	Cyan    Color = "\u001b[36m"
	White   Color = "\u001b[37m"

	BrightBlack   Color = "\u001b[30;1m"
	BrightRed     Color = "\u001b[31;1m"
	BrightGreen   Color = "\u001b[32;1m"
	BrightYellow  Color = "\u001b[33;1m"
	BrightBlue    Color = "\u001b[34;1m"
	BrightMagenta Color = "\u001b[35;1m"
	BrightCyan    Color = "\u001b[36;1m"
	BrightWhite   Color = "\u001b[37;1m"

	Reset Color = "\u001b[0m"

	Bold      Color = "\u001b[1m"
	Underline Color = "\u001b[4m"
	Reversed  Color = "\u001b[7m"

	Undefined Color = ""
)

type HTTP struct {
	HeaderName Color
	Status100  Color // 1xx statuses
	Status200  Color // 2xx statuses
	Status300  Color // 3xx statuses
	Status400  Color // 4xx statuses
	Status500  Color // 5xx statuses
}

type TokensColors struct {
	Identifier Color
	String     Color
	Number     Color
	Symbol     Color
	Brackets   Color
}

type Colors struct {
	Tokens TokensColors
	HTTP   HTTP
}

func (c Colors) Identifier(s string) string {
	return colorize(s, c.Tokens.Identifier)
}

func (c Colors) String(s string) string {
	return colorize(s, c.Tokens.String)
}

func (c Colors) Number(s string) string {
	return colorize(s, c.Tokens.Number)
}

func (c Colors) Symbol(s string) string {
	return colorize(s, c.Tokens.Symbol)
}

func (c Colors) Brackets(s string) string {
	return colorize(s, c.Tokens.Brackets)
}

func (c Colors) StatusColor(status string) Color {
	if len(status) == 0 {
		return Undefined
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
		return Undefined
	}
}

func (c Colors) Status(s, status string) string {
	if len(status) == 0 {
		return s
	}
	switch status[:1] {
	case "1":
		return colorize(s, c.HTTP.Status100)
	case "2":
		return colorize(s, c.HTTP.Status200)
	case "3":
		return colorize(s, c.HTTP.Status300)
	case "4":
		return colorize(s, c.HTTP.Status400)
	case "5":
		return colorize(s, c.HTTP.Status500)
	default:
	}
	return s
}

func (c Colors) Header(s, status string) string {
	return colorize(s, c.HTTP.HeaderName)
}

var DefaultColors = Colors{
	Tokens: TokensColors{
		Identifier: Green,
		String:     White,
		Number:     Magenta,
		Symbol:     BrightRed,
		Brackets:   White,
	},
	HTTP: HTTP{
		Status100: White,
		Status200: Green,
		Status300: Yellow,
		Status400: Red,
		Status500: Blue,
	},
}

func colorize(s string, color Color) string {
	return string(color) + s + string(color)
}

type MonoColors struct {
}

func (c MonoColors) Identifier(s string) string {
	return s
}

func (c MonoColors) String(s string) string {
	return s
}

func (c MonoColors) Number(s string) string {
	return s
}

func (c MonoColors) Symbol(s string) string {
	return s
}

func (c MonoColors) Brackets(s string) string {
	return s
}
