package modeltest

import "github.com/briand787b/rfs/core/models"

// MediaStoreMock is a mocked implementation of a MediaStore
type MediaStoreMock struct {
	// function return info
	GetByIDM *models.Media
	GetByIDE error
	SaveM    *models.Media
	SaveE    error
	DeleteM  *models.Media
	DeleteE  error

	// Spy fields
	CallCount int
}

// NewMediaStoreMockGetByID returns a MediaStore ready to mock the GetByID method
func NewMediaStoreMockGetByID(m *models.Media, e error) *MediaStoreMock {
	return &MediaStoreMock{
		GetByIDM: m,
		GetByIDE: e,
	}
}

// NewMediaStoreMockSave returns a MediaStore ready to mock the Save method
func NewMediaStoreMockSave(m *models.Media, e error) *MediaStoreMock {
	return &MediaStoreMock{
		SaveM: m,
		SaveE: e,
	}
}

// NewMediaStoreMockDelete returns a MediaStore ready to mock the Delete method
func NewMediaStoreMockDelete(m *models.Media, e error) *MediaStoreMock {
	return &MediaStoreMock{
		DeleteM: m,
		DeleteE: e,
	}
}

// GetByID returns the MediaStoreMock's GetByIDM and GetByIDE
func (msm *MediaStoreMock) GetByID(id int) (*models.Media, error) {
	msm.CallCount++
	return msm.GetByIDM, msm.GetByIDE
}

// Save returns the MediaStoreMock's SaveM and SaveE
func (msm *MediaStoreMock) Save(id int) (*models.Media, error) {
	msm.CallCount++
	return msm.SaveM, msm.SaveE
}

// Delete returns the MediaStoreMock's DeleteF and DeleteE
func (msm *MediaStoreMock) Delete(id int) (*models.Media, error) {
	msm.CallCount++
	return msm.DeleteM, msm.DeleteE
}
