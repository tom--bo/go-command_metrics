package main

import (
	"fmt"
	"github.com/tom--bo/go-command_metrics"
	"os/exec"
)

func main() {
	_cmd := exec.Command("sleep", "1")
	cmd := command_metrics.WrapCommand("test", _cmd)
	err := cmd.Run()
	if err != nil {
		fmt.Println("err: %v", err)
	}

	_cmd2 := exec.Command("sleep", "1")
	cmd2 := command_metrics.WrapCommand("test2", _cmd2)
	err = cmd2.Run()
	if err != nil {
		fmt.Println("err: %v", err)
	}

	_cmd3 := exec.Command("sleep", "1")
	cmd3 := command_metrics.WrapCommand("test", _cmd3)
	err = cmd3.Run()
	if err != nil {
		fmt.Println("err: %v", err)
	}

	_cmd4 := exec.Command("sleep", "9")
	cmd4 := command_metrics.WrapCommand("test", _cmd4)
	err = cmd4.Run()
	if err != nil {
		fmt.Println("err: %v", err)
	}
}
