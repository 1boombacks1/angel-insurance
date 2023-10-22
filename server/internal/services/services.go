package services

import (
	"context"
	"io"

	"github.com/1boombacks1/botInsurance/utils"
)

type Auth interface {
}

type Client interface {
	CreateNewCarApplication(ctx context.Context, carInsuranceApp CarInsuranceApplication) error
	// CreateNewHomeApplication()
	// CheckStatusForCurrentApplication()
	// EditCurrentApplication()
	// CheckHistoryApplications()
}

type Insurer interface {
	GetAllAvailableApplications()
}

type CarInsuranceApplicationBuilder interface {
	SetDescription(ctx context.Context, description string)
	// Провекра, что машина чистая
	// Проверка, что машина не разобрана или не в процессе ремонта
	SetVINPhoto(ctx context.Context, photo io.Reader) error
	// SetMainOutsidePhotos(ctx context.Context) error
	// SetWindshieldPhoto(ctx context.Context) error
	// SetMarkWindshieldPhoto(ctx context.Context) error
	// SetWheelPhoto(ctx context.Context) error
	// SetOdometerDataPhoto(ctx context.Context) error
	// SetCarInteriorPhotos(ctx context.Context) error
	// SetCarDamagePhotos(ctx context.Context) error
	// SetCarKeyPhotos(ctx context.Context) error

	// GetResultApplication()
}

type HomeInsuranceApplication interface {
}

type Services struct {
	CarInsuranceApplication CarInsuranceApplicationBuilder
}

func NewServices(photoHandler utils.PhotoHandler) *Services {
	return &Services{
		CarInsuranceApplication: NewCarInsuranceApplication(photoHandler),
	}
}
