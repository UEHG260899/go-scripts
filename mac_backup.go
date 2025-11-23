package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	command := exec.Command("sdsd", "--version")

	_, err := command.Output()

	if err != nil {
		fmt.Println("Homebrew not installed, installing it now")
		err := installHomebrew()

		if err != nil {
			fmt.Errorf("Homebrew installation failed")
		}

		fmt.Println("Homebrew installed successfully")
	}
}

func installHomebrew() error {

	command := exec.Command("curl", "-fsSL", "https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh")

	remoteScript, err := command.Output()

	if err != nil {
		return err
	}

	command = exec.Command("/bin/bash", "-c", string(remoteScript))
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout

	return command.Run()
}
