package appointment

import (
	"Final_go/internal/clinic/dentist"
	"Final_go/internal/clinic/patient"
)

type Appointment struct {
	ID          int             `json:"id"`
	Date        string          `json:"date" `
	Hour        string          `json:"hour" `
	Description string          `json:"description"`
	Patient     patient.Patient `json:"patient" `
	Dentist     dentist.Dentist `json:"dentist" `
}
