package photonGenerator

import (
	"bufio"
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"os"
	"path/filepath"
	"strings"
)

func ReadGoModModuleName(basePath string) (string, error) {
	file, err := os.Open(basePath + "/go.mod")
	if err != nil {
		return "", fmt.Errorf("failed to open go.mod: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "module ") {
			return strings.TrimSpace(strings.TrimPrefix(line, "module")), nil
		}
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error reading go.mod: %w", err)
	}

	return "", fmt.Errorf("module line not found in go.mod")
}

func Capitalize(s string) string {
	caser := cases.Title(language.English)
	return caser.String(s)
}

func CreateFileIfNotExists(path string, content []byte) error {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return fmt.Errorf("cannot resolve absolute path: %w", err)
	}

	if _, err := os.Stat(absPath); err == nil {
		return nil
	} else if !os.IsNotExist(err) {
		return fmt.Errorf("error checking file existence: %w", err)
	}

	if err := os.MkdirAll(filepath.Dir(absPath), 0755); err != nil {
		return fmt.Errorf("failed to create parent directories: %w", err)
	}

	if err := os.WriteFile(absPath, content, 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}
