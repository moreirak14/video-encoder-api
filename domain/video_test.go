package domain_test

import (
	"github.com/moreirak14/video-encoder-api/domain"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestValidateIfVideoIsEmpty(t *testing.T) {
	video := domain.NewVideo()
	err := video.Validate()

	require.Error(t, err)
}

func TestValidateIfVideoIDIsNotUUID(t *testing.T) {
	video := domain.NewVideo()
	video.ID = "xpto"
	err := video.Validate()

	require.Error(t, err)
}

func TestValidateIfVideoResourceIDIsEmpty(t *testing.T) {
	video := domain.NewVideo()
	video.ResourceID = ""
	err := video.Validate()

	require.Error(t, err)
}

func TestValidateIfVideoFilePathIsEmpty(t *testing.T) {
	video := domain.NewVideo()
	video.FilePath = ""
	err := video.Validate()

	require.Error(t, err)
}

func TestValidateIfVideoHasAllFields(t *testing.T) {
	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.ResourceID = "xpto"
	video.FilePath = "path"
	video.CreatedAt = time.Now()
	err := video.Validate()

	require.Nil(t, err)
}
