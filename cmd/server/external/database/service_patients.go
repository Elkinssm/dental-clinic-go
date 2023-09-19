package database

import (
	p "Final_go/internal/clinic/patients"
	"database/sql"
)

type SqlPatientsStorage struct {
	*sql.DB
}

func NewPatientsStorage(db *sql.DB) *SqlPatientsStorage {
	return &SqlPatientsStorage{db}
}

func (s *SqlPatientsStorage) CreatePatient(patient p.Patient) (p.Patient, error) {
	query := `INSERT INTO patients (name, last_name, address,dni,registration_date) VALUES ($1, $2, $3,$4,$5) RETURNING id`
	err := s.QueryRow(query, patient.Name, patient.LastName, patient.Address, patient.DNI, patient.RegistrationDate).Scan(&patient.ID)
	if err != nil {
		return p.Patient{}, err
	}
	return patient, nil
}

func (s *SqlPatientsStorage) GetPatientByID(id int) (p.Patient, error) {
	query := `SELECT * FROM patients WHERE id = $1`
	var patient p.Patient
	err := s.QueryRow(query, id).Scan(&patient.ID, &patient.Name, &patient.LastName, &patient.Address, &patient.DNI, &patient.RegistrationDate)
	if err != nil {
		return p.Patient{}, err
	}
	return patient, nil
}

func (s *SqlPatientsStorage) UpdatePatient(patient p.Patient) (p.Patient, error) {
	query := `UPDATE patients SET name = $2, last_name = $3, address = $4, dni = $5, registration_date = $6 WHERE id = $1`
	_, err := s.Exec(query, patient.ID, patient.Name, patient.LastName, patient.Address, patient.DNI, patient.RegistrationDate)
	if err != nil {
		return p.Patient{}, err
	}
	return patient, nil
}

func (s *SqlPatientsStorage) PatchPatient(patient p.Patient) (p.Patient, error) {
	query := `UPDATE patients SET name = $2, last_name = $3, address = $4, dni = $5, registration_date = $6 WHERE id = $1 RETURNING id`
	err := s.QueryRow(query, patient.ID, patient.Name, patient.LastName, patient.Address, patient.DNI, patient.RegistrationDate).Scan(&patient.ID)
	if err != nil {
		return p.Patient{}, err
	}
	return patient, nil
}

func (s *SqlPatientsStorage) DeletePatient(id int) error {
	query := `DELETE FROM patients WHERE id = $1`
	_, err := s.Exec(query, id)
	return err
}
