package lab2

import "io"

type ComputeHandler struct {
	r io.Reader
	w io.Writer
}

func (ch *ComputeHandler) Compute() error {
	var input []byte
	buffer := make([]byte, 1024)
	for {
		n, readErr := ch.r.Read(buffer)
		input = append(input, buffer[:n]...)
		if readErr != nil {
			if readErr == io.EOF {
				break
			} else {
				return readErr
			}
		}
	}
	res, err := PostfixToPrefix(string(input))
	if err != nil {
		return err
	}
	_, writeErr := ch.w.Write([]byte(res))
	return writeErr
}
