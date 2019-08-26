package cmds

import (
	"os/exec"
	"strings"
)

// RunCmd the specified command
func RunCmd(cmd string) {
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
