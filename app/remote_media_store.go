package app

import (
	"fmt"
	"net/http"

	rhttp "github.com/briand787b/rfs/remote/api/http"
	"github.com/pkg/errors"
)

// remoteMediaStore is an implementation of FileStore that
// persist the File structs on the remote filesystem
type remoteMediaStore map[string]*Media

// AddFileHTTP adds a file to the remote file system through HTTP call
func (rfs remoteMediaStore) AddMediaHTTP(s *Server, f *Media) error {
	// get file from local file system

	// build request for remote file server
	req, err := http.NewRequest("POST", s.GetBaseURL()+rhttp.AddFilePath, nil)
	if err != nil {
		return errors.Wrap(err,
			"could not build new request to remote file store",
		)
	}

	fmt.Printf("Request: %+v\n", req)

	// send request to remote server

	// add file to internal map

	// return error
	return nil
}
