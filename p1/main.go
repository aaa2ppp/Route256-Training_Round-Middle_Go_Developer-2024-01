package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func solution(nums []int) bool {
	var cnts [5]int
	for _, k := range nums {
		cnts[k]++
	}
	return cnts == [5]int{0, 4, 3, 2, 1}
}

func task(in io.Reader, out io.Writer) error {
	nums := make([]int, 10)
	for i := range nums {
		if _, err := fmt.Fscan(in, &nums[i]); err != nil {
			return err
		}
	}
	if solution(nums) {
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
