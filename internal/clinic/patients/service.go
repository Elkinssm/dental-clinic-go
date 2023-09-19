package patients

type PatientRepository interface {
	CreatePatient(patient Patient) (Patient, error)
	GetPatientByID(id int) (Patient, error)
	UpdatePatient(patient Patient) (Patient, error)
	DeletePatient(id int) error
	PatchPatient(patient Patient) (Patient, error)
}

type Service struct {
	repository PatientRepository
}

func NewPatientService(repository PatientRepository) *Service {
	return &Service{repository: repository}
}

func (s *Service) CreatePatient(patient Patient) (Patient, error) {
	return s.repository.CreatePatient(patient)
}

func (s *Service) GetPatientByID(id int) (Patient, error) {
	return s.repository.GetPatientByID(id)
}

func (s *Service) UpdatePatient(patient Patient) (Patient, error) {
	return s.repository.UpdatePatient(patient)
}

func (s *Service) DeletePatient(id int) error {
	return s.repository.DeletePatient(id)
}

func (s *Service) PatchPatient(patient Patient) (Patient, error) {
	return s.repository.PatchPatient(patient)
}
