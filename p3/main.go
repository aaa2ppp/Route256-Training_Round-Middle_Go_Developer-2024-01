package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func isLetter(c byte) bool {
	return 'A' <= c && c <= 'Z' || 'a' <= c && c <= 'z'
}

func isDigit(c byte) bool {
	return '0' <= c && c <= '9'
}

func solution(line string) (string, bool) {
	var sb strings.Builder
	sb.Grow(len(line) + (len(line)+3)/4)
	var i int
	for i < len(line) {
		if !(i < len(line) && isLetter(line[i])) {
			return "", false
		}
		sb.WriteByte(line[i])
		i++

		if !(i < len(line) && isDigit(line[i])) {
			return "", false
		}
		sb.WriteByte(line[i])
		i++

		if i < len(line) && isDigit(line[i]) {
			sb.WriteByte(line[i])
			i++
		}

		if !(i < len(line) && isLetter(line[i])) {
			return "", false
		}
		sb.WriteByte(line[i])
		i++

		if !(i < len(line) && isLetter(line[i])) {
			return "", false
		}
		sb.WriteByte(line[i])
		i++

		sb.WriteByte(' ')
	}

	res := sb.String()
	if len(res) > 0 {
		res = res[:len(res)-1]
	}
	return res, true
}

func run(r io.Reader, w io.Writer) error {
	in := bufio.NewReader(r)
	out := bufio.NewWriter(w)
	defer out.Flush()

	var t int
	if _, err := fmt.Fscan(in, &t); err != nil {
		return err
	}
	for i := 0; i < t; i++ {
		var line string
		if _, err := fmt.Fscan(in, &line); err != nil {
			return err
		}
		res, ok := solution(line)
		if ok {
			fmt.Fprintln(out, res)
		} else {
			fmt.Fprintln(out, "-")
		}
	}
	return nil
}

func main() {
	if err := run(os.Stdin, os.Stdout); err != nil {
		log.Fatal(err)
	}
}
