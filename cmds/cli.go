package cmds

import "github.com/urfave/cli"

// ReadFlags reads the Flag defined in the ParsedCmd
func ReadFlags(parsedCmd ParsedCmd) []cli.Flag {
	flags := make([]cli.Flag, len(parsedCmd.Vars))
	for i, v := range parsedCmd.Vars {
		flags[i] = cli.StringFlag{
			Name:     v.Name,
			Required: true,
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
