package commands

import (
	"fmt"
	"github.com/smtdfc/photon-cli/domain"
	"github.com/smtdfc/photon-cli/helpers"
	"github.com/urfave/cli/v2"
	"path/filepath"
)

func Dev(c *cli.Context) error {
	fmt.Println("Reading configuration file ...")
	cwd := helpers.GetCWD()
	configFilePath := filepath.Join(cwd, "photon.config.json")

	if !helpers.FileExists(configFilePath) {
		return fmt.Errorf("Cannot find configuration file !")
	}

	config, err := helpers.LoadJSONFile[*domain.Config](configFilePath)
	if err != nil {
		return fmt.Errorf("Cannot read configuration file: %s", err)
	}

	appName := config.Name
	entryPoint := config.EntryPoint

	fmt.Printf("[@%s] Starting application ...\n", appName)
	helpers.SpawnCommand("go", []string{"run", entryPoint}, true)
	return nil
}
