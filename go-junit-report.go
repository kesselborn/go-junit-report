package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jstemmer/go-junit-report/parser"
)

var (
	noXMLHeader bool
	packageName string
	setExitCode bool
	verbose     bool
)

func init() {
	flag.BoolVar(&noXMLHeader, "no-xml-header", false, "do not print xml header")
	flag.StringVar(&packageName, "package-name", "", "specify a package name (compiled test have no package name in output)")
	flag.BoolVar(&setExitCode, "set-exit-code", false, "set exit code to 1 if tests failed")
	flag.BoolVar(&verbose, "verbose", false, "include output of passing tests in junit report")
}

func main() {
	flag.Parse()

	// Read input
	report, err := parser.Parse(os.Stdin, packageName)
	if err != nil {
		fmt.Printf("Error reading input: %s\n", err)
		os.Exit(1)
	}

	// Write xml
	err = JUnitReportXML(report, noXMLHeader, verbose, os.Stdout)
	if err != nil {
		fmt.Printf("Error writing XML: %s\n", err)
		os.Exit(1)
	}

	if setExitCode && report.Failures() > 0 {
		os.Exit(1)
	}
}
