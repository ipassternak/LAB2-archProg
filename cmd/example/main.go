package main

import (
	"flag"
	"os"

	lab2 "github.com/ipassternak/LAB2-archProg"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	inputFile       = flag.String("f", "", "File to read")
	outputFile      = flag.String("o", "", "File to output")
)

func main() {
	flag.Parse()
	if len(*inputExpression) < 1 && len(*inputFile) < 1 && len(*outputFile) < 1 {
		os.Stderr.WriteString("You should provide atleast one flag as arg")
		return
	}
	handler := &lab2.ComputeHandler{
		// <paste your streams here>
	}
	err := handler.Compute()
	if err != nil {
		os.Stderr.WriteString("LOL it is test err, u r loooser")
	}
}
