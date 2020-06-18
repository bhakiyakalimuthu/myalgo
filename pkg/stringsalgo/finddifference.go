package stringsalgo

import (
	"encoding/csv"
	"errors"
	"os"
)

func Difference(a, b []string) []string {
	mb := make(map[string]string, len(b))
	for _, x := range b {
		mb[x] = x
	}
	var diff []string
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}

func Reader(path string, col int) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.New("failed to open file")
	}
	defer file.Close()
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, errors.New("failed to read file data")
	}
	recordTOReturn := make([]string, 0, len(records))
	for _, item := range records {
		recordTOReturn = append(recordTOReturn, item[col])
	}
	return recordTOReturn, nil

}
