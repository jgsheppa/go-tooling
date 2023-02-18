package counter

import (
	"bytes"
	"testing"
)

func TestNextCounterToWriter(t *testing.T) {
	t.Parallel()
	fakeTerminal := &bytes.Buffer{}

	c := &Counter{
		Current: 0,
		Output:  fakeTerminal,
	}
	c.Next()
	want := "1\n"
	got := fakeTerminal.String()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}
