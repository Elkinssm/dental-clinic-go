package patients

type Patient struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	LastName         string `json:"last_name"`
	Address          string `json:"address"`
	DNI              string `json:"dni"`
	RegistrationDate string `json:"registration_date"`
}
