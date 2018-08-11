package server

import (
	"reflect"
	"testing"
)

func TestNewServer(t *testing.T) {
	actual := NewServer(
		testServerName,
		testServerIP,
		testWorkingDir,
	)

	expected := Server{
		Name:       testServerName,
		IP:         testServerIP,
		WorkingDir: testWorkingDir,
	}

	if !reflect.DeepEqual(*actual, expected) {
		t.Fatal("new server not equal to expected server")
	}
}
