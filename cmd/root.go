package cmd

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Budry/prometheus-file-sd-updater-api/prometheus"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "prometheus-file-sd-updater-api",
	Short: "Add or remove hostname from Prometheus JSON file service discovery",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		r := mux.NewRouter()
		r.HandleFunc("/add/{hostname}", func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			targetFile := prometheus.NewTargetFile(args[0])
			targetFile.Append(vars["hostname"])
			w.WriteHeader(http.StatusNoContent)
		})
		r.HandleFunc("/remove/{hostname}", func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			targetFile := prometheus.NewTargetFile(args[0])
			targetFile.Remove(vars["hostname"])
			w.WriteHeader(http.StatusNoContent)
		})

		port, err := cmd.Flags().GetString("port")
		if err != nil {
			panic(err)
		}
		http.ListenAndServe(":" + port, r)
	},
}

func Execute() {

	rootCmd.Flags().StringP("port", "p", "80", "Set server port. Default 80")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
