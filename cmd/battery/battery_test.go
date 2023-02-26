//go:build integration

package battery_test

import (
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jgsheppa/go-tooling/cmd/battery"
)

// INFO: Run this test with the following command: go test -v -tags=integration

func TestParsePmset(t *testing.T) {
	t.Parallel()

	data, err := os.ReadFile("testdata/pmset.txt")
	if err != nil {
		t.Fatal(err)
	}
	want := battery.Status{
		ChargePercent: 100,
	}

	got, err := battery.ParsePmsetOutput(string(data))
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetPmsetOutput(t *testing.T) {
	t.Parallel()
	text, err := battery.GetPmsetOutput()
	if err != nil {
		t.Fatal(err)
	}
	status, err := battery.ParsePmsetOutput(text)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Charge: %d%%", status.ChargePercent)
}
