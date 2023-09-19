-- Comprobar si existe la base de datos
DROP TABLE IF EXISTS patients;

-- Crear la tabla de pacientes
CREATE TABLE patients (
                          id SERIAL PRIMARY KEY,
                          name VARCHAR(128) NOT NULL,
                          last_name VARCHAR(128) NOT NULL,
                          address VARCHAR(256),
                          dni VARCHAR(8) UNIQUE NOT NULL,
                          registration_date DATE NOT NULL
);

-- Comprobar si existe la base de datos
DROP TABLE IF EXISTS dentists;

-- Crear la tabla de dentistas
CREATE TABLE dentists (
                          id SERIAL PRIMARY KEY,
                          name VARCHAR(128) NOT NULL,
                          last_name VARCHAR(128) NOT NULL,
                          license VARCHAR(16) UNIQUE NOT NULL
);

-- Comprobar si existe la base de datos
DROP TABLE IF EXISTS appointments;

-- Crear la tabla de turnos
CREATE TABLE appointments (
                              id SERIAL PRIMARY KEY,
                              date VARCHAR(10) NOT NULL,
                              hour VARCHAR(10) NOT NULL,
                              description VARCHAR(256),
                              patient_id INTEGER NOT NULL,
                              dentist_id INTEGER NOT NULL,
                              FOREIGN KEY (patient_id) REFERENCES patients(id),
                              FOREIGN KEY (dentist_id) REFERENCES dentists(id)
)