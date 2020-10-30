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

	for scanner.Scan() {
		out, err := PostfixToInfix(scanner.Text())
		if err != nil {
			return err
		}
		ch.Output.Write([]byte(out + "\n"))
	}

	return nil
}
