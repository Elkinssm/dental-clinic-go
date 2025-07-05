package dentist

import (
	d "Final_go/internal/clinic/dentists"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// ------------------------
// Mock Repository
// ------------------------
type mockRepo struct {
	GetFunc    func(id int) (d.Dentist, error)
	CreateFunc func(d.Dentist) (d.Dentist, error)
	UpdateFunc func(d.Dentist) (d.Dentist, error)
	DeleteFunc func(id int) error
	PatchFunc  func(d.Dentist) (d.Dentist, error)
}

func (m *mockRepo) GetDentistByID(id int) (d.Dentist, error) {
	return m.GetFunc(id)
}

func (m *mockRepo) CreateDentist(p d.Dentist) (d.Dentist, error) {
	return m.CreateFunc(p)
}

func (m *mockRepo) UpdateDentist(p d.Dentist) (d.Dentist, error) {
	return m.UpdateFunc(p)
}

func (m *mockRepo) DeleteDentist(id int) error {
	return m.DeleteFunc(id)
}

func (m *mockRepo) PatchDentist(p d.Dentist) (d.Dentist, error) {
	return m.PatchFunc(p)
}

// ------------------------
// Tests
// ------------------------

func TestCreateDentist_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	body := d.Dentist{Name: "John"}
	jsonBody, _ := json.Marshal(body)

	req, _ := http.NewRequest(http.MethodPost, "/dentists", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	handler := &DentistHandler{
		Repository: &mockRepo{
			CreateFunc: func(d d.Dentist) (d.Dentist, error) {
				d.ID = 1
				return d, nil
			},
		},
	}

	handler.CreateDentist(c)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "John")
}

func TestCreateDentist_InvalidJSON(t *testing.T) {
	gin.SetMode(gin.TestMode)

	req, _ := http.NewRequest(http.MethodPost, "/dentists", bytes.NewBufferString("invalid"))
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	handler := &DentistHandler{Repository: &mockRepo{}}
	handler.CreateDentist(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetDentistByID_InvalidID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	req, _ := http.NewRequest(http.MethodGet, "/dentists/abc", nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "abc"}}
	c.Request = req

	handler := &DentistHandler{Repository: &mockRepo{}}
	handler.GetDentistByID(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetDentistByID_NotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)
	req, _ := http.NewRequest(http.MethodGet, "/dentists/1", nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "1"}}
	c.Request = req

	handler := &DentistHandler{
		Repository: &mockRepo{
			GetFunc: func(id int) (d.Dentist, error) {
				return d.Dentist{}, errors.New("not found")
			},
		},
	}
	handler.GetDentistByID(c)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestGetDentistByID_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	req, _ := http.NewRequest(http.MethodGet, "/dentists/1", nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "1"}}
	c.Request = req

	handler := &DentistHandler{
		Repository: &mockRepo{
			GetFunc: func(id int) (d.Dentist, error) {
				return d.Dentist{ID: id, Name: "Test"}, nil
			},
		},
	}
	handler.GetDentistByID(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Test")
}

func TestUpdateDentist_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	dentist := d.Dentist{Name: "Updated"}
	body, _ := json.Marshal(dentist)

	req, _ := http.NewRequest(http.MethodPut, "/dentists/1", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "1"}}
	c.Request = req

	handler := &DentistHandler{
		Repository: &mockRepo{
			UpdateFunc: func(d d.Dentist) (d.Dentist, error) {
				return d, nil
			},
		},
	}
	handler.UpdateDentist(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Updated")
}

func TestDeleteDentist_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	req, _ := http.NewRequest(http.MethodDelete, "/dentists/1", nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "1"}}
	c.Request = req

	handler := &DentistHandler{
		Repository: &mockRepo{
			DeleteFunc: func(id int) error { return nil },
		},
	}
	handler.DeleteDentist(c)

	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestPatchDentist_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	patch := d.Dentist{Name: "Patched"}
	body, _ := json.Marshal(patch)

	req, _ := http.NewRequest(http.MethodPatch, "/dentists/1", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "1"}}
	c.Request = req

	handler := &DentistHandler{
		Repository: &mockRepo{
			GetFunc: func(id int) (d.Dentist, error) {
				return d.Dentist{ID: id, Name: "Original"}, nil
			},
			UpdateFunc: func(d d.Dentist) (d.Dentist, error) {
				return d, nil
			},
		},
	}
	handler.PatchDentist(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Patched")
}
