package history

import (
	"path"

	"../../settings"
)

type viewedFile map[string]string

type viewedHistory []viewedFile

var history viewedHistory

func initHistory() {
	opened := settings.Get(settings.SetHistory, "opened")

	for _, file := range opened.([]interface{}) {
		// Convert the returned interface (from JSON) into usable map
		tmp := make(viewedFile)
		tmp["filename"] = file.(map[string]interface{})["filename"].(string)
		tmp["path"] = file.(map[string]interface{})["path"].(string)
		history = append(history, tmp)
	}
}

func (h *viewedHistory) Add(fullPath string) {
	_, filename := path.Split(fullPath)

	file := make(viewedFile)
	file["filename"] = filename
	file["path"] = fullPath

	// unshift the new element onto the front of the history
	*h = append(viewedHistory{file}, *h...)

	settings.Set(settings.SetHistory, "opened", h)
}
