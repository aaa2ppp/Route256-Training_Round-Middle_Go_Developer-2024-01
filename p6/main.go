package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

// func solution(...) (...) {
// 	// TODO
// }

type tTerminal struct {
	text [][]byte
	row  int
	pos  int
}

func newTerminal() *tTerminal {
	return &tTerminal{
		text: [][]byte{nil},
	}
}

func (t *tTerminal) writeText(out io.Writer, lineNums bool) {
	for i := range t.text {
		if lineNums {
			fmt.Fprintf(out, "%03d: ", i+1)
		}
		out.Write(t.text[i])
		out.Write([]byte{'\n'})
	}
	out.Write([]byte("-\n"))
}

func (t *tTerminal) input(c byte) {
	// строчные буквы латинского алфавита, цифры и буквы L, R, U, D, B, E, N.
	switch c {
	case 'L':
		t.pos--
		t.fixPos()
	case 'R':
		t.pos++
		t.fixPos()
	case 'U':
		t.row--
		t.fixPos()
	case 'D':
		t.row++
		t.fixPos()
	case 'B':
		t.pos = 0
	case 'E':
		t.pos = len(t.text[t.row])
	case 'N':
		t.enter()
	default:
		t.insertChar(c)
	}
}

func (t *tTerminal) fixPos() {
	if t.row < 0 {
		t.row = 0
	}
	if n := len(t.text) - 1; t.row > n {
		t.row = n
	}
	if t.pos < 0 {
		t.pos = 0
	}
	if n := len(t.text[t.row]); t.pos > n {
		t.pos = n
	}
}

func (t *tTerminal) insertChar(c byte) {
	old := t.text[t.row]
	buf := append(old, 0)
	copy(buf[t.pos+1:], old[t.pos:])
	buf[t.pos] = c
	t.text[t.row] = buf
	t.pos++
}

func (t *tTerminal) addNewLine() {
	old := t.text
	buf := append(old, nil)
	copy(buf[t.row+1:], old[t.row:])
	t.text = buf
	t.row++
	t.pos = 0
}

func (t *tTerminal) enter() {
	cur := t.text[t.row]
	buf := make([]byte, len(cur[t.pos:]))
	copy(buf, cur[t.pos:])
	t.text[t.row] = cur[:t.pos]
	t.addNewLine()
	t.text[t.row] = buf
}

func task(in *bufio.Reader, out *bufio.Writer) error {
	term := newTerminal()

	input, err := in.ReadBytes('\n')
	if err != nil && err != io.EOF {
		return err
	}
	input = bytes.TrimSpace(input)
	if debugEnable {
		log.Printf("input: %s", input)
	}

	for _, c := range input {
		term.input(c)
		if debugEnable {
			term.writeText(os.Stderr, true)
		}
	}

	term.writeText(out, false)

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
