package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
)

type tSearcher struct {
	res []int16
}

func newBorderSeacher() *tSearcher {
	return &tSearcher{}
}

func (s *tSearcher) search(matrix [][]int16, x1, y1, x2, y2 int16, deep int16) {

	for x, end := x1, x2-3; x <= end; x++ {
		for y, end := y1, y2-3; y <= end; {

			switch v := matrix[x][y]; {
			case v == 0:
				y++

			case v > 0:
				y = v

			case v == -1:
				// bingo
				s.res = append(s.res, deep)

				yy := y + 3
				for yy < y2 && matrix[x][yy] == -1 {
					yy++
				}

				// matrix[x][y] = yy
				xx := x + 1
				for xx < x2 && matrix[xx][y] == -1 {
					matrix[xx][y] = yy
					xx++
				}

				s.search(matrix, x+1, y+1, xx-1, yy-1, deep+1)
				y = yy
			}
		}
	}
}

func solution(matrix [][]int16) []int16 {
	n := len(matrix)
	m := len(matrix[0])

	bs := newBorderSeacher()
	bs.search(matrix, 0, 0, int16(n), int16(m), 0)
	res := bs.res

	sort.Slice(res, func(i, j int) bool {
		return res[i] < res[j]
	})

	return res
}

func task(in *bufio.Reader, out *bufio.Writer) error {
	var n, m int
	if _, err := fmt.Fscan(in, &n, &m); err != nil {
		return err
	}

	// read matrix
	matrix := makeMatrix(n, m)
	for i := 0; i < n; i++ {
		var line string
		if _, err := fmt.Fscan(in, &line); err != nil {
			return err
		}
		for j, c := range []byte(line) {
			if c == '*' {
				matrix[i][j] = -1
			}
		}
	}

	res := solution(matrix)

	for _, v := range res {
		fmt.Fprintf(out, "%d ", v)
	}
	fmt.Fprintln(out)

	return nil
}

func makeMatrix(n, m int) [][]int16 {
	buf := make([]int16, n*m)
	matrix := make([][]int16, n)
	for i, j := 0, 0; i < n; i, j = i+1, j+m {
		matrix[i] = buf[j : j+m]
	}
	return matrix
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

	for i := 0; i < t; i++ {
		if debugEnable {
			log.Printf("--- task %d", t)
		}
		if err := task(bIn, bOut); err != nil {
			return fmt.Errorf("task%d: %w", i+1, err)
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
