package http

import "net/http"

// DefaultRemoteServerHTTPPort is the port on which all servers should
// listen for inter-application http communication.  Master is
// the only server that should be listening on port 80
const DefaultRemoteServerHTTPPort = ":2269"

// Paths for Routing
const (
	// AddFilePath (POST) is the path for the creation of files on the remote server
	AddFilePath = "/files"
)

// Serve tells the remote server to liston on HTTP
func Serve() error {
	return http.ListenAndServe(DefaultRemoteServerHTTPPort, nil)
}
