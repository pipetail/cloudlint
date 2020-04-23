package main

import (
	"github.com/pipetail/cloudlint/internal/app/worker"
	"github.com/spf13/cobra"
)

var (
	cfgFile string
	rootCmd = &cobra.Command{
		Use:   "cloudlint",
		Short: "", // add some clever but short description
		Long:  "", // add even more clever description
	}

	analyzeCmd = &cobra.Command{
		Use:   "analyze",
		Short: "",
		Long:  "",
		Run:   analyze,
	}
)

func init() {
	cobra.OnInitialize(initConfig)

	// root command
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/cloudlint.yaml)")

	// analyze command
	rootCmd.AddCommand(analyzeCmd)
}

func main() {
	rootCmd.Execute()
}

func initConfig() {
}

func analyze(cmd *cobra.Command, args []string) {
	worker.Printhello()
	worker.Handle()
}
