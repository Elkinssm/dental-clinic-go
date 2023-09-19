package database

import (
	a "Final_go/internal/clinic/appointments"
	"database/sql"
	"errors"
	"fmt"
)

type SqlAppointmentStorage struct {
	*sql.DB
}

func NewSqlAppointmentStorage(db *sql.DB) *SqlAppointmentStorage {
	return &SqlAppointmentStorage{db}
}

func (s *SqlAppointmentStorage) CreateAppointment(appointment a.Appointment) (a.Appointment, error) {
	query := `
		INSERT INTO appointments (date, hour, description, patient_id, dentist_id) 
		VALUES ($1, $2, $3, $4, $5) 
		RETURNING id, patient_id, dentist_id
	`
	err := s.QueryRow(query, appointment.Date, appointment.Hour, appointment.Description, appointment.Patient.ID, appointment.Dentist.ID).
		Scan(&appointment.ID, &appointment.Patient.ID, &appointment.Dentist.ID)
	if err != nil {
		return a.Appointment{}, err
	}
	return appointment, nil
}

func (s *SqlAppointmentStorage) GetAppointmentByID(id int) (a.Appointment, error) {
	query := `SELECT a.id, a.date, a.hour, a.description, 
                     p.id as patient_id, p.name as patient_name, p.last_name as patient_last_name, p.address as patient_address, p.dni as patient_dni, p.registration_date as patient_registration_date,
                     d.id as dentist_id, d.name as dentist_name, d.last_name as dentist_last_name, d.license as dentist_license
              FROM appointments a
              JOIN patients p ON a.patient_id = p.id
              JOIN dentists d ON a.dentist_id = d.id
              WHERE a.id = $1`
	var appointment a.Appointment
	err := s.QueryRow(query, id).Scan(&appointment.ID, &appointment.Date, &appointment.Hour, &appointment.Description,
		&appointment.Patient.ID, &appointment.Patient.Name, &appointment.Patient.LastName, &appointment.Patient.Address, &appointment.Patient.DNI, &appointment.Patient.RegistrationDate,
		&appointment.Dentist.ID, &appointment.Dentist.Name, &appointment.Dentist.LastName, &appointment.Dentist.License)
	if err != nil {
		return a.Appointment{}, err
	}
	return appointment, nil
}

func (s *SqlAppointmentStorage) UpdateAppointment(appointment a.Appointment) (a.Appointment, error) {
	query := `UPDATE appointments SET date=$1, hour=$2, description=$3, patient_id=$4, dentist_id=$5 WHERE id=$6`
	_, err := s.Exec(query, appointment.Date, appointment.Hour, appointment.Description, appointment.Patient.ID, appointment.Dentist.ID, appointment.ID)
	if err != nil {
		return a.Appointment{}, err
	}

	return appointment, nil
}

func (s *SqlAppointmentStorage) PatchAppointment(appointment a.Appointment) (a.Appointment, error) {
	query := `UPDATE appointments SET date=$1, hour=$2, description=$3 WHERE id=$4`
	_, err := s.Exec(query, appointment.Date, appointment.Hour, appointment.Description, appointment.ID)
	if err != nil {
		return a.Appointment{}, err
	}

	return appointment, nil
}

func (s *SqlAppointmentStorage) DeleteAppointment(id int) error {
	query := `DELETE FROM appointments WHERE id = $1`
	_, err := s.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *SqlAppointmentStorage) GetAppointmentByPatientDNI(dni string) ([]a.Appointment, error) {

	query := `SELECT a.id, a.date, a.hour, a.description, 
                     p.id as patient_id, p.name as patient_name, p.last_name as patient_last_name, p.address as patient_address, p.dni as patient_dni, p.registration_date as patient_registration_date,
                     d.id as dentist_id, d.name as dentist_name, d.last_name as dentist_last_name, d.license as dentist_license
              FROM appointments a
              JOIN patients p ON a.patient_id = p.id
              JOIN dentists d ON a.dentist_id = d.id
              WHERE p.dni = $1`

	rows, err := s.Query(query, dni)
	if err != nil {
		fmt.Println("Error al ejecutar el query:", err.Error())
		return nil, err
	}
	defer rows.Close()

	var appointments []a.Appointment
	for rows.Next() {
		var appointment a.Appointment
		err := rows.Scan(
			&appointment.ID, &appointment.Date, &appointment.Hour, &appointment.Description,
			&appointment.Patient.ID, &appointment.Patient.Name, &appointment.Patient.LastName, &appointment.Patient.Address, &appointment.Patient.DNI, &appointment.Patient.RegistrationDate,
			&appointment.Dentist.ID, &appointment.Dentist.Name, &appointment.Dentist.LastName, &appointment.Dentist.License,
		)
		if err != nil {
			fmt.Println("Error al escanear las filas:", err.Error())
			return nil, err
		}
		appointments = append(appointments, appointment)
	}

	// Si no se encontraron turnos, devolvemos un error específico
	if len(appointments) == 0 {
		return nil, errors.New("No se encontraron turnos para el DNI proporcionado")
	}

	return appointments, nil
}
