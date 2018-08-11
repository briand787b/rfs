package server

// Store is the interface that covers the management
// of the persistence layer for the servers
type Store interface {
	AddServer(*Server) error
	RemoveServer(string) error
	GetAllServers() []*Server
	GetServerByName(string) (*Server, error)
}
