/*
Copyright © 2025 nanvenomous mrgarelli@gmail.com
*/
package cmd

import (
	"log"
	"net/http"
	"os"

	"github.com/nanvenomous/htmx_templ_daisy_starter/handle"
	"github.com/spf13/cobra"
)

func runServeCmd() error {
	var (
		mux = http.NewServeMux()
	)

	err := handle.Setup(mux)
	if err != nil {
		return err
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "5005"
	}

	log.SetFlags(0)
	log.SetOutput(os.Stdout)

	log.Println("Serving on port " + port)
	return http.ListenAndServe(":"+port, mux)
}

// serveCmd servers the starter project
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "serve the starter project",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runServeCmd()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
