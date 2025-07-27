package commands

import (
	"fmt"
	"github.com/smtdfc/photon-cli/domain"
	"github.com/smtdfc/photon-cli/helpers"
	"github.com/urfave/cli/v2"
	"path/filepath"
)

func Test(c *cli.Context) error {
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
	
	fmt.Printf("[@%s] Testing  ...\n", appName)
	helpers.SpawnCommand("go", []string{"test", "./test/..."}, true)
	return nil
}
