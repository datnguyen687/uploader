package usecase_test

import (
	"bytes"
	"context"
	"errors"
	"testing"
	"uploader/internal/usecase"
	"uploader/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUploadOk(t *testing.T) {
	name := "fake.json"
	data := bytes.NewBuffer([]byte("abcdefgh"))

	uploader := mocks.IUploader{}
	uploader.On("Upload", mock.Anything, mock.AnythingOfType("string"), data).Return(nil)

	uc := usecase.NewUsecase(&uploader)

	ctx := context.Background()
	err := uc.Upload(ctx, name, data)

	assert.NoError(t, err)
}

func TestUploadErr(t *testing.T) {
	name := "fake.json"
	data := bytes.NewBuffer([]byte("abcdefgh"))
	expectedErr := errors.New("failed to upload")

	uploader := mocks.IUploader{}
	uploader.On("Upload", mock.Anything, mock.AnythingOfType("string"), data).Return(expectedErr)

	uc := usecase.NewUsecase(&uploader)

	ctx := context.Background()
	err := uc.Upload(ctx, name, data)

	assert.ErrorIs(t, err, expectedErr)
}
