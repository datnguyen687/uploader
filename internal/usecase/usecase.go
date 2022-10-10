package usecase

import (
	"context"
	"io"
	"uploader/pkg/uploader"
)

func NewUsecase(uploader uploader.IUploader) *usecase {
	return &usecase{
		uploader: uploader,
	}
}

type IUsecase interface {
	Upload(ctx context.Context, name string, data io.Reader) error
}

type usecase struct {
	uploader uploader.IUploader
}
