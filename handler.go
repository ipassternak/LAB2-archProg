package lab2

import "io"

type ComputeHandler struct {
	r io.Reader
	w io.Writer
}

func (ch *ComputeHandler) Compute() error {
	input := make([]byte, 1024)
	for _, readErr := ch.r.Read(input); readErr != io.EOF; {
		if readErr != nil {
			return readErr
		}
	}
	res, err := PostfixToPrefix(string(input))
	if err != nil {
		return err
	}
	_, writeErr := ch.w.Write([]byte(res))
	return writeErr
}
