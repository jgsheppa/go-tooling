package hello

import (
	"bytes"
	"testing"
)

func TestPrintsHelloMessageToWriter(t *testing.T) {
	t.Parallel()
	fakeTerminal := &bytes.Buffer{}
	p := &Printer{
		Output: fakeTerminal,
	}
	p.Print()
	want := "Hello, world\n"
	got := fakeTerminal.String()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}
