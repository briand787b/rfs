package server

// Server represents a physical server that hosts
// and/or consumes content.
//
// FileList is not persisted because it is dynamic
// by the nature of this program and its accuracy
// can only be guaranteed on fresh retrievals from
// the server in question
type Server struct {
	Name       string   `json:"name"`
	IP         string   `json:"ip_addr"`
	WorkingDir string   `json:"working_dir"` // absolute path
	FileList   []string `json:"-"`
}

// NewServer instantiates a new server.
// It does NOT save it to the persistent
// storage
func NewServer(name, ip, dir string) *Server {
	return &Server{
		Name:       name,
		IP:         ip,
		WorkingDir: dir,
	}
}

// func (s *Server) MarshalJSON() ([]byte, error) {
// 	return
// }

// func (s *Server) String() string {

// }
