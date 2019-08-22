package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	b, err := ioutil.ReadFile(".xd.json")
	if err != nil {
		println("No .xd.json configuration file found in folder")
		return
	}

	m := map[string]string{}

	err = json.Unmarshal(b, &m)
	if err != nil {
		log.Fatalf("error: %v", err)
		return
	}

	commands := make([]cli.Command, 0)
	for k, v := range m {
		commands = append(commands, cli.Command{
			Name: k + ": " + v,
			Action: func(c *cli.Context) error {
				runCmd(v)
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

func runCmd(cmd string) {
	words := strings.Fields(cmd)
	command := exec.Command(words[0], words[1:]...)
	stdout, err := command.Output()
	println("Running command: " + cmd)
	if err != nil {
		println(err.Error())
		return
	}
	print(string(stdout))
}
