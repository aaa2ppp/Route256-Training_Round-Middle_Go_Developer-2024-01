package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func solution(d, m, y int) bool {
	// мне лень
	dt := time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.UTC)
	return dt.Year() == y && dt.Month() == time.Month(m) && dt.Day() == d
}

func task(in io.Reader, out io.Writer) error {
	var d, m, y int
	if _, err := fmt.Fscan(in, &d, &m, &y); err != nil {
		return err
	}

	if solution(d, m, y) {
		fmt.Fprintln(out, "YES")
	} else {
		fmt.Fprintln(out, "NO")
	}
	return nil
}

func run(in io.Reader, out io.Writer) error {
	var t int
	if _, err := fmt.Fscan(in, &t); err != nil {
		return err
	}
	for i := 0; i < t; i++ {
		if err := task(in, out); err != nil {
			return fmt.Errorf("task%d: %w", i+1, err)
		}
	}
	return nil
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	if err := run(in, out); err != nil {
		log.Fatal(err)
	}
}
