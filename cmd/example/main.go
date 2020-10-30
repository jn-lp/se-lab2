package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	lab2 "github.com/jn-lp/se-lab2"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	inputFile       = flag.String("f", "", "File with expression to compute")
	outputFile      = flag.String("o", "", "File to write result to")
)

func main() {
	flag.Parse()

	var err error
	var output io.Writer

	if *outputFile != "" {
		output, err = os.Create(*outputFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Unable to create file:", err)
			return
		}
	} else {
		output = os.Stdout
	}

	var input io.Reader

	if *inputFile != "" {
		input, err = os.Open(*inputFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Unable to open file:", err)
			return
		}
	} else {
		input = strings.NewReader(*inputExpression)
	}

	handler := &lab2.ComputeHandler{
		Input:  input,
		Output: output,
	}

	err = handler.Compute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
}
