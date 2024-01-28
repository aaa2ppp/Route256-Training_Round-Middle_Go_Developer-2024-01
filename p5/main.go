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

func sign(a int) int {
	if a < 0 {
		return -1
	}
	if a > 0 {
		return 1
	}
	return 0
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func task(in *bufio.Reader, out *bufio.Writer) error {
	var n int
	if _, err := fmt.Fscan(in, &n); err != nil {
		return err
	}

	// read sequence
	seq := make([]int, 0, n)
	for i := 0; i < n; i++ {
		var v int
		if _, err := fmt.Fscan(in, &v); err != nil {
			return err
		}
		seq = append(seq, v)
	}

	b := seq[0]
	c := 0

	// "сжимаем" жадно
	var res []int
	for _, v := range seq[1:] {
		if d := v - (b + c); abs(d) == 1 && (c == 0 || sign(c) == d) {
			c += d
			continue
		}
		res = append(res, b, c)
		b = v
		c = 0
	}
	res = append(res, b, c)

	// TODO: надо проверить, как тесты относятся к финальным пробелам
	fmt.Fprintf(out, "%d\n%d", len(res), res[0])
	for _, v := range res[1:] {
		fmt.Fprintf(out, " %d", v)
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
