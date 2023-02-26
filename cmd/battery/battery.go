package battery

import (
	"fmt"
	"os/exec"
	"regexp"
	"runtime"
	"strconv"
)

type Status struct {
	ChargePercent int
}

var pmsetOutput = regexp.MustCompile("([0-9]+)%")

func ParsePmsetOutput(data string) (Status, error) {
	if runtime.GOOS != "darwin" {
		panic("pmset command only functions on macOS")
	}

	matches := pmsetOutput.FindStringSubmatch(data)

	if len(matches) < 2 {
		return Status{}, fmt.Errorf("failed to parse pmset output: %q", data)
	}

	charge, err := strconv.Atoi(matches[1])
	if err != nil {
		return Status{}, fmt.Errorf("failed to parse charge percentage: %q", matches[1])
	}
	return Status{ChargePercent: charge}, nil
}

func GetPmsetOutput() (string, error) {
	if runtime.GOOS != "darwin" {
		panic("pmset command only functions on macOS")
	}

	data, err := exec.Command("/usr/bin/pmset", "-g", "ps").
		CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(data), nil
}
