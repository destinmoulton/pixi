package player

import (
	"os/exec"
	"path"
	"strings"

	"../settings"
)

var filetypes = map[string]string{
	".avi":  "video",
	".mpeg": "video",
	".mkv":  "video",
	".mp4":  "video",
}

// PlayVideo starts the video player
func PlayVideo(selectedFilePath string) {
	cmdString := settings.Get(settings.SetConfig, settings.KeyOmxplayerCommand).(string)
	cmdParts := strings.Split(cmdString, " ")
	cmdParts = append(cmdParts, selectedFilePath)

	cmd := exec.Command(cmdParts[0], cmdParts[1:]...)
	err := cmd.Run()

	if err != nil {

	}
}

// IsVideoFile determines if a file is a video
func IsVideoFile(filename string) bool {
	if val, ok := filetypes[path.Ext(filename)]; ok == true && val == "video" {
		return true
	}
	return false
}
