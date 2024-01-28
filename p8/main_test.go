package main

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
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
		debug   bool
	}
	tests := []tTest{
		{"1",
			args{strings.NewReader(`4
2
TS TC
AD AH
3
2H 3H
9S 9C
4D QS
3
4C 7H
4H 4D
6S 6H
3
2S 3H
2C 2D
3C 3D
`)},
			`2
TD
TH
0
3
7S
7C
7D
0
`,
			false,
			true,
		},
		{"2",
			args{strings.NewReader(`1
7
AS AC
AD AH
KS JH
9D 9C
5H 5D
3C 3S
TC TH
`)},
			`30
2S
2C
2D
2H
4S
4C
4D
4H
6S
6C
6D
6H
7S
7C
7D
7H
8S
8C
8D
8H
JS
JC
JD
QS
QC
QD
QH
KC
KD
KH
`,
			false,
			true,
		},
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
			false,
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			debugEnable = tt.debug
			out := &bytes.Buffer{}
			if err := run(tt.args.in, out); (err != nil) != tt.wantErr {
				t.Errorf("run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			gotOut := unsafeString(out.Bytes())
			if strings.TrimSuffix(gotOut, "\n") != strings.TrimSuffix(tt.wantOut, "\n") {
				t.Errorf("run() = --\n%v\n--, want --\n%v\n--", gotOut, tt.wantOut)
			}
		})
	}
}
