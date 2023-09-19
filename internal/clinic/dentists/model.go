package dentists

type Dentist struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	License  string `json:"license" `
}
