package findgo_test

import (
	"archive/zip"
	"os"
	"testing"
	"testing/fstest"

	"github.com/jgsheppa/go-tooling/cmd/findgo"
)

func TestFilesInZIP(t *testing.T) {
	t.Parallel()
	fsys, err := zip.OpenReader(
		"testdata.zip")
	if err != nil {
		t.Fatal(err)
	}
	want := 4
	got := findgo.Files(fsys)
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

// Tip: When you need to test some function that takes a filesystem, it’s faster to give it a
// MapFS than a real tree of files on disk.
func TestFilesInMemory(t *testing.T) {
	t.Parallel()
	fsys := fstest.MapFS{
		"file.go":                {},
		"subfolder/subfolder.go": {},
		"subfolder2/another.go":  {},
		"subfolder2/file.go":     {},
	}
	want := 4
	got := findgo.Files(fsys)
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func BenchmarkFilesOnDisk(b *testing.B) {
	fsys := os.DirFS("testdata/findgo")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findgo.Files(fsys)
	}
}

func BenchmarkFilesInMemory(b *testing.B) {
	fsys := fstest.MapFS{
		"file.go":                {},
		"subfolder/subfolder.go": {},
		"subfolder2/another.go":  {},
		"subfolder2/file.go":     {},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findgo.Files(fsys)
	}
}
