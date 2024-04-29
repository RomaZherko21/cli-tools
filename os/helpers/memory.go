package helpers

import (
	"fmt"
	"strconv"
)

const KB = 1024
const MB = 1024 * KB
const GB = 1024 * MB

func ConvertToGB(bytes uint64) float64 {
	return float64(bytes) / GB
}

func ConvertToMB(bytes uint64) float64 {
	return float64(bytes) / MB
}

func ConvertToKB(bytes uint64) float64 {
	return float64(bytes) / MB
}

func ReadableMemory(bytes uint64) string {
	gb := ConvertToGB(bytes)
	if gb < 1 {
		mb := ConvertToMB(bytes)
		if mb < 1 {
			kb := ConvertToKB(bytes)
			return fmt.Sprintf("%v KB", strconv.Itoa(int(kb)))
		}
		return fmt.Sprintf("%v MB", strconv.Itoa(int(mb)))
	}
	return fmt.Sprintf("%v GB", strconv.Itoa(int(gb)))
}
