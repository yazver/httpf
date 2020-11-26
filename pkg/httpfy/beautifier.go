package httpfy

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net/textproto"
	"strings"
)

type ErrIO struct {
	err error
}

func NewErrIO(err error) error {
	if err == nil {
		return nil
	}
	return ErrIO{err: err}
}

func (e ErrIO) Unwrap() error {
	return e.err
}

func (e ErrIO) Error() string {
	return e.err.Error()
}

func IsIOError(err error) bool {
	_, ok := err.(ErrIO)
	return ok
}

type ErrUnrecognizedFormat struct {
	error
}

func NewErrUnrecognizedFormat(line string) error {
	if len(line) > 80 {
		line = line[:80]
	}
	return ErrUnrecognizedFormat{error: fmt.Errorf("unrecognized format of «%s»", strings.TrimSpace(line))}
}

func IsUnrecognizedFormat(err error) bool {
	_, ok := err.(ErrUnrecognizedFormat)
	return ok
}

type Option func(h *HTTPfy)

func Colorize(c bool) Option {
	return func(h *HTTPfy) {
		h.colorize = c
	}
}

func Beautify(b bool) Option {
	return func(h *HTTPfy) {
		h.beautify = b
	}
}

func WithColors(c Colors) Option {
	return func(h *HTTPfy) {
		h.Colors = c
	}
}

type HTTPfy struct {
	colorize bool
	beautify bool

	src *bufio.Reader
	dst io.Writer

	Colors Colors

	contentType    ContentType
	contentCharSet string

	eof bool
}

func New(dst io.Writer, src io.Reader, opts ...Option) *HTTPfy {
	h := &HTTPfy{
		colorize: false,
		beautify: false,
		src:      bufio.NewReader(src),
		dst:      dst,
		Colors:   DefaultColors,
	}
	for _, opt := range opts {
		opt(h)
	}
	return h
}

func (h *HTTPfy) writeString(s string, color Color) error {
	if h.colorize && color != Undefined {
		s = string(color) + s + string(Reset)
	}
	_, err := h.dst.Write([]byte(s))
	return NewErrIO(err)
}

func (h *HTTPfy) writeBytes(b []byte, color Color) error {
	return h.writeString(string(b), color)
}

func (h *HTTPfy) readLine() ([]byte, error) {
	b, err := h.src.ReadBytes('\n')
	if err == io.EOF {
		h.eof = true
		err = nil
	}
	return b, NewErrIO(err)
}

func (h *HTTPfy) header() (empty bool, err error) {
	b, err := h.readLine()
	if err != nil {
		return false, err
	}

	if isEmpty(b) {
		return true, h.writeBytes(b, Undefined)
	}

	if spaceStarted(b) {
		return false, h.writeBytes(b, Undefined)
	}

	i := bytes.IndexByte(b, ':')
	if i < 0 {
		return false, h.writeBytes(b, Undefined)
	}

	key := b[:i]
	val := bytes.TrimPrefix(b[i:], []byte(":"))

	// Output header name
	if h.beautify {
		key = []byte(textproto.CanonicalMIMEHeaderKey(string(key)))
	}
	if err := h.writeBytes(key, h.Colors.Tokens.Identifier); err != nil { // HTTP.HeaderName
		return false, err
	}

	// Output separator
	if err := h.writeString(": ", h.Colors.Tokens.Symbol); err != nil {
		return false, err
	}

	// Output header value
	if err := h.writeBytes(bytes.TrimSpace(val), h.Colors.Tokens.String); err != nil {
		return false, err
	}

	// Get content type
	if strings.ToLower(string(key)) == "content-type" {
		parts := strings.Split(string(val), ";")
		if len(parts) > 0 {
			h.contentType = GetContentType(parts[0])
		}
		if len(parts) > 1 {
			h.contentCharSet = strings.ToLower(strings.TrimSpace(parts[1]))
		}
	}

	if err := h.writeString("\n", Undefined); err != nil {
		return false, err
	}

	return false, nil
}

func (h *HTTPfy) protocol() (continueStatus bool, err error) {
	b, err := h.readLine()
	if err != nil {
		return false, err
	}

	if !bytes.HasPrefix(b, []byte("HTTP")) {
		_ = h.writeBytes(b, Undefined)
		return false, NewErrUnrecognizedFormat(string(b))
	}

	parts := strings.Fields(string(b))
	if len(parts) < 2 {
		_ = h.writeBytes(b, Undefined)
		return false, NewErrUnrecognizedFormat(string(b))
	}

	if err := h.writeString(parts[0], Undefined); err != nil {
		return false, err
	}

	statusColor := h.Colors.StatusColor(parts[1])
	for i := 1; i < len(parts); i++ {
		if err := h.writeString(" ", Undefined); err != nil {
			return false, err
		}
		if err := h.writeString(parts[i], statusColor); err != nil {
			return false, err
		}
	}

	if err := h.writeString("\n", Undefined); err != nil {
		return false, err
	}

	return parts[1] == "100", nil
}

func (h *HTTPfy) Do() (err error) {
	defer func() {
		if err != nil && !IsIOError(err) {
			_, _ = io.Copy(h.dst, h.src)
		}
	}()

header:
	for {
		cont, err := h.protocol()
		if err != nil {
			return err
		}
		if h.eof {
			break
		}
		for {
			empty, err := h.header()
			if err != nil {
				return err
			}
			switch {
			case h.eof:
				return nil
			case empty && cont: // If the header is continue, then repeat header reading.
				break
			case empty && !cont: // If the header is not continue.
				break header
			}
		}
	}

	var tokenFormater TokenFormater = MonoColors{}
	if h.colorize {
		tokenFormater = h.Colors
	}
	return Format(h.dst, h.src, h.contentType, tokenFormater)
}

func isEmpty(b []byte) bool {
	return len(bytes.TrimSpace(b)) == 0
}

func spaceStarted(b []byte) bool {
	return !isEmpty(b) && (b[0] == ' ' || b[0] == '\t')
}
