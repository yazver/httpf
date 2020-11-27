package color

import "strconv"

type Escape = Color

// Screen and cursor operations.

// CursorBlinking enable/disable Blinking.
func CursorBlinking(b bool) Escape {
	if b {
		return "\u001b[?12h"
	}
	return "\u001b[?12l"
}

// CursorVisibility enable/disable Blinking.
func CursorVisibility(b bool) Escape {
	if b {
		return "\u001b[?25h"
	}
	return "\u001b[?25l"
}

// Move puts the cursor at line L and column C.
func CursorTo(line, column int) Escape {
	return Escape("\u001b[" + strconv.Itoa(line) + ";" + strconv.Itoa(column) + "H") // \u001b[<L>;<C>H or \u001b[<L>;<C>f
}

// CursorUp move the cursor up N lines.
func CursorUp(n int) Escape {
	return Escape("\u001b[" + strconv.Itoa(n) + "A") // \u001b[<N>A
}

// CursorDown move the cursor down N lines.
func CursorDown(n int) Escape {
	return Escape("\u001b[" + strconv.Itoa(n) + "B") // \u001b[<N>B
}

// Forward move the cursor forward N columns.
func CursorForward(n int) Escape {
	return Escape("\u001b[" + strconv.Itoa(n) + "C") // \u001b[<N>C
}

// Backward move the cursor backward N columns.
func CursorBackward(n int) Escape {
	return Escape("\u001b[" + strconv.Itoa(n) + "D") // \u001b[<N>D
}

// ScrollUp scroll text up by <n>. Also known as pan down, new lines fill in from the bottom of the screen.
func ScrollUp(n int) Escape {
	return Escape("\u001b[" + strconv.Itoa(n) + "S") // \u001b[<N>A
}

// ScrollDown scroll text down by <n>. Also known as pan up, new lines fill in from the top of the screen.
func ScrollDown(n int) Escape {
	return Escape("\u001b[" + strconv.Itoa(n) + "T") // \u001b[<N>B
}

const (
	// CursorBlinkingOn enable/disable Blinking.
	CursorBlinkingOn Escape = "\u001b[?12h"
	// CursorBlinkingOff disable Blinking.
	CursorBlinkingOff Escape = "\u001b[?12l"
	// CursorShow show the cursor.
	CursorShow Escape = "\u001b[?25h"
	// CursorHide hide the cursor.
	CursorHide Escape = "\u001b[?25l"

	// Clear the screen, move to (0,0).
	Clear Escape = "\u001b[2J"
	// Clears the screen from the cursor to the end of the screen.
	ClearDown Escape = "\u001b[J"
	// Clears the screen from the cursor to the start of the screen.
	ClearUp Escape = "\u001b[1J"
	// ClearLine clears all characters on the line.
	ClearLine Escape = "\u001b[2K"
	// ClearLineRight clears all characters from the cursor position to the end of the line.
	ClearLineRight Escape = "\u001b[K"
	// ClearLineLeft clears all characters from the cursor position to the start of the line.
	ClearLineLeft Escape = "\u001b[1K"
	// Save cursor position.
	SavePos Escape = "\u001b[s"
	// Restore cursor position.
	RestorePos Escape = "\u001b[u"
)
