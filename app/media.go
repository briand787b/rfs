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

// // NewLocalMedia takes the filename of an external
// // file  (external to the working dir of the server)
// // and
// func NewLocalMedia(filename string) *Media {
// 	// open external file
// 	xf, err := os.Open(filename)
// 	if err != nil {
// 		return errors.Wrap(err,
// 			"could not open external file",
// 		)
// 	}

// }

// func (m *Media)
