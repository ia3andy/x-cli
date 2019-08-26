// +build !windows

package cmds

import (
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/creack/pty"
)

// RunCmd the specified command
func RunCmd(cmd string) {
	words := strings.Fields(cmd)
	println("Running command: " + cmd + "\n")
	command := exec.Command(words[0], words[1:]...)

	f, err := pty.Start(command)
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, f)
}
