package cmd

import (
	"log"
	
	"user-app/internal/api"
	"user-app/internal/consts"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   consts.RootCommand,
	Short: consts.ProjectShortDescription,
	Long:  consts.ProjectLongDescription,
}

// Execute is the entry into the CLI, executing the root CMD.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

var apiCmd = &cobra.Command{
	Use:   consts.ApiCommand,
	Short: consts.ApiShortDescription,
	Long:  consts.ApiLongDescription,
	Run:   runApi,
}

func init() {
	RootCmd.AddCommand(apiCmd)
}

func runApi(*cobra.Command, []string) {
	server := api.NewServer()
	server.Start()
	server.Shutdown()
}
