package patient

import (
	p "Final_go/internal/clinic/patients"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// ---------------------------
// Mock del repositorio
// ---------------------------
type mockRepo struct {
	GetFunc    func(id int) (p.Patient, error)
	CreateFunc func(p.Patient) (p.Patient, error)
	UpdateFunc func(p.Patient) (p.Patient, error)
	DeleteFunc func(id int) error
	PatchFunc  func(p.Patient) (p.Patient, error)
}

func (m *mockRepo) GetPatientByID(id int) (p.Patient, error) {
	return m.GetFunc(id)
}

func (m *mockRepo) CreatePatient(pt p.Patient) (p.Patient, error) {
	return m.CreateFunc(pt)
}

func (m *mockRepo) UpdatePatient(pt p.Patient) (p.Patient, error) {
	return m.UpdateFunc(pt)
}

func (m *mockRepo) DeletePatient(id int) error {
	return m.DeleteFunc(id)
}

func (m *mockRepo) PatchPatient(pt p.Patient) (p.Patient, error) {
	return m.PatchFunc(pt)
}

// ---------------------------
// Tests
// ---------------------------

func TestCreatePatient_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	input := p.Patient{Name: "John"}
	jsonBody, _ := json.Marshal(input)

	req, _ := http.NewRequest(http.MethodPost, "/patients", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	handler := &PatientHandler{
		Repository: &mockRepo{
			CreateFunc: func(p p.Patient) (p.Patient, error) {
				p.ID = 1
				return p, nil
			},
		},
	}
	handler.CreatePatient(c)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "John")
}

func TestCreatePatient_InvalidJSON(t *testing.T) {
	gin.SetMode(gin.TestMode)
	req, _ := http.NewRequest(http.MethodPost, "/patients", bytes.NewBufferString("invalid"))
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	handler := &PatientHandler{Repository: &mockRepo{}}
	handler.CreatePatient(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetPatientByID_InvalidID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	req, _ := http.NewRequest(http.MethodGet, "/patients/abc", nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "abc"}}
	c.Request = req

	handler := &PatientHandler{Repository: &mockRepo{}}
	handler.GetPatientByID(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetPatientByID_NotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)
	req, _ := http.NewRequest(http.MethodGet, "/patients/1", nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "1"}}
	c.Request = req

	handler := &PatientHandler{
		Repository: &mockRepo{
			GetFunc: func(id int) (p.Patient, error) {
				return p.Patient{}, errors.New("not found")
			},
		},
	}
	handler.GetPatientByID(c)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestGetPatientByID_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	req, _ := http.NewRequest(http.MethodGet, "/patients/1", nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "1"}}
	c.Request = req

	handler := &PatientHandler{
		Repository: &mockRepo{
			GetFunc: func(id int) (p.Patient, error) {
				return p.Patient{ID: id, Name: "Test"}, nil
			},
		},
	}
	handler.GetPatientByID(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Test")
}

func TestUpdatePatient_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	input := p.Patient{Name: "Updated"}
	body, _ := json.Marshal(input)

	req, _ := http.NewRequest(http.MethodPut, "/patients/1", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "1"}}
	c.Request = req

	handler := &PatientHandler{
		Repository: &mockRepo{
			UpdateFunc: func(p p.Patient) (p.Patient, error) {
				return p, nil
			},
		},
	}
	handler.UpdatePatient(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Updated")
}

func TestPatchPatient_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	patch := p.Patient{Name: "Patched"}
	body, _ := json.Marshal(patch)

	req, _ := http.NewRequest(http.MethodPatch, "/patients/1", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "1"}}
	c.Request = req

	handler := &PatientHandler{
		Repository: &mockRepo{
			GetFunc: func(id int) (p.Patient, error) {
				return p.Patient{ID: id, Name: "Original"}, nil
			},
			UpdateFunc: func(p p.Patient) (p.Patient, error) {
				return p, nil
			},
		},
	}
	handler.PatchPatient(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Patched")
}
