package main

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"unsafe"
)

const testDir = "./test_data"

func Test_run(t *testing.T) {
	unsafeString := func(buf []byte) string {
		return *(*string)(unsafe.Pointer(&buf))
	}
	type args struct {
		in io.Reader
	}
	type tTest struct {
		name    string
		args    args
		wantOut string
		wantErr bool
	}
	tests := []tTest{
		{"1",
			args{strings.NewReader(`4
14
75 22 I'm fine. Thank you.
84 82     Ciao!
26 22 So-so
45 26 What's wrong?
22 -1 How are you?
72 45 Maybe I got sick
81 72 I wish you a speedy recovery!
97 26   Stick it!
2 97 Thanks
47 72 I also got sick recently.
25 -1 Hi!
82 -1 Bye
17 82 Good day!
29 72 Visit the doctor
8
5 4 e
6 5 f
7 6 g
1 -1 a
2 1 b
3 2 c
4 3 d
8 7 h
6
10 -1 x
20 10 x
40 -1 x
50 -1 x
11 20 x
30 10 x
1
1000000000 -1 root
`)},
			`How are you?
|
|--So-so
|  |
|  |--What's wrong?
|  |  |
|  |  |--Maybe I got sick
|  |     |
|  |     |--Visit the doctor
|  |     |
|  |     |--I also got sick recently.
|  |     |
|  |     |--I wish you a speedy recovery!
|  |
|  |--  Stick it!
|     |
|     |--Thanks
|
|--I'm fine. Thank you.

Hi!

Bye
|
|--Good day!
|
|--    Ciao!

a
|
|--b
   |
   |--c
      |
      |--d
         |
         |--e
            |
            |--f
               |
               |--g
                  |
                  |--h

x
|
|--x
|  |
|  |--x
|
|--x

x

x

root
`,
			false,
		},
		// {"2",
		// 	args{strings.NewReader(``)},
		// 	``,
		// 	false,
		// },
		// {"3",
		// 	args{strings.NewReader(``)},
		// 	``,
		// 	false,
		// },
		// TODO: Add test cases.
	}

	// load the full test suite
	for i := 1; ; i++ {
		testFile := filepath.Join(testDir, strconv.Itoa(i))
		test, err := os.ReadFile(testFile)
		if err != nil && os.IsNotExist(err) {
			break
		}
		if err != nil {
			t.Fatal(err)
		}
		answer, err := os.ReadFile(testFile + ".a")
		if err != nil {
			t.Fatal(err)
		}
		tests = append(tests, tTest{
			testFile,
			args{bytes.NewBuffer(test)},
			unsafeString(answer),
			false,
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			if err := run(tt.args.in, out); (err != nil) != tt.wantErr {
				t.Fatalf("run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			gotOut := unsafeString(out.Bytes())
			gotLines := splitToLines(gotOut)
			wantLines := splitToLines(tt.wantOut)

			if !reflect.DeepEqual(gotLines, wantLines) {
				t.Fatalf("run() = --\n%v\n--, want --\n%v\n--", gotOut, tt.wantOut)
			}
		})
	}
}

func splitToLines(text string) []string {
	text = strings.TrimSuffix(text, "\n")
	lines := strings.Split(text, "\n")
	for i := range lines {
		lines[i] = strings.TrimRight(lines[i], "\r")
	}
	return lines
}
