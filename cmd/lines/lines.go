package lines

import (
	"bufio"
	"errors"
	"io"
	"os"
)

type counter struct {
	Input  io.Reader
	Output io.Writer
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

func WithOutput(input io.Reader) option {
	return func(c *counter) error {
		if input == nil {
			return errors.New("nil output reader")
		}
		c.Input = input
		return nil
	}
}

func NewFileReader() *counter {
	return &counter{
		Input: os.Stdin,
	}
}

func (f *counter) Lines() int {
	lines := 0
	scanner := bufio.NewScanner(f.Input)
	for scanner.Scan() {
		lines++
	}
	return lines
}

func Lines() int {
	c, err := NewCounter()
	if err != nil {
		panic("Lines function: internal error")
	}
	return c.Lines()
}
