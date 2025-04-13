package fileHandling

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func WriteFloatToFile(filename string, value float64) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("%f\n", value))
	return err
}

func ReadFloatFromFile(filename string) (float64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		return strconv.ParseFloat(scanner.Text(), 64)
	}
	return 0, scanner.Err()
}
