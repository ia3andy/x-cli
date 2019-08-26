package cmds

import (
	"fmt"
	"strings"
)

// ProcessCmd is processing the command to make it executable
func ProcessCmd(parsedCmd ParsedCmd, params map[string]string) (string, error) {
	processed := parsedCmd.Text
	for _, v := range parsedCmd.Vars {
		param := params[v.Name]
		if param == "" {
			param = v.Value
		}
		if param == "" {
			return "", fmt.Errorf("param '%s' must be defined", v.Name)
		}
		processed = strings.ReplaceAll(processed, v.Expression, param)
	}
	return processed, nil
}
