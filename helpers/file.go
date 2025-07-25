package helpers

import (
	"bufio"
	"os"
)

func WriteFile(fileName string, content string) error {
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer f.Close()

	writer := bufio.NewWriter(f)
	_, err = writer.WriteString(content)
	if err != nil {
		return err
	}

	return writer.Flush()
}
