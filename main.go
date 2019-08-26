package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/ia3andy/xd/cmds"
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
)

func main() {

	app := cli.NewApp()

	b, err := ioutil.ReadFile(".xd.yml")
	if err != nil {
		println("No .xd.json configuration file found in folder")
		return
	}

	m := map[string]string{}

	err = yaml.Unmarshal(b, &m)
	if err != nil {
		log.Fatalf("error: %v", err)
		return
	}

	commands := make([]cli.Command, 0)
	for k, v := range m {
		parsedCmd := cmds.ParseCmd(v)

		commands = append(commands, cli.Command{
			Name:  k,
			Usage: "Execute the command: " + parsedCmd.Text,
			Flags: cmds.ReadFlags(parsedCmd),

			Action: func(c *cli.Context) error {
				params := cmds.ReadParams(parsedCmd, c)
				processed, err := cmds.ProcessCmd(parsedCmd, params)
				if err != nil {
					return err
				}
				cmds.RunCmd(processed)
				return nil
			},
		})
	}
	app.Commands = commands

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
