package main

import "github.com/jgsheppa/go-tooling/cmd/pipeline"

func main() {
	pipeline.FromString("hello, world\n").Stdout()
}
