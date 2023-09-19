package main

import (
	"Final_go/cmd/routers"
	"Final_go/cmd/server/config"
	"Final_go/cmd/server/external/database"
	"Final_go/cmd/server/handler/appointment"
	"Final_go/cmd/server/handler/dentist"
	"Final_go/cmd/server/handler/patient"
	"Final_go/cmd/server/middlewares"
	"Final_go/internal/clinic/appointments"
	"Final_go/internal/clinic/dentists"
	"Final_go/internal/clinic/patients"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load(".env")

	env := os.Getenv("ENV")
	if env == "" {
		env = "local"
	}

	if env == "local" {
		err := godotenv.Load()
		if err != nil {
			panic(err)
		}
	}

	cfg, err := config.NewConfig(env)

	if err != nil {
		panic(err)
	}

	router := routers.SetupRouter()

	// docs.SwaggerInfo.Host = os.Getenv("HOST")
	// router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	customRecovery := gin.CustomRecovery(middlewares.CustomRecoveryMiddleware)
	router.Use(customRecovery)

	authMidd := middlewares.NewAuthMiddleware(cfg.PublicConfig.PublicKey, cfg.PrivateConfig.SecretKey)
	router.Use(authMidd.AuthHeader)

	router.GET("/test-error", func(c *gin.Context) {
		err := fmt.Errorf("Este es un error de prueba")
		middlewares.CustomRecoveryMiddleware(c, err)
	})

	// Configurar base de datos
	postgresDatabase, err := database.NewPostgresSQLDatabase(cfg.PublicConfig.PostgresHost,
		cfg.PublicConfig.PostgresPort, cfg.PublicConfig.PostgresUser, cfg.PrivateConfig.PostgresPassword,
		cfg.PublicConfig.PostgresDBName)

	if err != nil {
		panic(err)
	}

	myDatabaseAppointments := database.NewSqlAppointmentStorage(postgresDatabase)
	myDatabaseDentists := database.NewDentistStorage(postgresDatabase)
	myDatabasePatients := database.NewPatientsStorage(postgresDatabase)

	appointmentsService := appointments.NewAppointmentService(myDatabaseAppointments)
	dentinstService := dentists.NewDentistService(myDatabaseDentists)
	patientService := patients.NewPatientService(myDatabasePatients)

	routers.SetupAppointmentRoutes(router, appointment.NewAppointmentHandler(appointmentsService), authMidd)
	routers.SetupDenstistRoutes(router, dentist.NewDentistHandler(dentinstService), authMidd)
	routers.SetupPatientRoutes(router, patient.NewPatientHandler(patientService), authMidd)

	err = router.Run()

	if err != nil {
		panic(err)
	}
}
