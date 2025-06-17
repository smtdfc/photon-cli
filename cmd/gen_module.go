package cmd

import (
	"fmt"
	"github.com/smtdfc/photon-cli/generator"
	"github.com/spf13/cobra"
	"os"
)

var genModuleCmd = &cobra.Command{
	Use:   "module",
	Short: "Generate a new module",
	Long:  `This command scaffolds a new module for your Photon project.`,
	Run: func(cmd *cobra.Command, args []string) {
		currentDir, err := os.Getwd()
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}

		if len(args) == 0 {
			fmt.Println("Please add module name")
			return
		}

		name := args[0]
		fmt.Printf("Generating module %s\n", name)
		err = photonGenerator.GenerateModule(name, currentDir)
		if err != nil {
			fmt.Printf("%s\n", err)
		} else {
			fmt.Printf("Module %s generated successfully ! \n", name)
		}
	},
}

func init() {
	genCmd.AddCommand(genModuleCmd)
}
