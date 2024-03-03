package main

import (
	"flag"
	"io"
	"os"
	"strings"

	lab2 "github.com/ipassternak/LAB2-archProg"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	inputFile       = flag.String("f", "", "File to read")
	outputFile      = flag.String("o", "", "File to output")
)

func main() {
	flag.Parse()
	if len(*inputExpression) < 1 && len(*inputFile) < 1 {
		os.Stderr.WriteString("You should provide atleast one flag as arg")
		os.Exit(1)
	}
	var readFile io.Reader
	if len(*inputExpression) > 0 {
		readFile = strings.NewReader(*inputExpression)
	} else {
		readFile, err := os.Open("file.go")
		if err == nil {
			os.Stderr.WriteString("Incorrect file to read")
			os.Exit(1)
		}
	}
	handler := &lab2.ComputeHandler{}
	err := handler.Compute()
	if err == nil {
		os.Stderr.WriteString("LOL it is test err, u r loooser")
		os.Exit(1)
	}
}
