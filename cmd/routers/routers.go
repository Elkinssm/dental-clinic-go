package routers

import (
	"Final_go/cmd/server/handler/appointment"
	"Final_go/cmd/server/handler/dentist"
	"Final_go/cmd/server/handler/patient"
	"Final_go/cmd/server/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter() *gin.Engine {
	router := gin.New()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"ok": "ok"})
	})

	return router
}

func SetupDenstistRoutes(router *gin.Engine, dentistHandler *dentist.DentistHandler, authMiddleware *middlewares.AuthMiddleware) {
	dentistsGroup := router.Group("/api/v1/dentists")

	dentistsGroup.POST("/", authMiddleware.AuthHeader, dentistHandler.CreateDentist)
	dentistsGroup.GET("/:id", dentistHandler.GetDentistByID)
	dentistsGroup.PUT("/:id", authMiddleware.AuthHeader, dentistHandler.UpdateDentist)
	dentistsGroup.DELETE("/:id", authMiddleware.AuthHeader, dentistHandler.DeleteDentist)
	dentistsGroup.PATCH("/:id", authMiddleware.AuthHeader, dentistHandler.PatchDentist)

}

func SetupPatientRoutes(router *gin.Engine, patientHandler *patient.PatientHandler, authMiddleware *middlewares.AuthMiddleware) {
	patientsGroup := router.Group("/api/v1/patients")

	patientsGroup.POST("/", authMiddleware.AuthHeader, patientHandler.CreatePatient)
	patientsGroup.GET("/:id", patientHandler.GetPatientByID)
	patientsGroup.PUT("/:id", authMiddleware.AuthHeader, patientHandler.UpdatePatient)
	patientsGroup.DELETE("/:id", authMiddleware.AuthHeader, patientHandler.DeletePatient)
	patientsGroup.PATCH("/:id", authMiddleware.AuthHeader, patientHandler.PatchPatient)
}

func SetupAppointmentRoutes(router *gin.Engine, appointmentHandler *appointment.AppointmentHandler, authMiddleware *middlewares.AuthMiddleware) {
	appointmentsGroup := router.Group("/api/v1/appointments")

	appointmentsGroup.GET("/:id", appointmentHandler.GetAppointmentById)
	appointmentsGroup.PUT("/:id", authMiddleware.AuthHeader, appointmentHandler.UpdateAppointment)
	appointmentsGroup.DELETE("/:id", authMiddleware.AuthHeader, appointmentHandler.DeleteAppointment)
	appointmentsGroup.POST("/", authMiddleware.AuthHeader, appointmentHandler.CreateAppointment)
	appointmentsGroup.PATCH("/:id", authMiddleware.AuthHeader, appointmentHandler.PatchAppointment)
	appointmentsGroup.GET("/by-patient-dni", appointmentHandler.GetAppointmentByPatientDNI)
}
