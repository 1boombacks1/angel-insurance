package services

import (
	"context"
	"io"

	"github.com/1boombacks1/botInsurance/utils"
)

type Photo struct {
	TypePhoto string
	Link      string
}

type CarInsuranceApplication struct {
	PhotoHandler utils.PhotoHandler

	Description string
	Photos      []Photo
}

func NewCarInsuranceApplication(photoHandler utils.PhotoHandler) *CarInsuranceApplication {
	return &CarInsuranceApplication{
		PhotoHandler: photoHandler,
	}
}

func (s *CarInsuranceApplication) SetDescription(ctx context.Context, description string) {
	//TODO: setDesc
}

func (s *CarInsuranceApplication) SetVINPhoto(ctx context.Context, photo io.Reader) error {
	ok, err := s.PhotoHandler.IsCorrectResolution(photo)
	if err != nil {
		return err
	}
	if !ok {
		return ErrBadResolution
	}

	err = s.PhotoHandler.RegisterMetadata(photo)
	if err != nil {
		return err
	}

	// ...

	return nil
}
