package main

import (
	"fmt"
	"os"
	"os/exec"
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
		cmd := exec.Command(c.Args()[0], c.Args()[1:]...)
		// psh.AddStream(cmd.Stdout, "stdout")
		// psh.AddStream(cmd.Stderr, "stderr")
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		fmt.Println(cmd.Run())
	} else {
		psh.Name = "<INPUT>"
		psh.AddStream(os.Stdin, "stdin")
	}

	psh.Start()
}
