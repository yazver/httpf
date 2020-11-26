package httpfy

import (
	"fmt"
	"io"
	"strings"
)

type TokenFormater interface {
	Identifier(s string) string
	String(s string) string
	Number(s string) string
	Symbol(s string) string
}

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

var formaters = map[ContentType]func(dst io.Writer, src io.Reader, formater TokenFormater) error{
	ContentHTML:    htmlFormater,
	ContentXML:     xmlFormater,
	ContentJSON:    jsonFormater,
	ContentUnknown: unknownFormater,
}

func jsonFormater(dst io.Writer, src io.Reader, formater TokenFormater) error {
	_, err := io.Copy(dst, src)
	return err
}

func htmlFormater(dst io.Writer, src io.Reader, formater TokenFormater) error {
	_, err := io.Copy(dst, src)
	return err
}

func xmlFormater(dst io.Writer, src io.Reader, formater TokenFormater) error {
	_, err := io.Copy(dst, src)
	return err
}

func unknownFormater(dst io.Writer, src io.Reader, formater TokenFormater) error {
	_, err := io.Copy(dst, src)
	return err
}

func Format(dst io.Writer, src io.Reader, typ ContentType, formater TokenFormater) error {
	f := formaters[typ]
	if f == nil {
		return fmt.Errorf("formater for context type «%s» is not founded", typ)
	}
	return f(dst, src, formater)
}
