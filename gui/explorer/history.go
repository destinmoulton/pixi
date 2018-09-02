package explorer

import "path"

type tViewedFile struct {
	filename string
	path     string
}

type viewHistory []tViewedFile

func (h *viewHistory) add(fullPath string) {
	_, filename := path.Split(fullPath)
	*h = append(*h, tViewedFile{filename, fullPath})
}
