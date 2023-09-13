package appointment

import(
	"Final_go/internal/clinic/patient"
	"Final_go/internal/clinic/dentist"
)

type Appointment struct {
	ID          int     `json:"id"`
	Date        string  `json:"date" binding:"required"`
	Hour        string  `json:"hour" binding:"required"`
	Description string  `json:"description"`
	Patient     patient.Patient `json:"patient" binding:"required"`
	Dentist     dentist.Dentist `json:"dentist" binding:"required"`
}


