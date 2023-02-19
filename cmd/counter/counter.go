package counter

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Counter struct {
	Current int
	Output  io.Writer
}

func NewCounter(start int) *Counter {
	return &Counter{Current: start,
		Output: os.Stdout,
	}
}

func (c *Counter) Next() {
	c.Current = c.Current + 1
	fmt.Printf("current iteration: %v \n", c.Current)
	fmt.Fprintln(c.Output, c.Current)
}

func Next() {
	fmt.Print("Enter starting number: ")
	reader := bufio.NewReader(os.Stdin)
	// ReadString will block until the delimiter is entered
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return
	}

	// remove the delimeter from the string
	input = strings.TrimSuffix(input, "\n")
	start, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Could not covert string to int", err)
		return
	}

	NewCounter(start).Next()
}
