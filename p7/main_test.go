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
			args{strings.NewReader(`7
8
7
8
1,7,1
8
1-5,1,7-7
10
1-5
10
1,2,3,4,5,6,8,9,10
3
1-2
100
1-2,3-7,10-20,100`)},
			`1-6,8
2-6,8
6,8
6-10
7
3
8-9,21-99`,
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
