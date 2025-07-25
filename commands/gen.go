package commands

import (
	"fmt"
	"github.com/smtdfc/photon-cli/domain"
	"github.com/smtdfc/photon-cli/helpers"
	"github.com/urfave/cli/v2"
	"path/filepath"
	"strings"
)

func GenModule(appName string, moduleName string, PkgName string, projPath string) error {

	fmt.Printf("[@%s] Generating module %s ...\n", appName, moduleName)

	if helpers.IsPascalCase(moduleName) == false {
		return fmt.Errorf("[@%s] Invalid module name. Module name must follow CamelCase format \n", appName)
	}

	callDir := GetCallerDir()
	normalizedName := strings.ToLower(moduleName)
	projectFilesMap := map[string]string{
		"modules/" + normalizedName + "/init.go": "../templates/module/init.go.tmpl",
		"domain/" + normalizedName + ".go":       "../templates/domain/module.go.tmpl",
	}

	data := map[string]any{
		"PkgName":              PkgName,
		"ModuleNormalizedName": normalizedName,
		"ModuleName":           moduleName,
	}

	for fileName, tmplFile := range projectFilesMap {
		realFilePath, err := helpers.EnsureDirAndResolve(filepath.Join(projPath, fileName))
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

	fmt.Printf("[@%s] Module %s created !\n", appName, moduleName)

	return nil
}

func Gen(c *cli.Context) error {
	if c.NArg() < 2 {
		return fmt.Errorf("Invalid arguments")
	}

	cwd := helpers.GetCWD()

	fmt.Println("Reading go.mod ...")
	goModFilePath := filepath.Join(cwd, "go.mod")
	if !helpers.FileExists(goModFilePath) {
		return fmt.Errorf("Cannot find go.mod file !")
	}

	fmt.Println("Reading configuration file ...")
	configFilePath := filepath.Join(cwd, "photon.config.json")

	if !helpers.FileExists(configFilePath) {
		return fmt.Errorf("Cannot find configuration file !")
	}

	config, err := helpers.LoadJSONFile[*domain.Config](configFilePath)
	if err != nil {
		return fmt.Errorf("Cannot read configuration file: %s", err)
	}

	appName := config.Name

	targetType := c.Args().Get(0)
	targetName := c.Args().Get(1)

	moduleName, err := helpers.GetModuleName(goModFilePath)
	if err != nil {
		moduleName = appName
	}

	if targetType == "module" {
		return GenModule(appName, targetName, moduleName, cwd)
	}

	return nil
}
