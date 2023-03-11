package store_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jgsheppa/go-tooling/cmd/store"
)

func TestStoreFile(t *testing.T) {
	t.Parallel()
	path := t.TempDir() + "/store.bin"
	output := store.Open(path)
	want := []int{2, 3, 5, 7, 11}
	err := output.Save(want)

	if err != nil {
		t.Fatal(err)
	}
	output.Close()
	input := store.Open(path)
	var got []int
	err = input.Load(&got)
	if err != nil {
		t.Fatal(err)
	}
	input.Close()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
