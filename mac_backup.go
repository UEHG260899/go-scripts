package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Please provide the path of your brewfile")
		return
	}

	command := exec.Command("brew", "--version")

	err := command.Run()

	if err != nil {
		fmt.Println("Homebrew not installed, installing it now")
		err = installHomebrew()

		if err != nil {
			fmt.Errorf("Homebrew installation failed")
			return
		}

		fmt.Println("Homebrew installed successfully")
	}

	fmt.Println("Proceding with the installation of caks/formulae as defined in the brewfile")

	if err = installBrewSoftware(); err != nil {
		fmt.Println("Something unexpected happened when installing the casks/formulae", err.Error())
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

func installBrewSoftware() error {
	pahtToBrewfile := os.Args[1]

	command := exec.Command("brew", "bundle", "--file", pahtToBrewfile)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	return command.Run()
}
