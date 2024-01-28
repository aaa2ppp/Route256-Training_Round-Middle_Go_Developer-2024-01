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
			args{strings.NewReader(`5
9
3 2 1 0 -1 0 6 6 7
1
1000
7
1 2 3 4 5 6 7
7
1 3 5 7 9 11 13
11
100 101 102 103 19 20 21 22 42 41 40
`)},
			`8
3 -4 0 0 6 0 6 1
2
1000 0
2
1 6
14
1 0 3 0 5 0 7 0 9 0 11 0 13 0
6
100 3 19 3 42 -2
`,
			false,
			true,
		},
		// {"2",
		// 	args{strings.NewReader(``)},
		// 	``,
		// 	false,
		// 	false,
		// },
		// {"3",
		// 	args{strings.NewReader(``)},
		// 	``,
		// 	false,
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
