package routes

import (
	"field-service/clients"
	"field-service/constants"
	"field-service/controllers"
	"field-service/middlewares"

	"github.com/gin-gonic/gin"
)

type FieldScheduleRoute struct {
	controller controllers.IControllerRegistry
	client     clients.IClientRegistry
	group      *gin.RouterGroup
}

type IFieldScheduleRoute interface {
	Run()
}

func NewFieldScheduleRoute(group *gin.RouterGroup, controller controllers.IControllerRegistry, client clients.IClientRegistry) IFieldScheduleRoute {
	return &FieldScheduleRoute{
		controller: controller,
		group:      group,
		client:     client,
	}
}

func (f *FieldScheduleRoute) Run() {
	group := f.group.Group("/field/schedule")
	group.PATCH("/update-status", middlewares.AuthenticateWithoutToken(), f.controller.GetFieldSchedule().UpdateStatus)
	group.Use(middlewares.Authenticate())
	group.GET("/pagination", middlewares.CheckRole([]string{
		constants.Admin,
		constants.Customer,
	}, f.client), f.controller.GetFieldSchedule().GetAllWithPagination)
	// group.GET("/pagination", f.controller.GetFieldSchedule().GetAllWithPagination)
	group.GET("/:uuid", middlewares.CheckRole([]string{
		constants.Admin,
		constants.Customer,
	}, f.client), f.controller.GetFieldSchedule().GetByUUID)
	// group.GET("/:uuid", f.controller.GetFieldSchedule().GetByUUID)
	group.POST("/generate-one-month", middlewares.CheckRole([]string{
		constants.Admin,
	}, f.client), f.controller.GetFieldSchedule().GenerateScheduleForOneMonth)
	// group.POST("/generate-one-month", f.controller.GetFieldSchedule().GenerateScheduleForOneMonth)
	group.POST("/create", middlewares.CheckRole([]string{
		constants.Admin,
	}, f.client), f.controller.GetFieldSchedule().Create)
	// group.POST("/create", f.controller.GetFieldSchedule().Create)
	group.PUT("/update", middlewares.CheckRole([]string{
		constants.Admin,
	}, f.client), f.controller.GetFieldSchedule().Update)
	// group.PUT("/update/:uuid", f.controller.GetFieldSchedule().Update)
	group.DELETE("/delete/:uuid", middlewares.CheckRole([]string{
		constants.Admin,
	}, f.client), f.controller.GetFieldSchedule().Delete)
	// group.DELETE("/delete/:uuid", f.controller.GetFieldSchedule().Delete)
}
