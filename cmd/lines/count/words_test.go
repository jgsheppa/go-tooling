package count

import (
	"bytes"
	"testing"
)

func TestLines(t *testing.T) {
	t.Parallel()
	inputBuf := bytes.NewBufferString("one\ntwo three \n four five six")
	c, err := NewCounter(
		WithInput(inputBuf),
	)
	if err != nil {
		t.Fatal(err)
	}
	want := 6
	got := c.Words()
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}
