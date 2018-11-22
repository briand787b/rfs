package models

// Media (as defined in README) This application defines the
// term 'media' (ignoring correct singularization) to
// represent a grouping of any number of entertainment artifacts
// that can be logically organized into a single entity.
// Organization of media is hierarchical and one media may be
// considered a 'parent' to zero or more media and a 'child' of
// exactly one media (unless it is basal to its hierarchy tree).
type Media struct {
	// DB-mapped
	ID            int    `json:"id"`
	Name          string `json:"name"`
	ParentID      int    `json:"parent_id"`
	FeatureFileID int    `json:"feature_file_id"`
	ReleaseYear   int    `json:"release_year"`

	// Cached -- might not want to make these exported
	ParentMedia *Media
	ChildMedia  []Media
	FeatureFile *File
	Files       []File
}

// Save saves a Media to storage
func (m *Media) Save(ms MediaStore) error {
	return ms.Save(m)
}
