package cmd

import (
	"fmt"
	"github.com/smtdfc/photon-cli/generator"
	"github.com/spf13/cobra"
	"os"
)

var genAppCmd = &cobra.Command{
	Use:   "app",
	Short: "Generate a new app",
	Long:  `This command scaffolds a new app for your Photon project.`,
	Run: func(cmd *cobra.Command, args []string) {
		currentDir, err := os.Getwd()
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}

		if len(args) == 0 {
			fmt.Println("Please add app name")
			return
		}
		name := args[0]
		fmt.Printf("Generating app: %s \n", name)
		err = photonGenerator.GenerateApp(name, currentDir)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("App %s generated succesfully ! \n", name)
		}
	},
}

func init() {
	genCmd.AddCommand(genAppCmd)
}
