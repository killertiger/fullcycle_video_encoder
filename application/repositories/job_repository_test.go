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

func TestJobRepositoryDbInsert(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	repoVideo := repositories.VideoRepositoryDb{Db: db}
	repoVideo.Insert(video)

	job, err := domain.NewJob("output_path", "Pending", video)
	require.Nil(t, err)

	repoJob := repositories.JobRepositoryDb{Db: db}
	repoJob.Insert(job)

	jobFound, err := repoJob.Find(job.ID)
	require.NotEmpty(t, jobFound.ID)
	require.Nil(t, err)
	require.Equal(t, job.ID, jobFound.ID)
	require.Equal(t, job.VideoID, jobFound.VideoID)
}

func TestJobRepositoryDbUpdate(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	repoVideo := repositories.VideoRepositoryDb{Db: db}
	repoVideo.Insert(video)

	job, err := domain.NewJob("output_path", "Pending", video)
	require.Nil(t, err)

	repoJob := repositories.JobRepositoryDb{Db: db}
	repoJob.Insert(job)

	job.Status = "Complete"
	repoJob.Update(job)

	jobFound, err := repoJob.Find(job.ID)
	require.NotEmpty(t, jobFound.ID)
	require.Nil(t, err)
	require.Equal(t, job.ID, jobFound.ID)
	require.Equal(t, job.Status, jobFound.Status)
}
