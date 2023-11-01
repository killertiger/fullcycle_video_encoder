package domain_test

import (
	"testing"
	"time"

	"github.com/killertiger/fullcycle_video_encoder/domain"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestValidateIfVideoIsEmpty(t *testing.T) {
	video := domain.NewVideo()
	err := video.Validate()
	require.Error(t, err)
}

func TestVideoIdIsNotAUuid(t *testing.T) {
	video := domain.NewVideo()

	video.ID = uuid.NewV4().String()
	video.ResourceID = "def"
	video.FilePath = "any_path"
	video.CreatedAt = time.Now()

	err := video.Validate()
	require.Nil(t, err)
}
