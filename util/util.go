package util

import "fmt"

// HumanReadableBytes prints the bytes into a linuxy human readable format
// Replicates the sizes shown via `ls -h`
// Inspired by: https://play.golang.org/p/68w_QCsE4F
func HumanReadableBytes(sizeB int64, suffix string) string {
	size := float64(sizeB)
	units := []string{"", "K", "M", "G", "T", "P", "E", "Z"}
	for _, unit := range units {
		if size < 1024.0 {
			if size < 10.0 {
				if size == 0.0 {
					return "0"
				}
				return fmt.Sprintf("%3.1f%s%s", size, unit, suffix)
			}
			return fmt.Sprintf("%3d%s%s", RoundUp(size), unit, suffix)
		}
		size = (size / 1024)
	}
	return fmt.Sprintf("%d%s%s", RoundUp(size), "Y", suffix)
}

// RoundUp rounds up to an int
func RoundUp(val float64) int {
	if val > 0 {
		return int(val + 1.0)
	}
	return int(val)
}
