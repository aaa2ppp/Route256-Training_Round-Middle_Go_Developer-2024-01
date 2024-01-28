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

var (
	cardVals  = []byte("23456789TJQKA")
	cardSuits = []byte("SCDH")
)

type tCard uint64

func parseCard(s string) (tCard, error) {
	val := s[0]
	suit := s[1]

	var card tCard
	var bingo bool

	bingo = false
	for i, v := range []byte(cardVals) {
		if v == val {
			card |= tCard(i) << 2
			bingo = true
			break
		}
	}
	if !bingo {
		return 0, fmt.Errorf("%s: unknown value value", s)
	}

	bingo = false
	for i, v := range []byte(cardSuits) {
		if v == suit {
			card |= tCard(i)
			bingo = true
			break
		}
	}
	if !bingo {
		return 0, fmt.Errorf("%s: unknown card suit", s)
	}

	return card, nil
}

func (c tCard) String() string {
	suit := cardSuits[c&0b11]
	val := cardVals[c>>2]
	return string([]byte{val, suit})
}

func handPower(c1, c2, c3 tCard) int {
	cv1 := int(c1 >> 2)
	cv2 := int(c2 >> 2)
	cv3 := int(c3 >> 2)

	v := cv1
	if v < cv2 {
		v = cv2
	}
	if v < cv3 {
		v = cv3
	}

	if cv1 == cv2 && cv1 == cv3 {
		return 3*16 + cv1
	}

	if cv1 == cv2 || cv1 == cv3 {
		v = 2*16 + cv1
	}

	if cv2 == cv3 {
		v2 := 2*16 + cv2
		if v < v2 {
			v = v2
		}
	}

	return v
}

func task(in *bufio.Reader, out *bufio.Writer) error {
	var n int
	if _, err := fmt.Fscan(in, &n); err != nil {
		return err
	}

	var desk uint64
	count := 52

	// read hands
	hands := make([][2]tCard, n)
	for i := range hands {
		var s1, s2 string
		if _, err := fmt.Fscan(in, &s1, &s2); err != nil {
			return err
		}

		c1, err := parseCard(s1)
		if err != nil {
			return fmt.Errorf("hand %d: %w", i+1, err)
		}
		c2, err := parseCard(s2)
		if err != nil {
			return fmt.Errorf("hand %d: %w", i+1, err)
		}

		if debugEnable {
			log.Println(c1, c2)
		}

		count -= 2
		desk |= 1 << c1
		desk |= 1 << c2
		hands[i] = [...]tCard{c1, c2}
	}

	res := make([]tCard, 0, count)

	for card := tCard(0); card < 52; card++ {
		if desk&(1<<card) != 0 {
			continue
		}

		var maximum int
		for _, h := range hands {
			v := handPower(h[0], h[1], card)
			if maximum < v {
				maximum = v
			}
		}

		h := hands[0]
		if maximum == handPower(h[0], h[1], card) {
			res = append(res, card)
		}
	}

	fmt.Fprintln(out, len(res))
	for _, card := range res {
		fmt.Fprintln(out, card)
	}

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

	for i := 0; i < t; i++ {
		if debugEnable {
			log.Printf("--- task %d", i+1)
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
