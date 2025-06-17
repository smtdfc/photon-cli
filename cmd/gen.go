package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate object ex: module,route ,etc",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gen called")
	},
}

func init() {
	rootCmd.AddCommand(genCmd)
}
