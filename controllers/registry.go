package controllers

import (
	fieldControllers "field-service/controllers/field"
	fieldSchedulecontrollers "field-service/controllers/fieldSchedule"
	timeControllers "field-service/controllers/time"
	"field-service/services"
)

type Registry struct {
	service services.IServiceRegistry
}

type IControllerRegistry interface {
	GetField() fieldControllers.IFieldController
	GetFieldSchedule() fieldSchedulecontrollers.IFieldScheduleController
	GetTime() timeControllers.ITimeController
}

func NewControllerRegistry(service services.IServiceRegistry) IControllerRegistry {
	return &Registry{service: service}
}

// GetFieldSchedule implements IControllerRegistry.
func (r *Registry) GetFieldSchedule() fieldSchedulecontrollers.IFieldScheduleController {
	return fieldSchedulecontrollers.NewFieldScheduleController(r.service)
}

// GetTime implements IControllerRegistry.
func (r *Registry) GetTime() timeControllers.ITimeController {
	return timeControllers.NewTimeController(r.service)
}

// etField implements IControllerRegistry.
func (r *Registry) GetField() fieldControllers.IFieldController {
	return fieldControllers.NewFieldController(r.service)
}
