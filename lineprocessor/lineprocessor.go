package lineprocessor

import (
	"bufio"
	"os"
)

type LineProcessFunc func(line string) error

func ReadInput(inputPath string, fn LineProcessFunc) error {
	path := inputPath
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineErr := fn(line)
		if lineErr != nil {
			return lineErr
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
