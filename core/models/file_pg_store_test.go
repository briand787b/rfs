package models

import (
	"fmt"
	"testing"
)

// I might not even need this
//
// CreateFile creates a File.  If the passed
// File pointer is nil, then the default File is
// created.  Otherwise, the provided File is created
// in the database.  The returned File pointer should
// only be used if the provided File pointer is nil.
func CreateFile(ts *TestStruct, f *File) *File {
	if f == nil {
		f = &File{}
	}

	if err := fileStore.Save(f); err != nil {
		defer (*ts.CF)()
		ts.Fatal("could not create File: ", err)
	}

	fmt.Println("created File with ID: ", f.ID)

	ts.CF.Add(func() {
		if err := fileStore.Delete(f.ID); err != nil {
			ts.Fatal("could not delete created File: ", err)
		}
	})

	ts.ParentIDs[FILE] = f.ID

	return f
}

func TestSaveGetFileByIDDelete(t *testing.T) {

	// f := &File{

	// }
	t.Fail()
}
