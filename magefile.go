//go:build mage
// +build mage

package main

import (
	"fmt"
	"strings"

	. "github.com/logrusorgru/aurora"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

const (
	packageName = "raw.tools/over"
)

var Default = All

// Helpers
//====

// split a single string command into a list of string
func split(cmd string) (string, []string) {
	var args = strings.Fields(cmd)
	return args[0], args[1:]
}

// execute a command with output on stdout
func run(cmd string) error {
	var args = strings.Fields(cmd)
	return sh.RunV(args[0], args[1:]...)
}

// execute a command with an OK/KO message
func runOK(cmd string, ok string, ko string) error {
	err := run(cmd)
	if err != nil {
		return exit(ko)
	}
	success(ok)
	return nil
}

// silently get the output of a command
func output(cmd string) (string, error) {
	var exe, args = split(cmd)
	return sh.Output(exe, args...)
}

// display a success message
func success(str string) {
	fmt.Printf("%s %s \n", Green("✔"), str)
}

// display a failure message
func failure(str string) {
	fmt.Printf("%s %s \n", Red("✖"), str)
}

// exit with an error message and exit code -1
func exit(str string) error {
	return mg.Fatalf(-1, "%s %s", Red("✖"), str)
}

// filter a list of strings using a test function
func grep(list []string, test func(string) bool) []string {
	out := make([]string, 0)
	for _, line := range list {
		if test(line) {
			out = append(out, line)
		}
	}
	return out
}

// Tasks
//====

// Clean the workdir
func Clean() error {
	return run("go clean")
}

// Run the test suite
func Test() error {
	return runOK("go test -v -race ./...", "Tests succeed", "Tests failed")
}

// Run tests with coverage
func Cover() error {
	return runOK(
		`go test -race ./... -cover -covermode=atomic -coverprofile=coverage.out`,
		"Tests (with coverage) succeed",
		"Tests (with coverage) failed",
	)
}

// Run tests with coverage and generate an HTML report
func CoverHtml() error {
	mg.Deps(Cover)
	return runOK(
		`go tool cover -html=coverage.out -o coverage.html`,
		`Coverage report generated in coverage.html`,
		`Coverage report generation failed`,
	)
}

// Execute static analysis
func Lint() error {
	return runOK("golangci-lint run", "Code is fine", "There is some lints to fix 👆")
}

// Execute the benchmark suite
func Bench() error {
	return runOK(
		`gotest -v -bench . -cpu 1,2,4 -tags bench -run Benchmark`,
		`Benchmark done`,
		`Benchmark failed`,
	)
}

// Compile code for the current platform
func Build() error {
	return runOK(`goreleaser build --snapshot --rm-dist --single-target`, `Build success`, `Build failed`)
}

// Build wk binary for all supported platforms
func BuildAll() error {
	return runOK(`goreleaser build --snapshot --rm-dist`, `Build success`, `Build failed`)
}

// Perform a release
func Release() error {
	return runOK(`goreleaser --rm-dist`, `over released :rocket:`, `Release failed`)
}

// Lint, Build, Test
func All() error {
	mg.Deps(Lint)
	mg.Deps(Test)
	mg.Deps(BuildAll)
	return nil
}
