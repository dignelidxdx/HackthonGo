package util

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func ConvertToCsv(fileName string) error {

	path := fmt.Sprintf("../../datos/txt/%v.txt", fileName)
	outPutPath := fmt.Sprintf("../../datos/csv/%v.csv", fileName)

	file, err := os.ReadFile(path)

	if err != nil {

		return err
	}

	data := string(file)
	newData := strings.ReplaceAll(data, "#$%#", ",")

	os.WriteFile(outPutPath, []byte(newData), 0644)
	return nil
}

func ReadCsv(fileName string) ([][]string, error) {

	path := fmt.Sprintf("../../datos/csv/%v.csv", fileName)
	// Open CSV file
	f, err := os.Open(path)
	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()

	// Read File into a Variable
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return lines, nil
}
