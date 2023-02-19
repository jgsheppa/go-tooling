package count

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

type counter struct {
	Input  io.Reader
	Output io.Writer
	Args   []string
}

type option func(*counter) error

func NewCounter(opts ...option) (counter, error) {
	c := counter{
		Input:  os.Stdin,
		Output: os.Stdout,
	}

	for _, opt := range opts {
		err := opt(&c)
		if err != nil {
			return counter{}, err
		}
	}
	return c, nil
}

func WithInput(input io.Reader) option {
	return func(c *counter) error {
		if input == nil {
			return errors.New("nil input reader")
		}
		c.Input = input
		return nil
	}
}

func WithOutput(output io.Writer) option {
	return func(c *counter) error {
		if output == nil {
			return errors.New("nil output writer")
		}
		c.Output = output
		return nil
	}
}

func WithInputFromArgs(args []string) option {
	return func(c *counter) error {
		if len(args) == 0 {
			return nil
		}
		f, err := os.Open(args[0])
		if err != nil {
			return err
		}
		c.Input = f
		return nil
	}
}

func NewFileReader() *counter {
	return &counter{
		Input: os.Stdin,
	}
}

func (c *counter) Lines() int {
	lines := 0
	scanner := bufio.NewScanner(c.Input)
	for scanner.Scan() {
		lines++
	}
	return lines
}

func Lines() int {
	c, err := NewCounter(
		WithInputFromArgs(os.Args[1:]),
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	return c.Lines()
}
