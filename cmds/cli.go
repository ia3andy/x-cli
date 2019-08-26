package cmds

import (
	"github.com/ia3andy/xd/configs"
	"github.com/urfave/cli"
)

// GenCommands generates the commands from the config
func GenCommands(config *configs.XDConfig) []cli.Command {
	commands := make([]cli.Command, 0)
	for k, v := range config.Scripts {
		parsedCmd := ParseCmd(v)

		commands = append(commands, cli.Command{
			Name:  k,
			Usage: "Execute the command: " + parsedCmd.Text,
			Flags: ReadFlags(parsedCmd),

			Action: func(c *cli.Context) error {
				params := ReadParams(parsedCmd, c)
				processed, err := ProcessCmd(parsedCmd, params)
				if err != nil {
					return err
				}
				RunCmd(processed)
				return nil
			},
		})
	}
	return commands
}

// ReadFlags reads the Flag defined in the ParsedCmd
func ReadFlags(parsedCmd ParsedCmd) []cli.Flag {
	flags := make([]cli.Flag, len(parsedCmd.Vars))
	for i, v := range parsedCmd.Vars {
		flags[i] = cli.StringFlag{
			Name:     v.Name,
			Required: v.Value == "",
			Value:    v.Value,
		}
	}
	return flags
}

// ReadParams reads the params definied in the ParsedCmd from the cli context
func ReadParams(parsedCmd ParsedCmd, c *cli.Context) map[string]string {
	params := make(map[string]string)

	for _, v := range parsedCmd.Vars {
		params[v.Name] = c.String(v.Name)
	}
	return params
}
