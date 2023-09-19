# Sistema de Reserva de Turnos - API

Este proyecto implementa una API para administrar la reserva de turnos en una clínica odontológica. Cumple con varios requisitos funcionales y técnicos detallados a continuación.
Los cuales fueron desarrollados para la materia de backend 3 de la especializacion de Digital House en backend

## Requerimientos Funcionales

### Administración de Odontólogos
- **/api/v1/dentists** : Url inicial Api
- **POST /**: Agregar un dentista.
- **GET /:id**: Traer información de un dentista por ID.
- **PUT /:id**: Actualizar información de un dentista.
- **PATCH /:id**: Actualizar un campo específico de un dentista.
- **DELETE /:id**: Eliminar un dentista.

### Administración de Pacientes
- **/api/v1/patients** : Url inicial Api
- **POST /**: Agregar un paciente.
- **GET /:id**: Traer información de un paciente por ID.
- **PUT /:id**: Actualizar información de un paciente.
- **PATCH /:id**: Actualizar un campo específico de un paciente.
- **DELETE /:id**: Eliminar un paciente.

### Registrar Turnos
- **/api/v1/appointments** : Url inicial Api
- **POST /**: Agregar un turno.
- **GET /:id**: Traer información de un turno por ID.
- **PUT /:id**: Actualizar información de un turno.
- **PATCH /:id**: Actualizar un campo específico de un turno.
- **DELETE /:id**: Eliminar un turno.
- **GET /by-patient-dni**: Traer detalles de los turnos de un paciente por su DNI.

### Seguridad y Middleware
- Se implementa un middleware para proveer autenticación en los métodos POST, PUT, PATCH y DELETE.


## Requerimientos Técnicos

- El proyecto sigue una estructura orientada a paquetes con las siguientes capas:
  - Capa de Entidades de Negocio.
  - Capa de Repositorio (Acceso a Datos).
  - Capa de Acceso a Datos (Base de Datos).
  - Capa de Servicios.
  - Capa de Controladores (Handlers).

- Se utiliza una base de datos relacional PostgreSQL.

### Scheme
![image](https://github.com/Elkinssm/dental-clinic-go/assets/52393397/2276e658-b077-4385-8799-258fac66f60e)


## Instalación y Uso

1. Clona el repositorio.
2. Configura la base de datos en el archivo de configuración correspondiente usar datos de docker-compose.yml
3. Instala las dependencias.
4. Ejecuta la aplicación.

## Coleccion de postman 
[Colección de Postman](./Dental_clinic_API_Go.postman_collection.json)
[Variables de entorno Postman](./dental_clinic_api.postman_environment.json)

## Licencia

Este proyecto está bajo la licencia [MT]. 

## Contacto

Elkin Silva
[<img src="https://cdn-icons-png.flaticon.com/512/174/174857.png" width="50" height="50">](https://www.linkedin.com/in/elkinssm/)

Federico Bonesi
[<img src="https://cdn-icons-png.flaticon.com/512/174/174857.png" width="50" height="50">](https://www.linkedin.com/in/federico-guillermo-bonesi-ale-307591186/)

María Gabriela Mateo
[<img src="https://cdn-icons-png.flaticon.com/512/174/174857.png" width="50" height="50">](https://www.linkedin.com/in/mariagabrielamateo/)

