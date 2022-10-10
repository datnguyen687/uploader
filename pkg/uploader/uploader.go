package uploader

import (
	"context"
	"io"
)

type IUploader interface {
	Upload(ctx context.Context, name string, file io.Reader) error
}
