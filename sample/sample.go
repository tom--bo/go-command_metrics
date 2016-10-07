package main

import (
	"fmt"
	"github.com/tom--bo/go-command_metrics"
	"os/exec"
)

func main() {
	command_metrics.Print(1)

	_cmd := exec.Command("sleep", "1")
	cmd := command_metrics.WrapCommand("test", _cmd)
	err := cmd.Run()

	// fmt.Println("1")
	if err != nil {
		fmt.Println("err: %v", err)
	}
	cmd2 := command_metrics.WrapCommand("test2", _cmd2)
	err = cmd2.Run()
	//fmt.Println("2")

	cmd3 := command_metrics.WrapCommand("test", _cmd3)
	err = cmd3.Run()
}
