package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "prometheus-file-sd-updater-api",
	Short: "Add or remove hostname from Prometheus JSON file service discovery",
	Args: cobra.ExactArgs(3),
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
