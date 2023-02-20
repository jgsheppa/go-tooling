package count

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func (c *counter) Words() int {
	words := 0
	scanner := bufio.NewScanner(c.Input)
	for scanner.Scan() {
		wordsInLine := strings.Split(strings.TrimSpace(scanner.Text()), " ")
		words += len(wordsInLine)
	}
	return words
}

func Words() int {
	c, err := NewCounter(
		FromArgs(os.Args[1:]),
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	return c.Words()
}
