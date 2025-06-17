package photonGenerator

import (
	"fmt"
	"path/filepath"
	"strings"
)

func GenerateModule(name, path string) error {
	if !IsDeclarationFileExist(path) {
		return fmt.Errorf("Cannot find declaration file!")
	}

	modulePkgName := strings.ToLower(name)
	moduleName := Capitalize(name)
	modulePath := filepath.Join(path, "internal", modulePkgName)

	declaration, _ := ReadDeclarationFile(path)
	declaration.Modules = append(declaration.Modules, AppModuleDeclaration{
		Name:        moduleName,
		PackageName: modulePkgName,
		Path:        "./internal/" + modulePkgName,
	})

	files := []string{
		"module.go",
		"init.go",
		"repository.go",
		"service.go",
		"routes.go",
	}

	for _, file := range files {
		err := CreateFileIfNotExists(filepath.Join(modulePath, file), []byte{})
		if err != nil {
			return fmt.Errorf("failed to create file %s: %w", file, err)
		}
	}

	templateFiles := map[string]string{
		"model.go":      "model.tmpl",
		"handler.go":    "handler.tmpl",
		"module.go":     "module.tmpl",
		"init.go":       "init.tmpl",
		"repository.go": "repository.tmpl",
		"service.go":    "service.tmpl",
		"routes.go":     "routes.tmpl",
	}

	data := ModuleTemplateData{
		PackageName: modulePkgName,
		ModuleName:  moduleName,
	}
	for file, tmpl := range templateFiles {
		err := GenerateFile(filepath.Join(modulePath, file), tmpl, data)
		if err != nil {
			return fmt.Errorf("failed to generate file %s from template %s: %w", file, tmpl, err)
		}
	}

	return WriteDeclarationFile(path, declaration)
}
