package appointments

import (
	"Final_go/internal/clinic/dentists"
	"Final_go/internal/clinic/patients"
)

type Appointment struct {
	ID          int              `json:"id"`
	Date        string           `json:"date" `
	Hour        string           `json:"hour" `
	Description string           `json:"description"`
	Patient     patients.Patient `json:"patient" `
	Dentist     dentists.Dentist `json:"dentist" `
}
