package main

import (
	"github.com/pipetail/cloudlint/pkg/worker"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	cfgFile      string   // path of configuration file
	runChecks    []string // list of checks we want to execute
	regionFilter []string // list of regions we want to check
	logLevel     string

	// default values
	runChecksDefault    = []string{"ami_old", "ebs_unused"}
	regionFilterDefault = []string{"us-east-1", "eu-central-1"}
	logLevelDefault     = "ERROR"

	// commands
	rootCmd = &cobra.Command{
		Use:   "cloudlint",
		Short: "", // add some clever but short description
		Long:  "", // add even more clever description
	}

	runCmd = &cobra.Command{
		Use:   "run",
		Short: "",
		Long:  "",
		Run:   run,
	}
)

func init() {
	// root command
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/cloudlint.yaml)")
	rootCmd.PersistentFlags().StringArrayVar(&runChecks, "checks", runChecksDefault, "list of checks you want to run against infrastructure")
	rootCmd.PersistentFlags().StringArrayVar(&regionFilter, "regions", regionFilterDefault, "list of regions you want to run checks for")
	rootCmd.PersistentFlags().StringVar(&logLevel, "log-level", logLevelDefault, "log level")

	// run command
	rootCmd.AddCommand(runCmd)

}

func setLogLevel() {
	switch logLevel {
	case "ERROR":
		log.SetLevel(log.ErrorLevel)
	case "WANT":
		log.SetLevel(log.WarnLevel)
	case "INFO":
		log.SetLevel(log.InfoLevel)
	case "DEBUG":
		log.SetLevel(log.DebugLevel)
	case "FATAL":
		log.SetLevel(log.FatalLevel)
	case "TRACE":
		log.SetLevel(log.TraceLevel)
	case "PANIC":
		log.SetLevel(log.PanicLevel)
	default:
		log.WithField("log-level", logLevel).Warning("Wrong log level set. Falling back to ERROR")
		log.SetLevel(log.ErrorLevel)
	}
}

func main() {
	rootCmd.Execute()
}

func run(cmd *cobra.Command, args []string) {

	// set loglevel
	setLogLevel()

	log.WithField("checks", runChecks).Debug("received list of checks")
	result := worker.Handle()
	worker.Print(result)
}
