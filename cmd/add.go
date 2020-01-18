package cmd

import (
	"github.com/Budry/prometheus-file-sd-updater-api/prometheus"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add [path] [hostname]",
	Short: "Add new server to Prometheus JSON file service discovery",
	Long: "Add new server to Prometheus JSON file service discovery",
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]
		hostname := args[1]
		targetFile := prometheus.NewTargetFile(path)
		targetFile.Append(hostname)
	},
}

