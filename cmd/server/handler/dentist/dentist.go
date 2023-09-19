package dentist

import (
	d "Final_go/internal/clinic/dentists"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type DentistHandler struct {
	Repository d.DentistRepository
}

func NewDentistHandler(repository d.DentistRepository) *DentistHandler {
	return &DentistHandler{Repository: repository}
}

// @Summary Create a new dentist
// @Description Create a new dentist
// @Tags dentists
func (h *DentistHandler) CreateDentist(c *gin.Context) {
	var newDentist d.Dentist
	if err := c.BindJSON(&newDentist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdDentist, err := h.Repository.CreateDentist(newDentist)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdDentist)
}

func (h *DentistHandler) GetDentistByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	foundDentist, err := h.Repository.GetDentistByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Dentist not found"})
		return
	}

	c.JSON(http.StatusOK, foundDentist)
}

func (h *DentistHandler) UpdateDentist(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var updatedDentist d.Dentist
	if err := c.BindJSON(&updatedDentist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedDentist.ID = id
	updatedDentist, err = h.Repository.UpdateDentist(updatedDentist)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedDentist)
}

func (h *DentistHandler) PatchDentist(c *gin.Context) {
	// Obtener el ID del dentista de la URL
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// Obtener el dentista existente por su ID
	existingDentist, err := h.Repository.GetDentistByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Dentist not found"})
		return
	}

	// Decodificar el JSON de la solicitud y aplicar actualizaciones parciales
	var partialDentist d.Dentist
	if err := c.BindJSON(&partialDentist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Aplicar las actualizaciones parciales al dentista existente
	if partialDentist.Name != "" {
		existingDentist.Name = partialDentist.Name
	}
	if partialDentist.LastName != "" {
		existingDentist.LastName = partialDentist.LastName
	}
	if partialDentist.License != "" {
		existingDentist.License = partialDentist.License
	}

	// Actualizar el dentista en el almacenamiento
	updatedDentist, err := h.Repository.UpdateDentist(existingDentist)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedDentist)
}

func (h *DentistHandler) DeleteDentist(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = h.Repository.DeleteDentist(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
