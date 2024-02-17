package cmd

import (
	"fmt"
	"github.com/atadzan/cobra-layout/internal/api"
	eventHandler "github.com/atadzan/cobra-layout/internal/event_handler"
	"github.com/atadzan/cobra-layout/internal/lib"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var (
	cliArgs lib.CliArgs
)

var rootCmd = &cobra.Command{
	Use:   "advertising service",
	Short: "Advertising service",
	Run: func(cmd *cobra.Command, args []string) {
		log.Fatal("No app specified")
	}}

var (
	apiArgs api.CliArgs

	apiCmd = &cobra.Command{
		Use:   "api",
		Short: "API service",
		Run: func(cmd *cobra.Command, args []string) {
			api.RunApplication(&cliArgs, &apiArgs)
		},
	}
)

var (
	eventArgs eventHandler.CliArgs

	eventCmd = &cobra.Command{
		Use:   "event-handler",
		Short: "Advertising events handler",
		Run: func(cmd *cobra.Command, args []string) {
			eventHandler.RunApplication(&cliArgs, &eventArgs)
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cliArgs.LogPath, "log-file", "", "path to log file")
	rootCmd.PersistentFlags().StringVar(&cliArgs.ConfigPath, "config", "", "path to config file")
	rootCmd.PersistentFlags().Uint16Var(&cliArgs.HTTPPort, "port", 80, "http port to start on")
	apiCmd.PersistentFlags().StringVar(&apiArgs.SomeAdditionalArgForApi, "some-add-arg", "", "path to access log file for api service")
	apiCmd.PersistentFlags().StringVar(&apiArgs.ConfigPath, "config-path", "default-value", "path to config file for api service")
	eventCmd.PersistentFlags().StringVar(&eventArgs.ConfigPath, "config-path", "default-value", "path to config file for api service")

	rootCmd.AddCommand(apiCmd)
	rootCmd.AddCommand(eventCmd)
}
