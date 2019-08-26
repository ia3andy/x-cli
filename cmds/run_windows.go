// +build windows

package cmds

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"
)

// RunCmd the specified command
func RunCmd(cmd string) {
	words := strings.Fields(cmd)
	println("Running command: " + cmd + "\n")
	command := exec.Command(words[0], words[1:]...)
	c := make(chan struct{})
	go run(command, c)
	c <- struct{}{}
	command.Start()
	<-c
	if err := command.Wait(); err != nil {
		fmt.Println(err)
	}
}

func run(cmd *exec.Cmd, c chan struct{}) {
	defer func() { c <- struct{}{} }()
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	<-c
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		m := scanner.Text()
		fmt.Println(m)
	}
	fmt.Println("EOF")
}
