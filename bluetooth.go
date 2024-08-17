package main

import (
	"fmt"
	"os/exec"
)

// Only for linux
func main() {
	command := exec.Command("rfkill", "block", "bluetooth")
	_, err := command.Output()

	if err != nil {
		fmt.Errorf("Could not execute")
	}
}
