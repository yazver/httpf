package httpfy

import (
	"io"
	"strings"

	"github.com/gotidy/httpf/pkg/fy"
	"github.com/gotidy/httpf/pkg/jsonfy"
)

type ContentType string

const (
	ContentHTML    ContentType = "text/html"
	ContentXML     ContentType = "text/xml"
	ContentJSON    ContentType = "application/json"
	ContentUnknown ContentType = ""
)

var contentTypes = map[string]ContentType{
	"text/html":        ContentHTML,
	"text/xml":         ContentXML,
	"application/xml":  ContentXML,
	"application/json": ContentJSON,
}

func GetContentType(contentType string) ContentType {
	return contentTypes[strings.ToLower(strings.TrimSpace(contentType))]
}

var formatters = map[ContentType]func(dst io.Writer, src io.Reader, formatter fy.Formatter) error{
	ContentHTML:    htmlFormatter,
	ContentXML:     xmlFormatter,
	ContentJSON:    jsonFormatter,
	ContentUnknown: unknownFormatter,
}

func jsonFormatter(dst io.Writer, src io.Reader, formatter fy.Formatter) error {
	// _, err := io.Copy(dst, src)
	return jsonfy.Format(src, fy.New(dst, formatter))
}

func htmlFormatter(dst io.Writer, src io.Reader, formatter fy.Formatter) error {
	_, err := io.Copy(dst, src)
	return err
}

func xmlFormatter(dst io.Writer, src io.Reader, formatter fy.Formatter) error {
	_, err := io.Copy(dst, src)
	return err
}

func unknownFormatter(dst io.Writer, src io.Reader, formatter fy.Formatter) error {
	_, err := io.Copy(dst, src)
	return err
}

func Format(dst io.Writer, src io.Reader, typ ContentType, formatter fy.Formatter) error {
	f := formatters[typ]
	if f == nil {
		f = unknownFormatter
	}
	if err := f(dst, src, formatter); err != nil && err != io.EOF {
		return err
	}
	return nil
}
