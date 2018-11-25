package models

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

// CreateMedia creates a Media.  If the passed
// Media pointer is nil, then the default Media is
// created.  Otherwise, the provided Media is created
// in the database.  The returned Media pointer should
// only be used if the provided Media pointer is nil.
func CreateMedia(ts *TestStruct, m *Media) *Media {
	if m == nil {
		m = &Media{}
	}

	if err := mediaStore.Save(m); err != nil {
		defer (*ts.CF)()
		ts.Fatal("could not create media: ", err)
	}

	fmt.Println("created media with ID: ", m.ID)

	ts.CF.Add(func() {
		if err := mediaStore.Delete(m.ID); err != nil {
			ts.T.Fatal("could not delete created media: ", err)
		}
	})

	ts.ParentIDs[MEDIA] = m.ID

	return m
}

func TestSaveGetByIDDeleteParentlessChildlessMedia(t *testing.T) {
	m := &Media{
		Name:        t.Name(),
		ReleaseYear: time.Now().Year(),
	}

	if mediaStore == nil {
		t.Fatal("MEDIA STORE IS NIL")
	}

	if err := mediaStore.Save(m); err != nil {
		t.Fatal(err)
	}

	retM, err := mediaStore.GetByID(m.ID)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(retM, m) {
		t.Fatalf("returned Media %+v does not equal saved media %+v\n",
			retM, m,
		)
	}

	if err := mediaStore.Delete(m.ID); err != nil {
		t.Fatal(err)
	}

	if _, err := mediaStore.GetByID(m.ID); err == nil {
		t.Fatalf("deleted media is still found by id %v\n", m.ID)
	}
}
