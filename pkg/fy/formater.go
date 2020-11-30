package fy

import "io"

type Prettifier interface {
	Format(r io.Reader, formatter Formatter) error
}

// Formatter provides the interface to format strings.
type Formatter interface {
	Identifier(s string) string
	Symbol(s string) string
	String(s string) string
	Brackets(s string) string
	Number(s string) string
	Bool(s string) string
	Null(s string) string
	Unknown(s string) string
	NewLine() string
	Indent(level int) string
}

// Writer provides the interface to formated output.
type Writer interface {
	Identifier(s string) error
	Symbol(s string) error
	Brackets(s string) error
	String(s string) error
	Number(s string) error
	Bool(s string) error
	Null(s string) error
	Unknown(s string) error
	NewLine() error
	Indent(level int) error
}

type FormatedWriter struct {
	w io.Writer
	f Formatter
}

func New(w io.Writer, f Formatter) *FormatedWriter {
	return &FormatedWriter{w: w, f: f}
}

func wrapResult(n int, err error) error {
	return err
}

func (f FormatedWriter) Identifier(s string) error {
	return wrapResult(f.w.Write([]byte(f.f.Identifier(s))))
}

func (f FormatedWriter) Symbol(s string) error {
	return wrapResult(f.w.Write([]byte(f.f.Symbol(s))))
}

func (f FormatedWriter) Brackets(s string) error {
	return wrapResult(f.w.Write([]byte(f.f.Brackets(s))))
}

func (f FormatedWriter) String(s string) error {
	return wrapResult(f.w.Write([]byte(f.f.String(s))))
}

func (f FormatedWriter) Number(s string) error {
	return wrapResult(f.w.Write([]byte(f.f.Number(s))))
}

func (f FormatedWriter) Bool(s string) error {
	return wrapResult(f.w.Write([]byte(f.f.Bool(s))))
}

func (f FormatedWriter) Null(s string) error {
	return wrapResult(f.w.Write([]byte(f.f.Null(s))))
}

func (f FormatedWriter) Unknown(s string) error {
	return wrapResult(f.w.Write([]byte(f.f.Unknown(s))))
}

func (f FormatedWriter) NewLine() error {
	return wrapResult(f.w.Write([]byte(f.f.NewLine())))
}

func (f FormatedWriter) Indent(level int) error {
	return wrapResult(f.w.Write([]byte(f.f.Indent(level))))
}
