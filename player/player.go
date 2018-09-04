package player

import (
	"os/exec"
	"path"
)

var filetypes = map[string]string{
	".avi":  "video",
	".mpeg": "video",
	".mkv":  "video",
	".mp4":  "video",
}

// PlayVideo starts the video player
func PlayVideo(selectedFilePath string) {
	cmd := exec.Command("xterm", "-e", "omxplayer", "-b", selectedFilePath)
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
