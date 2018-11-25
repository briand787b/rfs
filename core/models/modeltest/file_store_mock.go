package modeltest

import (
	"github.com/briand787b/rfs/core/models"
)

// FileStoreMock is a mocked implementation of a FileStore
// intended to only be used for tests
type FileStoreMock struct {
	// function return info
	GetByIDF *models.File
	GetByIDE error
	SaveF    *models.File
	SaveE    error
	DeleteF  *models.File
	DeleteE  error

	// Spy fields
	CallCount int
}

// NewFileStoreMockGetByID returns a FileStore ready to mock the GetByID method
func NewFileStoreMockGetByID(f *models.File, e error) *FileStoreMock {
	return &FileStoreMock{
		GetByIDF: f,
		GetByIDE: e,
	}
}

// NewFileStoreMockSave returns a FileStore ready to mock the Save method
func NewFileStoreMockSave(f *models.File, e error) *FileStoreMock {
	return &FileStoreMock{
		SaveF: f,
		SaveE: e,
	}
}

// NewFileStoreMockDelete returns a FileStore ready to mock the Delete method
func NewFileStoreMockDelete(f *models.File, e error) *FileStoreMock {
	return &FileStoreMock{
		DeleteF: f,
		DeleteE: e,
	}
}

// GetByID returns the FileStoreMock's GetByIDF and GetByIDE
func (fsm *FileStoreMock) GetByID(id int) (*models.File, error) {
	fsm.CallCount++
	return fsm.GetByIDF, fsm.GetByIDE
}

// Save returns the FileStoreMock's SaveF and SaveE
func (fsm *FileStoreMock) Save(id int) (*models.File, error) {
	fsm.CallCount++
	return fsm.SaveF, fsm.SaveE
}

// Delete returns the FileStoreMock's DeleteF and DeleteE
func (fsm *FileStoreMock) Delete(id int) (*models.File, error) {
	fsm.CallCount++
	return fsm.DeleteF, fsm.DeleteE
}
