package appointments

type AppointmentRepository interface {
	CreateAppointment(appointment Appointment) (Appointment, error)
	GetAppointmentByID(id int) (Appointment, error)
	UpdateAppointment(appointment Appointment) (Appointment, error)
	DeleteAppointment(id int) error
	GetAppointmentByPatientDNI(dni string) ([]Appointment, error)
}

type Service struct {
	repository AppointmentRepository
}

func NewAppointmentService(repository AppointmentRepository) *Service {
	return &Service{repository: repository}
}

func (s *Service) CreateAppointment(appointment Appointment) (Appointment, error) {
	return s.repository.CreateAppointment(appointment)
}

func (s *Service) GetAppointmentByID(id int) (Appointment, error) {
	return s.repository.GetAppointmentByID(id)
}

func (s *Service) UpdateAppointment(appointment Appointment) (Appointment, error) {
	return s.repository.UpdateAppointment(appointment)
}

func (s *Service) DeleteAppointment(id int) error {
	return s.repository.DeleteAppointment(id)
}

func (s *Service) GetAppointmentByPatientDNI(dni string) ([]Appointment, error) {
	return s.repository.GetAppointmentByPatientDNI(dni)
}
