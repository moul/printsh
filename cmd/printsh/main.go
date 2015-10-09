package main

import (
	"os"
	"path"

	"github.com/codegangsta/cli"
	"github.com/moul/printsh"
)

// main is the entrypoint
func main() {
	app := cli.NewApp()
	app.Name = path.Base(os.Args[0])
	app.Author = "Manfred Touron"
	app.Email = "https://github.com/moul"
	app.Version = "1.0.0"
	app.Usage = "Run command in a shell and print debugging information"

	app.Action = action
	app.Run(os.Args)
}

func action(c *cli.Context) {
	psh := printsh.New()

	if len(c.Args()) > 0 {
	} else {
		psh.Name = "<INPUT>"
		psh.AddStream(os.Stdin, "stdin")
	}

	psh.Start()
}
