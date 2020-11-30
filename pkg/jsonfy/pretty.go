// Package prettyjson provides JSON pretty print.
package jsonfy

import (
	"io"

	"github.com/gotidy/httpf/pkg/fy"
	"github.com/pkg/json"
)

func newline(w fy.Writer, depth int) error {
	if err := w.NewLine(); err != nil {
		return err
	}
	if err := w.Indent(depth); err != nil {
		return err
	}
	return nil
}

func Format(src io.Reader, w fy.Writer) error {
	scan := json.NewScanner(src)

	needIndent := false
	depth := 0
	savedToken := ""
	for {
		tok := string(scan.Next())
		if err := scan.Error(); err != nil {
			return err
		}

		if len(tok) == 0 {
			break
		}
		tokID := tok[0]

		if savedToken != "" {
			switch {
			case tokID == ':':
				if err := w.Identifier(savedToken); err != nil {
					return err
				}
			default:
				if err := w.String(savedToken); err != nil {
					return err
				}
			}
			savedToken = ""
		}

		if needIndent && tokID != json.ObjectEnd && tokID != json.ArrayEnd {
			needIndent = false
			depth++
			if err := newline(w, depth); err != nil {
				return err
			}
		}

		// Add spacing around real punctuation.
		switch t := tokID; t {
		case '{', '[':
			// delay indent so that empty object and array are formatted as {} and [].
			needIndent = true
			if err := w.Brackets(string(tok)); err != nil {
				return err
			}

		case ',':
			if err := w.Symbol(tok); err != nil {
				return err
			}
			if err := newline(w, depth); err != nil {
				return err
			}
		case ':':
			if err := w.Symbol(": "); err != nil {
				return err
			}
		case '}', ']':
			if needIndent {
				// suppress indent in empty object/array
				needIndent = false
			} else {
				depth--

				if err := newline(w, depth); err != nil {
					return err
				}
			}
			if err := w.Brackets(tok); err != nil {
				return err
			}

		case '"':
			savedToken = tok
		case '-', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			if err := w.Number(tok); err != nil {
				return err
			}
		case json.True, json.False:
			if err := w.Bool(tok); err != nil {
				return err
			}
		case json.Null:
			if err := w.Null(tok); err != nil {
				return err
			}
		default:
			if err := w.Unknown((tok)); err != nil {
				return err
			}
		}
	}
	return nil
}
