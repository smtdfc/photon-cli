package commands

import (
	"fmt"
	"github.com/smtdfc/photon-cli/helpers"
	"github.com/urfave/cli/v2"
	"path/filepath"
	"runtime"
)

func GetCallerDir() string {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		return ""
	}
	return filepath.Dir(filename)
}

func Init(c *cli.Context) error {

	if c.NArg() < 1 {
		return fmt.Errorf("Pleasd input name of project")
	}

	cwd := helpers.GetCWD()

	fmt.Println("Reading go.mod ...")
	goModFilePath := filepath.Join(cwd, "go.mod")
	if !helpers.FileExists(goModFilePath) {
		return fmt.Errorf("Cannot find go.mod file !")
	}

	projectName := c.Args().Get(0)
	fmt.Printf("Initializing project %s ...\n", projectName)

	moduleName, err := helpers.GetModuleName(goModFilePath)
	if err != nil {
		moduleName = projectName
	}

	callDir := GetCallerDir()
	projectFilesMap := map[string]string{
		"main.go":          "../templates/project/main.go.tmpl",
		"app/app.go":       "../templates/project/app/app.go.tmpl",
		"app/module.go":    "../templates/project/app/module.go.tmpl",
		"domain/domain.go": "../templates/project/domain/domain.go.tmpl",
	}

	data := map[string]any{
		"ModuleName": moduleName,
	}

	for fileName, tmplFile := range projectFilesMap {
		realFilePath, err := helpers.EnsureDirAndResolve(filepath.Join(cwd, fileName))
		if err != nil {
			return err
		}

		realTemplPath := filepath.Join(callDir, tmplFile)
		code, err := helpers.RenderTemplateFile(realTemplPath, data)
		if err != nil {
			return err
		}

		err = helpers.WriteFile(realFilePath, code)
		if err != nil {
			return err
		}

	}

	fmt.Printf("Project created successfully \n")

	return nil
}
