package modeltest

import "github.com/briand787b/rfs/core/models"

// MediaTypeStoreMock is a mocked implementation of MediaTypeStore
type MediaTypeStoreMock struct {
	// function return info
	GetByIDMT *models.MediaType
	GetByIDE error
	SaveMT    *models.MediaType
	SaveE    error
	DeleteMT  *models.MediaType
	DeleteE  error

	// Spy fields
	CallCount int
}

// NewMediaTypeStoreMockGetByID returns a MediaTypeStore ready to mock the GetByID method
func NewMediaTypeStoreMockGetByID(mt *models.MediaType, e error) *MediaTypeStoreMock {
	return &MediaTypeStoreMock{
		GetByIDMT: mt,
		GetByIDE: e,
	}
}

// NewMediaTypeStoreMockSave returns a MediaTypeStore ready to mock the Save method
func NewMediaTypeStoreMockSave(mt *models.MediaType, e error) *MediaTypeStoreMock {
	return &MediaTypeStoreMock{
		SaveMT: mt,
		SaveE: e,
	}
}

// NewMediaTypeStoreMockDelete returns a MediaTypeStore ready to mock the Delete method
func NewMediaTypeStoreMockDelete(mt *models.MediaType, e error) *MediaTypeStoreMock {
	return &MediaTypeStoreMock{
		DeleteMT: mt,
		DeleteE: e,
	}
}

// GetByID returns the MediaTypeStoreMock's GetByIDMT and GetByIDE
func (mtsm *MediaTypeStoreMock) GetByID(id int) (*models.MediaType, error) {
	mtsm.CallCount++
	return mtsm.GetByIDMT, mtsm.GetByIDE
}

// Save returns the MediaTypeStoreMock's SaveMT and SaveE
func (mtsm *MediaTypeStoreMock) Save(id int) (*models.MediaType, error) {
	mtsm.CallCount++
	return mtsm.SaveMT, mtsm.SaveE
}

// Delete returns the MediaTypeStoreMock's DeleteMT and DeleteE
func (mtsm *MediaTypeStoreMock) Delete(id int) (*models.MediaType, error) {
	mtsm.CallCount++
	return mtsm.DeleteMT, mtsm.DeleteE
}
