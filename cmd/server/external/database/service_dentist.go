package database

import (
	d "Final_go/internal/clinic/dentists"
	"database/sql"
)

type SqlDentistStorage struct {
	*sql.DB
}

func NewDentistStorage(db *sql.DB) *SqlDentistStorage {
	return &SqlDentistStorage{db}
}

func (s *SqlDentistStorage) CreateDentist(dentist d.Dentist) (d.Dentist, error) {
	query := `INSERT INTO dentists (name, last_name, license) VALUES ($1, $2, $3) RETURNING id`
	err := s.QueryRow(query, dentist.Name, dentist.LastName, dentist.License).Scan(&dentist.ID)
	if err != nil {
		return d.Dentist{}, err
	}
	return dentist, nil
}

func (s *SqlDentistStorage) GetDentistByID(id int) (d.Dentist, error) {
	query := `SELECT * FROM dentists WHERE id = $1`
	var dentist d.Dentist
	err := s.QueryRow(query, id).Scan(&dentist.ID, &dentist.Name, &dentist.LastName, &dentist.License)
	if err != nil {
		return d.Dentist{}, err
	}
	return dentist, nil
}

func (s *SqlDentistStorage) UpdateDentist(dentist d.Dentist) (d.Dentist, error) {
	query := `UPDATE dentists SET name = $2, last_name = $3, license = $4 WHERE id = $1`
	_, err := s.Exec(query, dentist.ID, dentist.Name, dentist.LastName, dentist.License)
	if err != nil {
		return d.Dentist{}, err
	}
	return dentist, nil
}

func (s *SqlDentistStorage) PatchDentist(dentist d.Dentist) (d.Dentist, error) {
	query := `UPDATE dentists SET name = $2, last_name = $3, license = $4 WHERE id = $1 RETURNING id`
	err := s.QueryRow(query, dentist.ID, dentist.Name, dentist.LastName, dentist.License).Scan(&dentist.ID)
	if err != nil {
		return d.Dentist{}, err
	}
	return dentist, nil
}

func (s *SqlDentistStorage) DeleteDentist(id int) error {
	query := `DELETE FROM dentists WHERE id = $1`
	_, err := s.Exec(query, id)
	return err
}
