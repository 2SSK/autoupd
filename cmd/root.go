package cmd

import (
	"fmt"
	"os"

	"github.com/2SSK/autoupd/internal/utils"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "autoupd",
	Short: "A brief description of your application",
	Long:  ``,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting autoupd...")

		// Check if the user is running the command as root
		if os.Geteuid() != 0 {
			fmt.Fprintln(os.Stderr, "Please run this command with sudo.")
			os.Exit(1)
		}

		// Check if the log directory exists, if not, create it
		utils.CheckLogDir()
		utils.SetupLogger()

		// Check if update has already been performed today
		if ok := utils.WasUpdateSuccessful(); ok {
			utils.Logger.Println("Update already completed successfully today. Skipping...")
			os.Exit(0)
		}

		// Main functionality of the application
		if err := utils.PerformPackageUpdate(); err != nil {
			utils.Logger.Println("Update failed: ", err)
			os.Exit(1)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
