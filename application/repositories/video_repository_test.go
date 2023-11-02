package repositories_test

import (
	"testing"
	"time"

	"github.com/killertiger/fullcycle_video_encoder/application/repositories"
	"github.com/killertiger/fullcycle_video_encoder/domain"
	"github.com/killertiger/fullcycle_video_encoder/framework/database"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestVideoRepositoryDbInsert(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	repo := repositories.VideoRepositoryDb{Db: db}
	repo.Insert(video)

	v, err := repo.Find(video.ID)

	require.NotEmpty(t, v.ID)
	require.Nil(t, err)
	require.Equal(t, video.ID, v.ID)
}
