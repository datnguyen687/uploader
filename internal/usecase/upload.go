package usecase

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/sirupsen/logrus"
)

func (uc *usecase) Upload(ctx context.Context, name string, data io.Reader) error {
	now := time.Now().UTC().UnixNano()

	newName := fmt.Sprintf("%s.%d", name, now)

	if err := uc.uploader.Upload(ctx, newName, data); err != nil {
		logrus.WithError(err).Errorf("unable to upload %s", name)
		return err
	}

	return nil
}
