package cmds

import (
	"regexp"
)

// Variable is a part of the command that will be replaced when called
type Variable struct {
	Expression string
	Name       string
	Value      string
}

// ParsedCmd is a command and all its extracted information
type ParsedCmd struct {
	Text string
	Vars []Variable
}

// ParseCmd is parsing the command and extracting revelent information
func ParseCmd(cmd string) ParsedCmd {
	varNamesRegexp := regexp.MustCompile(`\$\{([^}|]*)(\|([^}]*))?\}`)
	matches := varNamesRegexp.FindAllStringSubmatch(cmd, -1)
	if matches == nil {
		return ParsedCmd{cmd, make([]Variable, 0)}
	}
	count := len(matches)
	var vars = make([]Variable, count)
	for i, match := range matches {
		vars[i] = Variable{Expression: match[0], Name: match[1], Value: match[3]}
	}
	return ParsedCmd{cmd, vars}
}
