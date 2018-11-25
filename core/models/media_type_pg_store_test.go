package models

import (
	"fmt"
	"reflect"
	"testing"
)

// CreateMediaType creates a MediaType.  If the passed
// MediaType pointer is nil, then the default MediaType is
// created.  Otherwise, the provided MediaType is created
// in the database.  The returned MediaType pointer should
// only be used if the provided MediaType pointer is nil.
func CreateMediaType(ts *TestStruct, mt *MediaType) *MediaType {
	if mt == nil {
		mt = &MediaType{
			Name: ts.T.Name(),
		}
	}

	if err := mediaTypeStore.Save(mt); err != nil {
		defer (*ts.CF)()
		ts.Fatal("could not create media: ", err)
	}

	fmt.Println("created media_type with ID: ", mt.ID)

	ts.CF.Add(func() {
		if err := mediaTypeStore.Delete(mt.ID); err != nil {
			ts.Fatal("could not delete created media_type: ", err)
		}
	})

	ts.ParentIDs[MEDIA_TYPE] = mt.ID

	return mt
}

func TestSaveGetByIDDeleteMediaType(t *testing.T) {
	mt := &MediaType{
		Name: t.Name(),
	}

	if mediaTypeStore == nil {
		t.Fatal("MediaType STORE IS NIL")
	}

	if err := mediaTypeStore.Save(mt); err != nil {
		t.Fatal(err)
	}

	defer mediaTypeStore.Delete(mt.ID)

	retMT, err := mediaTypeStore.GetByID(mt.ID)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(retMT, mt) {
		t.Fatalf("returned MediaType %+v does not equal saved media_type %+v\n",
			retMT, mt,
		)
	}

	if err := mediaTypeStore.Delete(mt.ID); err != nil {
		t.Fatal(err)
	}

	if _, err := mediaTypeStore.GetByID(mt.ID); err == nil {
		t.Fatalf("deleted media_type is still found by id %v\n", mt.ID)
	}
}
