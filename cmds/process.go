package cmds

import (
	"fmt"
	"strings"
)

// ProcessCmd is processing the command to make it executable
func ProcessCmd(parsedCmd ParsedCmd, params map[string]string) (string, error) {
	processed := parsedCmd.Text
	for _, v := range parsedCmd.Vars {
		if params[v.Name] == "" {
			return "", fmt.Errorf("param '%s' must be defined", v.Name)
		}
		processed = strings.ReplaceAll(processed, v.Expression, params[v.Name])
	}
	return processed, nil
}
