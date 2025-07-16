package cmd

import (
	"fmt"
	"os"

	"github.com/2SSK/autoupd/internal/ui/status"
	"github.com/2SSK/autoupd/internal/utils"
	"github.com/spf13/cobra"
)

var force bool
var showStatus bool

var rootCmd = &cobra.Command{
	Use:   "autoupd",
	Short: "Automatically update system packages with daily automation",
	Long:  ``,

	Run: func(cmd *cobra.Command, args []string) {
		if showStatus {
			status.RunDashboard()
			os.Exit(0)
		}

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
		if !force && utils.WasUpdateSuccessful() {
			utils.Logger.Println("Update already completed successfully today. Skipping...")
			os.Exit(0)
		}

		if !utils.IsTimerActive() {
			utils.Logger.Println("autoupd.timer not active. Setting up systemd...")
			if err := utils.SetupSystemdService(); err != nil {
				utils.Logger.Printf("Systemd setup failed: %v", err)
			}
		}

		// Main functionality of the application
		if err := utils.PerformPackageUpdate(); err != nil {
			utils.Logger.Println("Update failed: ", err)
			utils.NotifyFailure("System Update Failed.")
			os.Exit(1)
		}
		utils.NotifySuccess("System Update completed Successfully.")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolVarP(&force, "force", "f", false, "Force update even if already done today")
	rootCmd.Flags().BoolVarP(&showStatus, "status", "s", false, "Show autoupd status without performing update")
}
