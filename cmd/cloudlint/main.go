package main

import (
	"github.com/pipetail/cloudlint/internal/app/worker"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	cfgFile       string   // path of cnfiguration file
	analyzeChecks []string // list of checks we want to execute
	debug         bool     // indicates verbose output

	// default values
	analyzeChecksDefault = []string{"jedna", "dva"}

	// commands
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
	// root command
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/cloudlint.yaml)")
	rootCmd.PersistentFlags().StringArrayVar(&analyzeChecks, "checks", analyzeChecksDefault, "list of checks you want to run against infrastructure")
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "verbose output")

	// analyze command
	rootCmd.AddCommand(analyzeCmd)

}

func main() {
	rootCmd.Execute()
}

func analyze(cmd *cobra.Command, args []string) {

	log.SetLevel(log.ErrorLevel)
	if debug {
		log.SetLevel(log.DebugLevel)
	}

	log.WithField("checks", analyzeChecks).Debug("received list of checks")
	result := worker.Handle()
	worker.Print(result)
}
