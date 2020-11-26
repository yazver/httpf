package main

import (
	"log"
	"os"

	"github.com/gotidy/httpf/pkg/httpfy"
	"github.com/urfave/cli/v2"
)

var Version = "1.0.0"

const description = `httpf is a tool for processing HTTP inputs, applying the formation and colorization to
HTTP response from standard input and producing it on standard output. 
Example: curl  -i https://www.google.com/ | httpf -c"`

func main() {
	app := &cli.App{
		Name:                 "httpf",
		Usage:                "commandline HTTP response beautifier",
		Version:              Version,
		UsageText:            "httpf [-c] [-B] [-i FILE] [-o FILE]",
		EnableBashCompletion: true,
		Authors:              []*cli.Author{{Name: "Evgeny Safonov"}},
		Description:          description,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "no-beautify",
				Aliases: []string{"B"},
				Usage:   "Off beautify output.",
			},
			&cli.BoolFlag{
				Name:    "colorize",
				Aliases: []string{"c"},
				Usage:   "Colorize output.",
			},
			&cli.PathFlag{
				Name:    "input",
				Aliases: []string{"i"},
				Usage:   "To get data from the file insted of STDIN.",
			},
			&cli.PathFlag{
				Name:    "output",
				Aliases: []string{"o"},
				Usage:   "To save data to the file insted of STDOUT.",
			},
		},
		Action: func(c *cli.Context) error {
			var err error

			in := os.Stdin
			out := os.Stdout

			if c.IsSet("input") {
				if in, err = os.Open(c.String("input")); err != nil {
					log.Fatalf("Unable to open the source file «%s»", err)
				}
			}

			if c.IsSet("output") {
				if out, err = os.Open(c.String("output")); err != nil {
					log.Fatalf("Unable to open the source file «%s»", err)
				}
			}

			fy := httpfy.New(out, in, httpfy.Beautify(!c.Bool("no-beautify")), httpfy.Colorize(c.Bool("colorize")))

			if err := fy.Do(); err != nil {
				log.Fatalf("HTTP Beautify failed: %s", err)
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
