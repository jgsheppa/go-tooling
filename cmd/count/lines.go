package count

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

type counter struct {
	Input     io.Reader
	Output    io.Writer
	Args      []string
	wordCount bool
	verbose   bool
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

func FromArgs(args []string) option {
	return func(c *counter) error {
		fset := flag.NewFlagSet(os.Args[0],
			flag.ContinueOnError)
		wordCount := fset.Bool("w", false,
			"Count words instead of lines")
		verbose := fset.Bool("v", false,
			"Gives context to output")
		fset.SetOutput(c.Output)
		err := fset.Parse(args)
		if err != nil {
			return err
		}
		c.wordCount = *wordCount
		c.verbose = *verbose

		args = fset.Args()
		if len(args) < 1 {
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
		FromArgs(os.Args[1:]),
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	return c.Lines()
}

func RunCLI() {
	c, err := NewCounter(
		FromArgs(os.Args[1:]),
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if c.wordCount {
		fmt.Println(c.Words())
	} else if c.verbose {
		fmt.Printf("%v words\n", c.Words())
	} else {
		fmt.Println(c.Lines())
	}
}
