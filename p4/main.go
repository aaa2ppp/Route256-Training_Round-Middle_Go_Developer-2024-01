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
	var n int
	if _, err := fmt.Fscan(in, &n); err != nil {
		return err
	}

	minimum := 15
	maximum := 30
	current := 22

	for i := 0; i < n; i++ {
		var (
			con string
			val int
		)
		fmt.Fscan(in, &con, &val)

		switch con {
		case ">=":
			if minimum < val {
				minimum = val
			}
		case "<=":
			if maximum > val {
				maximum = val
			}
		default:
			return fmt.Errorf("%s: unknown condition", con)
		}

		if maximum < minimum {
			current = -1
		} else if current < minimum {
			current = minimum
		} else if current > maximum {
			current = maximum
		}

		fmt.Fprintln(out, current)
	}

	fmt.Fprintln(out)
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
