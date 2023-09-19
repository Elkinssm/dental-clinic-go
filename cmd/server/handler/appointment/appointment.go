package appointment

import (
	a "Final_go/internal/clinic/appointments"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AppointmentHandler struct {
	Repository a.AppointmentRepository
}

func NewAppointmentHandler(repository a.AppointmentRepository) *AppointmentHandler {
	return &AppointmentHandler{Repository: repository}
}

func (h *AppointmentHandler) CreateAppointment(c *gin.Context) {
	var newAppointment a.Appointment

	if err := c.BindJSON(&newAppointment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdAppointment, err := h.Repository.CreateAppointment(newAppointment)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := gin.H{
		"id":          createdAppointment.ID,
		"date":        createdAppointment.Date,
		"hour":        createdAppointment.Hour,
		"description": createdAppointment.Description,
		"patients": gin.H{
			"id": createdAppointment.Patient.ID,
		},
		"dentists": gin.H{
			"id": createdAppointment.Dentist.ID,
		},
	}

	c.JSON(http.StatusCreated, response)
}

// @Summary 		Create a new appointment
// @Description 	Create a new appointment
// @Tags 			appointments
// @Accept  		json
// @Produce  		json
// @Router 			/appointments [post]
// @Security 		ApiKeyAuth
// @Failure 		400 {object} string "Invalid request"
// @Failure			500 {object} string "Internal error"
// @Failure 		401 {object} string "Unauthorized"
// @Failure 		403 {object} string "Forbidden"
// @Failure 		404 {object} string "Not found"
// @Failure 		409 {object} string "Conflict"
// @Failure 		422 {object} string "Unprocessable entity"
// @Failure 		default {object} string "Error"
// @Success 		200 {object} a.Appointment
func (h *AppointmentHandler) GetAppointmentById(c *gin.Context) {
	id, error := strconv.Atoi(c.Param("id"))
	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	foundAppointment, err := h.Repository.GetAppointmentByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Appointment not found"})
		return
	}
	c.JSON(http.StatusOK, foundAppointment)

}

// @Summary 		Update an appointment
// @Description 	Update an appointment
// @Tags 			appointments
// @Accept  		json
// @Produce  		json
// @Router 			/appointments/{id} [put]
// @Security 		ApiKeyAuth
// @Param 			id path int true "Appointment ID"
// @Param 			appointment body Appointment true "Appointment"
// @Failure 		400 {object} string "Invalid request"
// @Failure			500 {object} string "Internal error"
// @Failure 		401 {object} string "Unauthorized"
// @Failure 		403 {object} string "Forbidden"
// @Failure 		404 {object} string "Not found"
// @Failure 		409 {object} string "Conflict"
// @Failure 		422 {object} string "Unprocessable entity"
// @Failure 		default {object} string "Error"
// @Success 		200 {object} Appointment
func (h *AppointmentHandler) UpdateAppointment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var updatedAppointment a.Appointment
	if err := c.BindJSON(&updatedAppointment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedAppointment.ID = id
	updatedAppointment, err = h.Repository.UpdateAppointment(updatedAppointment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedAppointment)
}

func (h *AppointmentHandler) PatchAppointment(c *gin.Context) {
	// Obtener el ID de la cita de la URL
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// Obtener la cita existente por su ID
	existingAppointment, err := h.Repository.GetAppointmentByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Appointment not found"})
		return
	}

	// Decodificar el JSON de la solicitud y aplicar actualizaciones parciales
	var partialAppointment a.Appointment
	if err := c.BindJSON(&partialAppointment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Aplicar las actualizaciones parciales a la cita existente
	if partialAppointment.Date != "" {
		existingAppointment.Date = partialAppointment.Date
	}
	if partialAppointment.Hour != "" {
		existingAppointment.Hour = partialAppointment.Hour
	}
	if partialAppointment.Description != "" {
		existingAppointment.Description = partialAppointment.Description
	}
	// Actualizar la cita en el almacenamiento
	updatedAppointment, err := h.Repository.UpdateAppointment(existingAppointment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedAppointment)
}

func (h *AppointmentHandler) DeleteAppointment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = h.Repository.DeleteAppointment(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func (h *AppointmentHandler) GetAppointmentByPatientDNI(c *gin.Context) {
	dni := c.Query("dni")
	appointments, err := h.Repository.GetAppointmentByPatientDNI(dni)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, appointments)
}
