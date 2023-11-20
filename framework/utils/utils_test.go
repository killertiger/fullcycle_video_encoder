package utils_test

import (
	"testing"

	"github.com/killertiger/fullcycle_video_encoder/framework/utils"
	"github.com/stretchr/testify/require"
)

func TestIsJson(t *testing.T) {
	json := `{
		"id": "de30dca9-d736-45a6-bd88-4a4abfc18828",
		"file_path": "convite.mp4",
		"status": "pending"
	}`

	err := utils.IsJson(json)
	require.Nil(t, err)

	json = `fake json`
	err = utils.IsJson(json)
	require.Error(t, err)
}
