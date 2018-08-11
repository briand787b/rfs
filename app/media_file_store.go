package app

// mediaFileStore implements the MediaStore interface
// using the local file sysystem as persistence
type mediaFileStore map[string]*Media

func NewMediaFileStore() MediaStore {

}
