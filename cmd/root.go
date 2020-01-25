package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/Budry/prometheus-file-sd-updater-api/prometheus"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
)

type RequestBody struct {
	Hostname string `json:"hostname"`
}

var rootCmd = &cobra.Command{
	Use:   "prometheus-file-sd-updater-api [path] [token]",
	Short: "Add or remove hostname from Prometheus JSON file service discovery",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {

		requiredToken := args[1]

		r := mux.NewRouter()
		r.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
			if isAuthorized(requiredToken, r) {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			requestBody, err := getRequestBody(r)
			if err != nil {
				return
			}

			targetFile := prometheus.NewTargetFile(args[0])
			targetFile.Append(requestBody.Hostname)
			w.WriteHeader(http.StatusNoContent)

		}).Methods("POST")
		r.HandleFunc("/remove", func(w http.ResponseWriter, r *http.Request) {
			if isAuthorized(requiredToken, r) {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			requestBody, err := getRequestBody(r)
			if err != nil {
				return
			}

			targetFile := prometheus.NewTargetFile(args[0])
			targetFile.Remove(requestBody.Hostname)
			w.WriteHeader(http.StatusNoContent)
		}).Methods("POST")

		port, err := cmd.Flags().GetString("port")
		if err != nil {
			panic(err)
		}
		err = http.ListenAndServe(":"+port, r)
		if err != nil {
			panic(err)
		}
	},
}

func Execute() {

	rootCmd.Flags().StringP("port", "p", "80", "Set server port. Default 80")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func getRequestBody(r *http.Request) (*RequestBody, error) {
	decoder := json.NewDecoder(r.Body)
	requestBody := &RequestBody{}
	if err := decoder.Decode(requestBody); err != nil {
		return nil, err
	}
	defer r.Body.Close()

	return requestBody, nil
}

func isAuthorized(requiredToken string, r *http.Request) bool {
	token := r.Header.Get("Authorization")
	token = token[7:len(token)]

	return requiredToken != token
}
