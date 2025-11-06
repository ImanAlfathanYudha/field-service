package error

import "errors"

var (
	ErrFieldScheduleNotFound = errors.New("Field schedule not found")
	ErrFieldScheduleIsExist  = errors.New("Field schedule already exists")
)

var FieldScheduleErrors = []error{
	ErrFieldScheduleNotFound, ErrFieldScheduleIsExist,
}
