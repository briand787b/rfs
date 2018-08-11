package server

import (
	"encoding/json"
	"fmt"
	"net"
	"os"

	"github.com/pkg/errors"
)

const (
	serversPath = "servers.json"
)

var (
	// ErrServerNotFound is the type of error returned when the
	// ServerStore cannot find the named server
	ErrServerNotFound = errors.New("Server does not exist in store")
)

// fileStore implements the Store
// interface with flat file storage in JSON
//
// The file which persists data will only
// be opened when actively reading or writing.
// Since reads/writes only happen when things
// change, and the server configs should change
// infrequently, every function invocation that
// opens the file will be close will close it
// upon exiting
type fileStore map[string]*Server

// NewFileStore sets up and returns a new
// Store using flat file storage
func NewFileStore() (Store, error) {
	return newFileStore(serversPath)
}

// newFileStore exists to facilitate testing of the
// fileStore struct
func newFileStore(path string) (fileStore, error) {
	fs := fileStore{}
	return fs, fs.getState(path)
}

// saveState adds the current state of sfs to disk
func (fs fileStore) saveState(filename string) error {
	fd, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		return errors.Wrap(err,
			"could not open file for file store",
		)
	}
	defer fd.Close()

	if err := fd.Truncate(0); err != nil {
		return errors.Wrap(err,
			"cannot truncate server store file")
	}

	if err := json.NewEncoder(fd).Encode(&fs); err != nil {
		return errors.Wrap(err,
			"could not encode fs into file as JSON",
		)
	}

	return nil
}

func (fs fileStore) getState(path string) error {
	fd, err := os.OpenFile(
		path,
		os.O_CREATE|os.O_RDWR,
		0600,
	)
	if err != nil {
		return errors.Wrap(err,
			"could not open file for file store")
	}
	defer fd.Close()

	fStat, err := fd.Stat()
	if err != nil {
		return errors.Wrapf(err,
			"could not get stats of file %s", path)
	}

	if fStat.Size() == 0 {
		// file is empty, fs should be empty too
		fmt.Println("opened empty file, returning empty store")
		fs = fileStore{}
		return nil
	}

	if err := json.NewDecoder(fd).Decode(&fs); err != nil {
		return errors.Wrap(err,
			"could not decode json in file to fileStore")
	}

	return nil
}

// Add adds a server and saves it to persistent storage
func (fs fileStore) AddServer(s *Server) error {
	return fs.add(s, serversPath)
}

func (fs fileStore) add(s *Server, path string) error {
	if s.Name == "" {
		return errors.New("Server must have name")
	}

	if net.ParseIP(s.IP) == nil {
		return errors.New("IP address is not valid")
	}

	if s.WorkingDir == "" {
		return errors.New("WorkingDir cannot be empty")
	}

	fs[s.Name] = s
	return fs.saveState(path)
}

func (fs fileStore) RemoveServer(name string) error {
	return fs.removeServer(name, serversPath)
}

func (fs fileStore) removeServer(name, path string) error {
	delete(fs, name)
	return fs.saveState(path)
}

func (fs fileStore) GetAllServers() (ss []*Server) {
	for _, s := range fs {
		ss = append(ss, s)
	}

	return
}

func (fs fileStore) GetServerByName(name string) (*Server, error) {
	s, ok := fs[name]
	if ok {
		return s, nil
	}

	return nil, ErrServerNotFound
}
