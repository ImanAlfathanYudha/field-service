package repositories

import (
	"context"
	"errors"
	errWrap "field-service/common/error"
	errConstant "field-service/constants/error"
	errTime "field-service/constants/error/time"
	"field-service/domain/models"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TimeRepository struct {
	db *gorm.DB
}

type ITimeRepository interface {
	FindAll(context.Context) ([]models.Time, error)
	FindByUUID(context.Context, string) (*models.Time, error)
	FindByID(context.Context, int) (*models.Time, error)
	Create(context.Context, *models.Time) (*models.Time, error)
}

func NewTimeRepository(db *gorm.DB) ITimeRepository {
	return &TimeRepository{db: db}
}

func (t *TimeRepository) FindAll(ctx context.Context) ([]models.Time, error) {
	var times []models.Time
	err := t.db.WithContext(ctx).Find(&times).Error
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errWrap.WrapError(errConstant.ErrSQLError), err)
	}
	return times, nil
}

func (t *TimeRepository) FindByUUID(ctx context.Context, uuid string) (*models.Time, error) {
	var times models.Time
	err := t.db.WithContext(ctx).Where("uuid = ?", uuid).First(&times).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("%w: %v", errWrap.WrapError(errTime.ErrTimeNotFound), err)
		}
		return nil, fmt.Errorf("%w: %v", errWrap.WrapError(errConstant.ErrSQLError), err)
	}
	return &times, nil
}

func (t *TimeRepository) FindByID(ctx context.Context, id int) (*models.Time, error) {
	var time models.Time
	err := t.db.WithContext(ctx).Where("id = ?", id).First(&time).Error
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errWrap.WrapError(errConstant.ErrSQLError), err)
	}
	return &time, nil
}

func (t *TimeRepository) Create(ctx context.Context, req *models.Time) (*models.Time, error) {
	req.UUID = uuid.New()
	err := t.db.WithContext(ctx).Create(&req).Error
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errWrap.WrapError(errConstant.ErrSQLError), err)
	}
	return req, nil
}
