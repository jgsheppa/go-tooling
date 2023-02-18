package counter

import (
	"fmt"
	"io"
	"os"
)

type Counter struct {
	Current int
	Output  io.Writer
}

func NewCounter() *Counter {
	return &Counter{Current: 0,
		Output: os.Stdout,
	}
}

func (c *Counter) Next() {
	fmt.Fprintln(c.Output, c.Current+1)
}

func Next() {
	NewCounter().Next()
}