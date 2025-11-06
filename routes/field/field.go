package routes

import (
	"field-service/clients"
	"field-service/constants"
	"field-service/controllers"
	"field-service/middlewares"

	"github.com/gin-gonic/gin"
)

type FieldRoute struct {
	controller controllers.IControllerRegistry
	client     clients.IClientRegistry
	group      *gin.RouterGroup
}

type IFieldRoute interface {
	Run()
}

func NewFieldRoute(group *gin.RouterGroup, controller controllers.IControllerRegistry, client clients.IClientRegistry) IFieldRoute {
	return &FieldRoute{
		controller: controller,
		group:      group,
		client:     client,
	}
}

func (f *FieldRoute) Run() {
	group := f.group.Group("/field")
	group.GET("", middlewares.AuthenticateWithoutToken(), f.controller.GetField().GetAllWithoutPagination)
	group.GET("/:uuid", middlewares.AuthenticateWithoutToken(), f.controller.GetField().GetByUUID)
	group.Use(middlewares.Authenticate())
	group.GET("/pagination", middlewares.CheckRole([]string{
		constants.Admin,
		constants.Customer,
	}, f.client), f.controller.GetField().GetAllWithPagination)
	// group.GET("/pagination", f.controller.GetField().GetAllWithPagination)
	group.POST("/create", middlewares.CheckRole([]string{
		constants.Admin,
	}, f.client), f.controller.GetField().Create)
	// group.POST("/create", f.controller.GetField().Create)
	group.PUT("/update/:uuid", middlewares.CheckRole([]string{
		constants.Admin,
	}, f.client), f.controller.GetField().Update)
	// group.PUT("/update/:uuid", f.controller.GetField().Update)
	group.DELETE("/delete/:uuid", middlewares.CheckRole([]string{
		constants.Admin,
	}, f.client), f.controller.GetField().Update)
	// group.DELETE("/delete/:uuid", f.controller.GetField().Delete)
}
