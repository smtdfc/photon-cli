package photonGenerator

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

type AppModuleDeclaration struct {
	Name        string `json:"name"`
	PackageName string `json:"package"`
	Path        string `json:"path"`
}

type AppDeclaration struct {
	Name    string                 `json:"name"`
	Version string                 `json:"version"`
	Project string                 `json:"project"`
	Modules []AppModuleDeclaration `json:"modules"`
}

func IsDeclarationFileExist(basePath string) bool {
	configPath := filepath.Join(basePath, "app.photon.json")

	info, err := os.Stat(configPath)
	if err != nil {
		return false
	}

	return !info.IsDir()
}

func ReadDeclarationFile(basePath string) (*AppDeclaration, error) {
	configPath := filepath.Join(basePath, "app.photon.json")

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read declaration file: %w", err)
	}

	var app AppDeclaration
	if err := json.Unmarshal(data, &app); err != nil {
		return nil, fmt.Errorf("invalid JSON in declaration file: %w", err)
	}

	if app.Name == "" || app.Version == "" || app.Project == "" {
		return nil, errors.New("missing required fields in declaration")
	}

	return &app, nil
}

func WriteDeclarationFile(basePath string, app *AppDeclaration) error {
	configPath := filepath.Join(basePath, "app.photon.json")

	data, err := json.MarshalIndent(app, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal declaration to JSON: %w", err)
	}

	if err := os.WriteFile(configPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write declaration file: %w", err)
	}

	return nil
}
