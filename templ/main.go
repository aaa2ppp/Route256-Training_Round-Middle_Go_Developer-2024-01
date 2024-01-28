package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

// func solution(...) (...) {
// 	// TODO
// }

func task(in *bufio.Reader, out *bufio.Writer) error {
	// TODO
	return nil
}

func run(in io.Reader, out io.Writer) (err error) {
	bIn := bufio.NewReader(in)
	bOut := bufio.NewWriter(out)
	defer func() {
		if e := bOut.Flush(); e != nil && err == nil {
			err = e
		}
	}()

	var t int
	if _, err := fmt.Fscan(bIn, &t); err != nil {
		return err
	}
	// skip eol
	bIn.ReadBytes('\n')

	for i := 1; i <= t; i++ {
		if debugEnable {
			log.Printf("--- task %d", i)
		}
		if err := task(bIn, bOut); err != nil {
			return fmt.Errorf("task%d: %w", i, err)
		}
	}

	return nil
}

var debugEnable bool

func main() {
	_, debugEnable = os.LookupEnv("DEBUG")

	if err := run(os.Stdin, os.Stdout); err != nil {
		log.Fatal(err)
	}
}
