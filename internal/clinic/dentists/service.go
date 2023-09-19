package dentists

type DentistRepository interface {
	CreateDentist(dentist Dentist) (Dentist, error)
	GetDentistByID(id int) (Dentist, error)
	UpdateDentist(dentist Dentist) (Dentist, error)
	DeleteDentist(id int) error
	PatchDentist(dentist Dentist) (Dentist, error)
}

type Service struct {
	repository DentistRepository
}

func NewDentistService(repository DentistRepository) *Service {
	return &Service{repository: repository}
}

func (s *Service) CreateDentist(dentist Dentist) (Dentist, error) {
	return s.repository.CreateDentist(dentist)
}

func (s *Service) GetDentistByID(id int) (Dentist, error) {
	return s.repository.GetDentistByID(id)
}

func (s *Service) UpdateDentist(dentist Dentist) (Dentist, error) {
	return s.repository.UpdateDentist(dentist)
}

func (s *Service) DeleteDentist(id int) error {
	return s.repository.DeleteDentist(id)
}

func (s *Service) PatchDentist(dentist Dentist) (Dentist, error) {
	return s.repository.PatchDentist(dentist)
}
