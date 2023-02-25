package main

import (
	"os"
	"os/exec"
)

func main() {
	// count.RunCLI()

	c := exec.Command("/usr/bin/ls")
	c.Stdout = os.Stdout
	c.Run()
}
