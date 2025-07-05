package appointment

import (
	a "Final_go/internal/clinic/appointments"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// -----------------------------
// Mock del repositorio
// -----------------------------
type mockRepo struct {
	GetFunc             func(id int) (a.Appointment, error)
	CreateFunc          func(a.Appointment) (a.Appointment, error)
	UpdateFunc          func(a.Appointment) (a.Appointment, error)
	DeleteFunc          func(id int) error
	GetByPatientDNIFunc func(dni string) ([]a.Appointment, error)
}

func (m *mockRepo) GetAppointmentByID(id int) (a.Appointment, error) {
	return m.GetFunc(id)
}
func (m *mockRepo) CreateAppointment(app a.Appointment) (a.Appointment, error) {
	return m.CreateFunc(app)
}
func (m *mockRepo) UpdateAppointment(app a.Appointment) (a.Appointment, error) {
	return m.UpdateFunc(app)
}
func (m *mockRepo) DeleteAppointment(id int) error {
	return m.DeleteFunc(id)
}
func (m *mockRepo) GetAppointmentByPatientDNI(dni string) ([]a.Appointment, error) {
	return m.GetByPatientDNIFunc(dni)
}

// -----------------------------
// Test: GetAppointmentById
// -----------------------------
func TestGetAppointmentById_InvalidID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	req, _ := http.NewRequest(http.MethodGet, "/appointments/abc", nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "abc"}}
	c.Request = req

	handler := &AppointmentHandler{Repository: &mockRepo{}}
	handler.GetAppointmentById(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	ErrInvalidID := "Invalid ID"
	assert.Contains(t, w.Body.String(), ErrInvalidID)
}

func TestGetAppointmentById_NotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)
	req, _ := http.NewRequest(http.MethodGet, "/appointments/123", nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "123"}}
	c.Request = req

	handler := &AppointmentHandler{
		Repository: &mockRepo{
			GetFunc: func(id int) (a.Appointment, error) {
				return a.Appointment{}, errors.New("not found")
			},
		},
	}
	handler.GetAppointmentById(c)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), "Appointment not found")
}

func TestGetAppointmentById_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	req, _ := http.NewRequest(http.MethodGet, "/appointments/1", nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "1"}}
	c.Request = req

	handler := &AppointmentHandler{
		Repository: &mockRepo{
			GetFunc: func(id int) (a.Appointment, error) {
				return a.Appointment{ID: id, Description: "test"}, nil
			},
		},
	}
	handler.GetAppointmentById(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "test")
}

// -----------------------------
// Test: CreateAppointment
// -----------------------------
func TestCreateAppointment_InvalidJSON(t *testing.T) {
	gin.SetMode(gin.TestMode)
	body := bytes.NewBufferString(`invalid`)
	req, _ := http.NewRequest(http.MethodPost, "/appointments", body)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	handler := &AppointmentHandler{Repository: &mockRepo{}}
	handler.CreateAppointment(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestCreateAppointment_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	appointment := a.Appointment{ID: 1, Description: "test"}
	jsonBody, _ := json.Marshal(appointment)

	req, _ := http.NewRequest(http.MethodPost, "/appointments", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	handler := &AppointmentHandler{
		Repository: &mockRepo{
			CreateFunc: func(app a.Appointment) (a.Appointment, error) {
				return appointment, nil
			},
		},
	}
	handler.CreateAppointment(c)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "test")
}

// -----------------------------
// Test: UpdateAppointment
// -----------------------------
func TestUpdateAppointment_InvalidID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	req, _ := http.NewRequest(http.MethodPut, "/appointments/abc", nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "abc"}}
	c.Request = req

	handler := &AppointmentHandler{Repository: &mockRepo{}}
	handler.UpdateAppointment(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUpdateAppointment_InvalidJSON(t *testing.T) {
	gin.SetMode(gin.TestMode)
	req, _ := http.NewRequest(http.MethodPut, "/appointments/1", bytes.NewBufferString("invalid"))
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "1"}}
	c.Request = req

	handler := &AppointmentHandler{Repository: &mockRepo{}}
	handler.UpdateAppointment(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUpdateAppointment_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	app := a.Appointment{ID: 1, Description: "Updated"}
	body, _ := json.Marshal(app)

	req, _ := http.NewRequest(http.MethodPut, "/appointments/1", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "1"}}
	c.Request = req

	handler := &AppointmentHandler{
		Repository: &mockRepo{
			UpdateFunc: func(app a.Appointment) (a.Appointment, error) {
				return app, nil
			},
		},
	}
	handler.UpdateAppointment(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Updated")
}

// -----------------------------
// Test: DeleteAppointment
// -----------------------------
func TestDeleteAppointment_InvalidID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	req, _ := http.NewRequest(http.MethodDelete, "/appointments/abc", nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "abc"}}
	c.Request = req

	handler := &AppointmentHandler{Repository: &mockRepo{}}
	handler.DeleteAppointment(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestDeleteAppointment_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	req, _ := http.NewRequest(http.MethodDelete, "/appointments/1", nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "1"}}
	c.Request = req

	handler := &AppointmentHandler{
		Repository: &mockRepo{
			DeleteFunc: func(id int) error {
				return nil
			},
		},
	}
	handler.DeleteAppointment(c)

	assert.Equal(t, http.StatusNoContent, w.Code)
}

// -----------------------------
// Test: GetAppointmentByPatientDNI
// -----------------------------
func TestGetAppointmentByPatientDNI_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	req, _ := http.NewRequest(http.MethodGet, "/appointments?dni=123", nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	handler := &AppointmentHandler{
		Repository: &mockRepo{
			GetByPatientDNIFunc: func(dni string) ([]a.Appointment, error) {
				return []a.Appointment{{ID: 1, Description: "DNI test"}}, nil
			},
		},
	}
	handler.GetAppointmentByPatientDNI(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "DNI test")
}
