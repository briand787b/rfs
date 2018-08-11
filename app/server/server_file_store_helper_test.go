package server

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"sync"
	"testing"
)

const (
	testServerName = "test_server"
	testServerIP   = "127.0.0.1"
	testWorkingDir = "working_dir"
)

var (
	whiteSpaceRegx *regexp.Regexp
	wsrOnce        sync.Once
)

// **************************************************************
// Why does the tesdata directory exist?  testdata directories
// are not analyzed by the go compiler unless tests are being
// run.  This prevents test fixtures from being included in the
// binary that `go build` outputs.
// **************************************************************

// prepares for running fileStore tests
func setupFS() {
	validTestServer1 = &Server{
		Name:       testServerName + "1",
		IP:         testServerIP,
		WorkingDir: testWorkingDir,
	}

	validTestServer2 = &Server{
		Name:       testServerName + "2",
		IP:         "127.0.0.2",
		WorkingDir: testWorkingDir,
	}
}

func tearDownFS() {
	// all test file store files are suffixed with .test.json
	ns, err := filepath.Glob("*.test.json")
	if err != nil {
		fmt.Println("cannot glob filepath for test.json files")
		os.Exit(1)
	}

	for _, n := range ns {
		if err := os.Remove(n); err != nil {
			fmt.Printf("cannot remove file %s: %s", n, err)
			os.Exit(1)
		}
	}
}

func getTestFileStoreName(t *testing.T) string {
	return t.Name() + ".test.json"
}

func cutWhiteSpace(b []byte, t *testing.T) []byte {
	wsrOnce.Do(func() {
		var err error
		whiteSpaceRegx, err = regexp.Compile(`\s`)
		if err != nil {
			t.Fatal("could not compile regexp: ", err)
		}
	})

	return whiteSpaceRegx.ReplaceAll(b, []byte(""))
}

func getGoldenFilePath(t *testing.T) string {
	return filepath.Join("testdata", t.Name()+".golden")
}

func parseGoldenFile(t *testing.T) *bytes.Buffer {
	filepath := getGoldenFilePath(t)
	bs, err := ioutil.ReadFile(filepath)
	if err != nil {
		t.Fatalf("could not read golden file %s: %s", t.Name(), err)
	}

	pwd, err := os.Getwd()
	if err != nil {
		t.Fatal("could not get working dir: ", err)
	}

	b := &bytes.Buffer{}
	if err := template.Must(template.New(
		filepath,
	).Parse(string(cutWhiteSpace(bs, t)))).Execute(b, pwd); err != nil {
		t.Fatal("could not execute parsed template: ", err)
	}

	t.Logf("%v\n", b.String())

	return b
}

func overwriteGoldenFile(t *testing.T, actual []byte) {
	p := getGoldenFilePath(t)
	if err := ioutil.WriteFile(p, actual, 0644); err != nil {
		t.Fatal("could not overwrite file: ", err)
	}
}
