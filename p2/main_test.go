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
	}
	tests := []tTest{
		{"1",
			args{strings.NewReader(`10
10 9 2022
21 9 2022
29 2 2022
31 2 2022
29 2 2000
29 2 2100
31 11 1999
31 12 1999
29 2 2024
29 2 2023`)},
			`YES
YES
NO
NO
YES
NO
NO
YES
YES
NO`,
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
				t.Errorf("run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			gotOut := unsafeString(out.Bytes())
			if strings.TrimSpace(gotOut) != strings.TrimSpace(tt.wantOut) {
				t.Errorf("run() = --\n%v\n--, want --\n%v\n--", gotOut, tt.wantOut)
			}
		})
	}
}
