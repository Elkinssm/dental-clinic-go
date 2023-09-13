package dentist

type Dentist struct {
	ID       int    `json:"id"`
	Name     string `json:"name" binding:"required"`
	LastName string `json:"last_name"`
	License  string `json:"license" binding:"required"`
}
