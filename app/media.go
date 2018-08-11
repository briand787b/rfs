package app

import (
	"os"
)

// Media represents a directory devoted to a single piece
// of media. The directory can contain other files, but
// only one file, the 'core' file, can be considered the
// defining "media".  This file is the what the checksum
// is calculated against.
//
// All Media exist in a flat hierarchy.  Media are not
// organized by their type (yet).
type Media struct {
	Name     string
	Checksum string
	files    []*os.File
}

func NewLocalMedia() *Media {

}

// func (m *Media)
