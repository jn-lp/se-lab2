package lab2

import (
	"bufio"
	"io"
)

// ComputeHandler should be constructed with input io.Reader and output io.Writer.
type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (ch *ComputeHandler) Compute() (err error) {
	scanner := bufio.NewScanner(ch.Input)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	for _, eachline := range txtlines {
		out, err := PostfixToInfix(eachline)
		if err != nil {
			return err
		}
		ch.Output.Write([]byte(out + "\n"))
	}

	return nil
}
