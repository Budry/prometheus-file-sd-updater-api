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
	Use:   "prometheus-file-sd-updater-api [path] [token]",
	Short: "Add or remove hostname from Prometheus JSON file service discovery",
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {

		requiredToken := args[1]

		r := mux.NewRouter()
		r.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
			if isAuthorized(requiredToken, r) {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			vars := mux.Vars(r)
			targetFile := prometheus.NewTargetFile(args[0])
			targetFile.Append(vars["hostname"])
			w.WriteHeader(http.StatusNoContent)

		}).Methods("POST")
		r.HandleFunc("/remove", func(w http.ResponseWriter, r *http.Request) {
			if isAuthorized(requiredToken, r) {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			vars := mux.Vars(r)
			targetFile := prometheus.NewTargetFile(args[0])
			targetFile.Remove(vars["hostname"])
			w.WriteHeader(http.StatusNoContent)
		}).Methods("POST")

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

func isAuthorized(requiredToken string, r *http.Request) bool {
	token := r.Header.Get("Authorization")
	token = token[7:len(token)]

	return requiredToken != token
}
