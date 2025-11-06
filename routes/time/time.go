package routes

import (
	"field-service/clients"
	"field-service/constants"
	"field-service/controllers"
	"field-service/middlewares"

	"github.com/gin-gonic/gin"
)

type TimeRoute struct {
	controller controllers.IControllerRegistry
	group      *gin.RouterGroup
	client     clients.IClientRegistry
}

type ITimeRoute interface {
	Run()
}

func NewTimeRoute(group *gin.RouterGroup, controller controllers.IControllerRegistry, client clients.IClientRegistry) ITimeRoute {
	return &TimeRoute{
		controller: controller,
		group:      group,
		client:     client,
	}
}

func (f *TimeRoute) Run() {
	group := f.group.Group("/time")
	group.Use(middlewares.Authenticate())
	group.GET("", middlewares.CheckRole([]string{
		constants.Admin,
	}, f.client), f.controller.GetTime().GetAll)
	// group.GET("", f.controller.GetTime().GetAll)
	group.GET("/:uuid", middlewares.CheckRole([]string{
		constants.Admin,
	}, f.client), f.controller.GetTime().GetByUUID)
	// group.GET("/:uuid", f.controller.GetTime().GetByUUID)
	group.POST("/create", middlewares.CheckRole([]string{
		constants.Admin,
	}, f.client), f.controller.GetTime().Create)
	// group.POST("/create", f.controller.GetTime().Create)
}
