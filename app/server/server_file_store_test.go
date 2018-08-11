package server

import (
	"bytes"
	"flag"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

var (
	validTestServer1 *Server
	validTestServer2 *Server
)

var update = flag.Bool("update", false, "update .golden files")

func TestMain(m *testing.M) {
	setupFS()
	exitStatus := m.Run()
	tearDownFS()
	os.Exit(exitStatus)
}

// Verify that no error is generated when
// instantiating new file-based server store
func TestGetNewFileStore(t *testing.T) {
	n := getTestFileStoreName(t)
	if _, err := newFileStore(n); err != nil {
		t.Fatal("unable to instantiate new file server store: ", err)
	}
}

func TestFileStoreSaveFailsNoName(t *testing.T) {
	n := getTestFileStoreName(t)
	fs, err := newFileStore(n)
	if err != nil {
		t.Fatal("unable to instantiate new file server store: ", err)
	}

	if err := fs.add(&Server{
		IP:         "0.0.0.0",
		WorkingDir: "/valid/path",
	}, t.Name()); err == nil {
		t.Fatal("should not be able to save server with no name")
	}
}

func TestFileStoreSaveFailsBadIPv4(t *testing.T) {
	n := getTestFileStoreName(t)
	fs, err := newFileStore(n)
	if err != nil {
		t.Fatal("unable to instantiate new file server store: ", err)
	}

	badIPs := [][]string{
		[]string{"", "IP is empty string"},
		[]string{"0.0.0", "not enough dots"},
		[]string{"0.0.0.cat", "not a number"},
		[]string{"0.0.0.256", "number too large"},
		[]string{"0.0.0.-1", "number too small"},
	}

	for _, ip := range badIPs {
		if err := fs.add(&Server{
			Name:       testServerName,
			IP:         ip[0],
			WorkingDir: "/valid",
		}, t.Name()); err == nil {
			t.Fatal("should not be able to save because: ", ip[1])
		}
	}
}

func TestFileStoreSaveFailsNoWorkingDir(t *testing.T) {
	n := getTestFileStoreName(t)
	fs, err := newFileStore(n)
	if err != nil {
		t.Fatal("unable to instantiate new file server store: ", err)
	}

	if err := fs.add(&Server{
		IP:   "0.0.0.0",
		Name: "ValidName",
	}, n); err == nil {
		t.Fatal("should not be able to save server with no working dir")
	}
}

func TestFileStoreAddServerFromCleanState(t *testing.T) {
	n := getTestFileStoreName(t)
	fs, err := newFileStore(n)
	if err != nil {
		t.Fatal("unable to instantiate new file server store: ", err)
	}

	if err := fs.add(validTestServer1, n); err != nil {
		t.Fatal("could not save single validTestServer1: ", err)
	}

	actual, err := ioutil.ReadFile(n)
	if err != nil {
		t.Fatal("could not open written JSON file: ", err)
	}

	tActual := cutWhiteSpace(actual, t)
	t.Logf("tActual: %v\n", string(tActual))

	if *update {
		overwriteGoldenFile(t, tActual)
	}

	goldenFile := parseGoldenFile(t).Bytes()

	// validate saved JSON is correct
	if bytes.Compare(goldenFile, tActual) != 0 {
		t.Fatal("written JSON file does not match golden file")
	}
}

func TestFileStoreAddSecondServer(t *testing.T) {
	n := getTestFileStoreName(t)
	fs, err := newFileStore(n)
	if err != nil {
		t.Fatal("unable to instantiate new file server store: ", err)
	}

	// first add, starts at 0
	if err := fs.add(validTestServer1, n); err != nil {
		t.Fatal("could not save single validTestServer1")
	}

	// second add, starts at 1
	if err := fs.add(validTestServer2, n); err != nil {
		t.Fatal("could not save single validTestServer1")
	}

	actual, err := ioutil.ReadFile(n)
	if err != nil {
		t.Fatal("could not open written JSON file: ", err)
	}

	tActual := cutWhiteSpace(actual, t)
	t.Logf("tActual: %v\n", string(tActual))

	if *update {
		overwriteGoldenFile(t, tActual)
	}

	goldenFile := parseGoldenFile(t).Bytes()

	// validate saved JSON is correct
	if bytes.Compare(goldenFile, tActual) != 0 {
		// t.Logf("actual: %v\n", actual)
		// t.Logf("expected: %v\n", goldenFile)
		t.Fatal("written JSON file does not match golden file")
	}
}

func TestFileStoreRemoveFirstServer(t *testing.T) {
	n := getTestFileStoreName(t)
	fs, err := newFileStore(n)
	if err != nil {
		t.Fatal("unable to instantiate new file server store: ", err)
	}

	// first add, starts at 0
	if err := fs.add(validTestServer1, n); err != nil {
		t.Fatal("could not save single validTestServer1")
	}

	// second add, starts at 1
	if err := fs.add(validTestServer2, n); err != nil {
		t.Fatal("could not save single validTestServer1")
	}

	if err := fs.removeServer(validTestServer1.Name, n); err != nil {
		t.Fatal("could not remove first server: ", err)
	}

	actual, err := ioutil.ReadFile(n)
	if err != nil {
		t.Fatal("could not open written JSON file: ", err)
	}

	tActual := cutWhiteSpace(actual, t)
	t.Logf("tActual: %v\n", string(tActual))

	if *update {
		overwriteGoldenFile(t, tActual)
	}

	goldenFile := parseGoldenFile(t).Bytes()

	// validate saved JSON is correct
	if bytes.Compare(goldenFile, tActual) != 0 {
		t.Fatal("written JSON file does not match golden file")
	}
}

func TestGetAllServers(t *testing.T) {
	n := getTestFileStoreName(t)
	fs, err := newFileStore(n)
	if err != nil {
		t.Fatal("unable to instantiate new file server store: ", err)
	}

	exp := fileStore{
		validTestServer1.Name: validTestServer1,
		validTestServer2.Name: validTestServer2,
	}

	t.Log("number of servers upon initialization: ", len(fs))
	var count int
	for _, s := range exp {
		count++
		t.Log("iteration #", count)

		if err := fs.add(s, n); err != nil {
			t.Fatalf("could not save server #%v", *s)
		}

		ss := fs.GetAllServers()

		// ensure that # of servers matches call count to fs.add
		if len(ss) != count {
			t.Log("length of file store: ", len(ss))
			t.Fatalf("number of servers should equal %v, is %v", count, len(ss))
		}

		// ensure the recently saved server can be found in return list
		var act *Server
		for _, sRet := range ss {
			if sRet.Name == s.Name {
				act = sRet
			}
		}

		if act == nil {
			t.Fatal("could not find recently saved Server from GetAllServers")
		}

		if !reflect.DeepEqual(*act, *s) {
			t.Logf("saved server: %+v\n", *s)
			t.Logf("retrieved server: %+v\n", *act)
			t.Fatal("saved server is not equal to retrieved server")
		}
	}
}

func TestGetServerByName(t *testing.T) {
	n := getTestFileStoreName(t)
	fs, err := newFileStore(n)
	if err != nil {
		t.Fatal("unable to instantiate new file server store: ", err)
	}

	if err := fs.add(validTestServer1, n); err != nil {
		t.Fatal("could not save single validTestServer1")
	}

	for _, s := range []struct {
		name string
		err  error
	}{
		{
			validTestServer1.Name,
			nil,
		}, {
			"name_does_not_exist",
			ErrServerNotFound,
		}, {
			"",
			ErrServerNotFound,
		},
	} {
		if _, err := fs.GetServerByName(s.name); err != s.err {
			t.Fatalf("expected error response for name %s to be %s, was %s",
				s.name, err, s.err)
		}
	}
}
