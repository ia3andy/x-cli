package cmds_test

import (
	"testing"

	"github.com/ia3andy/xd/cmds"
	"gotest.tools/assert"
)

// should parse a command when it contains one paramameter
func TestParseOneParamCmd(t *testing.T) {
	cmd := "echo \"Hello ${target}\""
	result := cmds.ParseCmd(cmd)

	assert.Equal(t, result.Text, cmd)
	assert.Equal(t, len(result.Vars), 1)
	assert.Equal(t, result.Vars[0], cmds.Variable{"${target}", "target"})
}

// should parse a command when it doesn't contain any paramameter
func TestParseNoParamCmd(t *testing.T) {
	cmd := "echo \"Hello\""
	result := cmds.ParseCmd(cmd)

	assert.Equal(t, result.Text, cmd)
	assert.Equal(t, len(result.Vars), 0)
}

// should parse a command when it contains two paramameters
func TestParseTwoParamCmd(t *testing.T) {
	cmd := "mv ${src} ${dst}"
	result := cmds.ParseCmd(cmd)

	assert.Equal(t, result.Text, cmd)
	assert.Equal(t, len(result.Vars), 2)
	assert.Equal(t, result.Vars[0], cmds.Variable{"${src}", "src"})
	assert.Equal(t, result.Vars[1], cmds.Variable{"${dst}", "dst"})
}
