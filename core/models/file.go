package models

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
