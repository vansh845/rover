package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// launchCmd represents the launch command
var launchCmd = &cobra.Command{
	Use:   "launch",
	Short: "Deploys and runs your application on your server.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("launch called")
	},
  
}

func init() {
	rootCmd.AddCommand(launchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// launchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// launchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
