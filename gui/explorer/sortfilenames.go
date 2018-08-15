package explorer

import (
	"os"
	"strings"
)

// SortByLowerCaseFilename sorts a slice of file info by lower case
type SortByLowerCaseFilename []os.FileInfo

func (f SortByLowerCaseFilename) Len() int {
	return len(f)
}

func (f SortByLowerCaseFilename) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

func (f SortByLowerCaseFilename) Less(i, j int) bool {

	pathA := strings.ToLower(f[i].Name())
	pathB := strings.ToLower(f[j].Name())

	return pathA < pathB
}
