package shell_test

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jgsheppa/go-tooling/cmd/shell"
)

func TestCmdFromString(t *testing.T) {
	t.Parallel()
	input := "/bin/ls -l main.go"
	want := []string{"/bin/ls", "-l", "main.go"}
	cmd, err := shell.CmdFromString(input)
	if err != nil {
		t.Fatal(err)
	}

	got := cmd.Args
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestCmdFromStringErrorsOnEmptyInput(t *testing.T) {
	t.Parallel()
	_, err := shell.CmdFromString("")
	if err == nil {
		t.Fatal("want error on empty input, got nil")
	}
}

func TestNewSession(t *testing.T) {
	t.Parallel()
	stdin := os.Stdin
	stdout := os.Stdout
	stderr := os.Stderr

	want := shell.Session{
		Stdin:  stdin,
		Stdout: stdout,
		Stderr: stderr,
	}
	got := *shell.NewSession(stdin, stdout, stderr)
	if want != got {
		t.Errorf("want %#v, got %#v", want, got)
	}
}

func TestRun(t *testing.T) {
	t.Parallel()
	stdin := strings.NewReader("echo hello\n\n")
	stdout := &bytes.Buffer{}
	session := shell.NewSession(stdin, stdout, io.Discard)
	session.DryRun = true
	session.Run()
	want := "> echo hello\n> > \nBe seeing you!\n"
	got := stdout.String()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
