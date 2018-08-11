package app

import (
	"github.com/briand787b/rfs/remote/api/http"
)

// Server represents a physical server that hosts content
//
//
type Server struct {
	Name       string     `json:"name"`
	IP         string     `json:"ip_addr"`
	WorkingDir string     `json:"working_dir"` // absolute path
	MediaStore MediaStore `json:"media"`       // name-value pairs
	Port       string     `json:"port"`
}

// NewServer instantiates a new server.
// It does NOT save it to the persistent
// storage
//
// NEEDS TO BE DIFFERENTIATED BTWN LOCAL AND REMOTE
func NewServer(name, ip, dir string) *Server {

	return &Server{
		Name:       name,
		IP:         ip,
		WorkingDir: dir,
		Port:       http.DefaultRemoteServerHTTPPort,
	}
}

// GetBaseURL returns the base url for the server
func (s *Server) GetBaseURL() string {
	return s.IP + ":" + s.Port + "/"
}
