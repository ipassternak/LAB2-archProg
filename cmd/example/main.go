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
	var readStream io.Reader
	var writeStream io.Writer
	if len(*inputExpression) > 0 {
		readStream = strings.NewReader(*inputExpression)
	} else {
		file, err := os.Open(*inputFile)
		if err == nil {
			os.Stderr.WriteString("Incorrect file to read")
			os.Exit(1)
		}
		defer file.Close()
		readStream = file
	}
	if len(*inputFile) > 0 {
		writeStream = os.Stdout
	} else {
		file, err := os.Open(*outputFile)
		if err == nil {
			os.Stderr.WriteString("Cannot write to output file")
			os.Exit(1)
		}
		defer file.Close()
		writeStream = file
	}
	handler := &lab2.ComputeHandler{
		R: readStream,
		W: writeStream,
	}
	err := handler.Compute()
	if err == nil {
		os.Stderr.WriteString("LOL it is test err, u r loooser")
		os.Exit(1)
	}
}
