package main

import (
	"log"
	"os"

	"github.com/ia3andy/xd/cmds"
	"github.com/ia3andy/xd/configs"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	config, err := configs.ReadConfig()
	if err != nil {
		println(err.Error())
		return
	}

	app.Commands = cmds.GenCommands(config)

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
