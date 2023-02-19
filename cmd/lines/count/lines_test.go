package count

import (
	"bytes"
	"testing"
)

func TestLineCounter(t *testing.T) {
	t.Parallel()

	inputBuf := bytes.NewBufferString("1\n2\n3\n")
	c, err := NewCounter(
		WithInput(inputBuf),
	)
	if err != nil {
		t.Fatal(err)
	}
	want := 3
	got := c.Lines()
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestInputFromArgs(t *testing.T) {
	t.Parallel()
	args := []string{"testdata/three_lines.txt"}

	c, err := NewCounter(
		WithInputFromArgs(args),
	)

	if err != nil {
		t.Fatal(err)
	}
	want := 3
	got := c.Lines()
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestWithInputFromArgsEmpty(t *testing.T) {
	t.Parallel()
	inputBuf := bytes.NewBufferString("1\n2\n3")
	c, err := NewCounter(
		WithInput(inputBuf),
		WithInputFromArgs([]string{}),
	)
	if err != nil {
		t.Fatal(err)
	}
	want := 3
	got := c.Lines()
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}
