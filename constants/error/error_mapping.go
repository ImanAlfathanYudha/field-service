package error

import (
	errField "field-service/constants/error/field"
	errFieldSchedule "field-service/constants/error/fieldSchedule"
	errTime "field-service/constants/error/time"
)

func ErrMapping(err error) bool {
	allErrors := make([]error, 8)
	// allErrors = append(append(append(GeneralErrors[:], errField.FieldErrors[:]...), errFieldSchedule.FieldScheduleErrors[:]...), errTime.TimeErrors[:]...)
	allErrors = append(allErrors, GeneralErrors...)
	allErrors = append(allErrors, errField.FieldErrors...)
	allErrors = append(allErrors, errFieldSchedule.FieldScheduleErrors...)
	allErrors = append(allErrors, errTime.TimeErrors...)

	for _, item := range allErrors {
		if err.Error() == item.Error() {
			return true

		}
	}
	return false
}
