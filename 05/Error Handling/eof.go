package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

// EOF is the error returned by Read when no more input is
// available.
var EOF = errors.New("EOF")

func endOfFile() error {
	in := bufio.NewReader(os.Stdin)
	for {
		r, _, err := in.ReadRune()
		if err == io.EOF {
			break // finished reading
		}
		if err != nil {
			return fmt.Errorf("read failed: %v\n", err)
		}
		fmt.Println(r)
	}
	return nil
}
