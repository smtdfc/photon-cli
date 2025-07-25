package helpers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetModuleName(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("Failed to open go.mod file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "module ") {
			return strings.TrimSpace(strings.TrimPrefix(line, "module ")), nil
		}
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("Error while reading go.mod file: %w", err)
	}

	return "", fmt.Errorf("Module declaration not found in go.mod")
}
