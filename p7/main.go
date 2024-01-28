package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

// func solution(...) (...) {
// 	// TODO
// }

type tSegment struct {
	begin, end int
}

func parseSegment(s string) (tSegment, error) {
	var seg tSegment

	items := strings.Split(s, "-")
	var (
		v1, v2 int
		err    error
	)

	switch len(items) {
	case 1:
		if v1, err = strconv.Atoi(items[0]); err != nil {
			return seg, fmt.Errorf("%s: %w", s, err)
		}
		v2 = v1
	case 2:
		if v1, err = strconv.Atoi(items[0]); err != nil {
			return seg, fmt.Errorf("%s: %w", s, err)
		}
		if v2, err = strconv.Atoi(items[1]); err != nil {
			return seg, fmt.Errorf("%s: %w", s, err)
		}
	default:
		return seg, fmt.Errorf("%s: bad segment", s)
	}

	if v1 < 1 || v2 < 1 || v2 < v1 {
		return seg, fmt.Errorf("%s: bad segment", s)
	}

	seg = tSegment{v1, v2}

	return seg, nil
}

func (s tSegment) String() string {
	if s.begin == s.end {
		return strconv.Itoa(s.begin)
	}
	return fmt.Sprintf("%d-%d", s.begin, s.end)
}

func task(in *bufio.Reader, out *bufio.Writer) error {
	var k int
	if _, err := fmt.Fscan(in, &k); err != nil {
		return err
	}
	// skip eol
	in.ReadString('\n')

	line, err := in.ReadString('\n')
	if err != nil && err != io.EOF {
		return err
	}
	line = strings.TrimSpace(line)
	if debugEnable {
		log.Println("line:", line)
	}

	// отметим каждую распечатанную страницу (для максимум 100 страниц это проще,
	// чем создавать список отрезков, а потом его инвертировать)

	pages := make([]bool, k+1)
	for _, s := range strings.Split(line, ",") {
		seg, err := parseSegment(s)
		if err != nil {
			return err
		}
		if debugEnable {
			log.Println("seg:", seg)
		}
		for i := seg.begin; i <= seg.end; i++ {
			pages[i] = true
		}
	}

	var list []string
	for i := 1; i < len(pages); {
		for i < len(pages) && pages[i] {
			i++
		}
		if i == len(pages) {
			break
		}
		seg := tSegment{begin: i}
		for i < len(pages) && !pages[i] {
			seg.end++
			i++
		}
		seg.end = i - 1
		list = append(list, seg.String())
	}

	res := strings.Join(list, ",")
	if debugEnable {
		log.Println("res:", res)
	}

	fmt.Fprintln(out, res)

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
