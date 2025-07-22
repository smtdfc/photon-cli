package photonGenerator

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"text/template"
)

type ModuleTemplateData struct {
	PackageName string
	ModuleName  string
	Module string
}

func GenerateFile(path string, tmplPath string, data any) error {
	_, callerFile, _, ok := runtime.Caller(0)
	if !ok {
		return fmt.Errorf("cannot determine caller location")
	}

	baseDir := filepath.Dir(callerFile)

	absTmplPath := filepath.Join(baseDir, "..", "templates", tmplPath)

	tmplContent, err := os.ReadFile(absTmplPath)
	if err != nil {
		return err
	}

	t, err := template.New(filepath.Base(tmplPath)).Parse(string(tmplContent))
	if err != nil {
		return err
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	return t.Execute(f, data)
}
