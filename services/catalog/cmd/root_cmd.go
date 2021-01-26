package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/eavillacis/velociraptor/pkg/httputils"
	"github.com/spf13/cobra"
)

var configFile = ""

func init() {
	rootCmd.PersistentFlags().StringVar(&configFile, "config", httputils.GetEnv(), "config file (default is $HOME/.env or .env.test if exist)")
}

var rootCmd = &cobra.Command{
	Use:   "catalog",
	Short: "Catalog is the Catalog Service for Velociraptor",
	Long:  `REST Service for Velociraptor Catalog Service`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("use: catalog serve")
	},
}

// Execute commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
