package app

// ServerStore is the interface that covers the management
// of the persistence layer for the servers
type ServerStore interface {
	AddServer(*Server) error
	RemoveServer(string) error
	GetAllServers() []*Server
	GetServerByName(string) (*Server, error)
}
