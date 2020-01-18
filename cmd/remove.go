package cmd

import (
	"github.com/Budry/prometheus-file-sd-updater-api/prometheus"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(removeCmd)
}

var removeCmd = &cobra.Command{
	Use:   "remove [path] [hostname]",
	Short: "Remove server from Prometheus JSON file service discovery",
	Long: "Remove server from Prometheus JSON file service discovery",
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]
		hostname := args[1]
		targetFile := prometheus.NewTargetFile(path)
		targetFile.Remove(hostname)
	},
}

