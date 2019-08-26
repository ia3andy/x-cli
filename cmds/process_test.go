package cmds_test

import (
	"testing"

	"github.com/ia3andy/xd/cmds"
	"gotest.tools/assert"
)

// should process a command when it contains one paramameter
func TestProcessOneParamCmd(t *testing.T) {
	cmd := "echo \"Hello ${target}\""
	parsedCmd := cmds.ParseCmd(cmd)
	processed, err := cmds.ProcessCmd(parsedCmd, map[string]string{"target": "Bobby"})
	assert.NilError(t, err)
	assert.Equal(t, processed, `echo "Hello Bobby"`)
}

// should return an error when a param is missing
func TestProcessMissingParamCmd(t *testing.T) {
	cmd := "echo \"Hello ${target}\""
	parsedCmd := cmds.ParseCmd(cmd)
	processed, err := cmds.ProcessCmd(parsedCmd, map[string]string{})
	assert.Error(t, err, "param 'target' must be defined")
	assert.Equal(t, processed, "")
}

// should process a command when it contains two paramameters (using space and double quote)
func TestProcessTwoParamsCmd(t *testing.T) {
	cmd := `mv ${from} ${to}`
	parsedCmd := cmds.ParseCmd(cmd)
	processed, err := cmds.ProcessCmd(parsedCmd, map[string]string{"from": "./rep/file.zip", "to": `"./anotherrep/some name.zip"`})
	assert.NilError(t, err)
	assert.Equal(t, processed, `mv ./rep/file.zip "./anotherrep/some name.zip"`)
}
