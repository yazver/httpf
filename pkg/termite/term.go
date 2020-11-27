package termite

import (
	"io"
	"strings"

	"github.com/gotidy/httpf/pkg/termite/color"
)

type bytes = []byte

type Mite struct {
	colors    []bytes
	w         io.Writer
	noEscapes bool

	err error
}

func New(w io.Writer) *Mite {
	return &Mite{w: w, colors: []bytes{bytes(color.Reset)}}
}

func (c *Mite) lastColor() []byte {
	return c.colors[len(c.colors)-1]
}

func (c *Mite) Push(colors ...color.Color) {
	b := strings.Builder{}
	for _, color := range colors {
		_, _ = b.WriteString(string(color))
	}
	c.colors = append(c.colors, append(c.lastColor(), bytes(b.String())...))
	c.apply()
}

func (c *Mite) Pop() {
	if len(c.colors) > 1 {
		c.colors = (c.colors)[:len(c.colors)-1]
	}
	c.apply()
}

func (c *Mite) Reset() {
	c.colors = c.colors[:1]
	c.apply()
}

func (c Mite) apply() {
	c.err = c.WriteRawEscape(c.colors[len(c.colors)-1])
}

func (c *Mite) WriteRawEscape(b []byte) (err error) {
	if c.noEscapes {
		return nil
	}
	_, err = c.Write(b)
	return err
}

func (c *Mite) WriteEscape(e color.Escape) (err error) {
	return c.WriteRawEscape(bytes(e))
}

func (c *Mite) Write(p []byte) (n int, err error) {
	if err := c.err; err != nil {
		c.err = nil
		return 0, err
	}
	return c.w.Write(p)
}

func (c *Mite) WriteWithEscape(p []byte, e color.Escape) (n int, err error) {
	if err := c.err; err != nil {
		c.err = nil
		return 0, err
	}

	if err := c.WriteEscape(e); err != nil {
		return 0, err
	}

	if n, err = c.w.Write(p); err != nil {
		return n, err
	}

	c.apply()
	err = c.err
	return n, err
}

func (c *Mite) WriteString(s string) (n int, err error) {
	return c.Write(bytes(s))
}

// CursorBlinking enable/disable Blinking.
func (c *Mite) CursorBlinking(b bool) (err error) {
	return c.WriteEscape(color.CursorBlinking(b))
}

// CursorVisibility enable/disable Blinking.
func (c *Mite) CursorVisibility(b bool) (err error) {
	return c.WriteEscape(color.CursorVisibility(b))
}

func (c *Mite) CursorTo(line, column int) (err error) {
	return c.WriteEscape(color.CursorTo(line, column))
}

func (c *Mite) CursorUp(n int) (err error) {
	return c.WriteEscape(color.CursorUp(n))
}

func (c *Mite) CursorDown(n int) (err error) {
	return c.WriteEscape(color.CursorDown(n))
}

func (c *Mite) CursorForward(n int) (err error) {
	return c.WriteEscape(color.CursorForward(n))
}

func (c *Mite) CursorBackward(n int) (err error) {
	return c.WriteEscape(color.CursorBackward(n))
}

func (c *Mite) ScrollUp(n int) (err error) {
	return c.WriteEscape(color.ScrollUp(n))
}

func (c *Mite) ScrollDown(n int) (err error) {
	return c.WriteEscape(color.ScrollDown(n))
}

// CursorBlinkingOn enable/disable Blinking.
func (c *Mite) CursorBlinkingOn() (err error) {
	return c.WriteEscape(color.CursorBlinkingOn)
}

// CursorBlinkingOff disable Blinking.
func (c *Mite) CursorBlinkingOff() (err error) {
	return c.WriteEscape(color.CursorBlinkingOff)
}

// CursorShow show the cursor.
func (c *Mite) CursorShow() (err error) {
	return c.WriteEscape(color.CursorShow)
}

// CursorHide hide the cursor.
func (c *Mite) CursorHide() (err error) {
	return c.WriteEscape(color.CursorHide)
}

// Clear the screen, move to (0,0).
func (c *Mite) Clear() (err error) {
	return c.WriteEscape(color.Clear)
}

// Clears the screen from the cursor to the end of the screen.
func (c *Mite) ClearDown() (err error) {
	return c.WriteEscape(color.ClearDown)
}

// Clears the screen from the cursor to the start of the screen.
func (c *Mite) ClearUp() (err error) {
	return c.WriteEscape(color.ClearUp)
}

// ClearLine clears all characters on the line.
func (c *Mite) ClearLine() (err error) {
	return c.WriteEscape(color.ClearLine)
}

// ClearLineRight clears all characters from the cursor position to the end of the line.
func (c *Mite) ClearLineRight() (err error) {
	return c.WriteEscape(color.ClearLineRight)
}

// ClearLineLeft clears all characters from the cursor position to the start of the line.
func (c *Mite) ClearLineLeft() (err error) {
	return c.WriteEscape(color.ClearLineLeft)
}

// Save cursor position.
func (c *Mite) SavePos() (err error) {
	return c.WriteEscape(color.SavePos)
}

// Restore cursor position.
func (c *Mite) RestorePos() (err error) {
	return c.WriteEscape(color.RestorePos)
}
