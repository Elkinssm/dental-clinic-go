package patient

import (
	p "Final_go/internal/clinic/patients"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type PatientHandler struct {
	Repository p.PatientRepository
}

func NewPatientHandler(repository *p.Service) *PatientHandler {
	return &PatientHandler{Repository: repository}
}

func (h *PatientHandler) CreatePatient(c *gin.Context) {
	var newPatient p.Patient
	if err := c.BindJSON(&newPatient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdPatient, err := h.Repository.CreatePatient(newPatient)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdPatient)
}

func (h *PatientHandler) GetPatientByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	foundPatient, err := h.Repository.GetPatientByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}

	c.JSON(http.StatusOK, foundPatient)
}

func (h *PatientHandler) UpdatePatient(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var updatedPatient p.Patient
	if err := c.BindJSON(&updatedPatient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedPatient.ID = id
	updatedPatient, err = h.Repository.UpdatePatient(updatedPatient)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedPatient)
}

func (h *PatientHandler) PatchPatient(c *gin.Context) {
	// Obtener el ID del paciente de la URL
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// Obtener el paciente existente por su ID
	existingPatient, err := h.Repository.GetPatientByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}

	// Decodificar el JSON de la solicitud y aplicar actualizaciones parciales
	var partialPatient p.Patient
	if err := c.BindJSON(&partialPatient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Aplicar las actualizaciones parciales al paciente existente
	if partialPatient.Name != "" {
		existingPatient.Name = partialPatient.Name
	}
	if partialPatient.LastName != "" {
		existingPatient.LastName = partialPatient.LastName
	}
	if partialPatient.Address != "" {
		existingPatient.Address = partialPatient.Address
	}
	if partialPatient.DNI != "" {
		existingPatient.DNI = partialPatient.DNI
	}
	if partialPatient.RegistrationDate != "" {
		existingPatient.RegistrationDate = partialPatient.RegistrationDate
	}

	// Actualizar el paciente en el almacenamiento
	updatedPatient, err := h.Repository.UpdatePatient(existingPatient)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedPatient)
}

func (h *PatientHandler) DeletePatient(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = h.Repository.DeletePatient(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}