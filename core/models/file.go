package models

import (
	"github.com/pkg/errors"
)

// File (from README) A file is a concrete artifact that relates
// to exactly one media.  It has a 1:1 relationship with a file
// in the traditional sense.
type File struct {
	// DB-mapped fields
	ID          int      `json:"id"`
	MediaID     int      `json:"media_id"`
	MD5Checksum [16]byte `json:"md5_checksum"`

	// Cached fields
	Media *Media
}

// Save saves a File to storage
func (f *File) Save(fs FileStore) error {
	return fs.Save(f)
}

// GetMD5Checksum computes the md5 checksum of the file
func (f *File) GetMD5Checksum() ([16]byte, error) {
	// find file on actual machine (may not be this one)
	// read the file
	// compute md5 checksum against contents

	return [16]byte{}, errors.New("NOT IMPLEMENTED")
}
