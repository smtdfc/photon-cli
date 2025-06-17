package photonGenerator

import (
	"fmt"
	"os"
	"path/filepath"
)

type AppTemplateData struct {
	AppName string
	Module  string
}

func GenerateApp(name string, path string) error {
	modName, err := ReadGoModModuleName(".")
	if err != nil {
		return err
	}

	data := AppTemplateData{
		AppName: name,
		Module:  modName,
	}

	files := map[string]string{
		"app/app.go":      "app/app.go.tmpl",
		"app/module.go":   "app/module.go.tmpl",
		"main.go":         "app/main.go.tmpl",
		"app.photon.json": "app/app.photon.json.tmpl",
	}

	for relPath, tmpl := range files {
		fullPath := filepath.Join(path, relPath)

		if err := os.MkdirAll(filepath.Dir(fullPath), os.ModePerm); err != nil {
			return fmt.Errorf("failed to create directories: %w", err)
		}

		if err := GenerateFile(fullPath, tmpl, data); err != nil {
			return fmt.Errorf("failed to generate file %s: %w", fullPath, err)
		}
	}

	return nil
}
