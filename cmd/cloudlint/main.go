package main

import (
    ins "github.com/pipetail/cloudlint/pkg/inspection"
    "github.com/pipetail/cloudlint/pkg/worker"
    log "github.com/sirupsen/logrus"
    "github.com/spf13/cobra"
)

var (
	cfgFile      string   // path of configuration file
	checkFilter  []string // list of checks we want to execute
	regionFilter []string // list of regions we want to check
	logLevel     string
    checkLevel   string

	// default values
	checkFilterDefault  = []string{}
	regionFilterDefault = []string{}
	logLevelDefault     = "ERROR"
    checkLevelDefault   = "BASE"

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
	rootCmd.PersistentFlags().StringArrayVar(&checkFilter, "checks", checkFilterDefault, "list of checks you want to run against infrastructure")
	rootCmd.PersistentFlags().StringArrayVar(&regionFilter, "regions", regionFilterDefault, "list of regions you want to run checks for")
	rootCmd.PersistentFlags().StringVar(&logLevel, "log-level", logLevelDefault, "log level")
    rootCmd.PersistentFlags().StringVar(&checkLevel, "check-level", checkLevelDefault, "level of a check")

	// run command
	rootCmd.AddCommand(runCmd)

}

func setLogLevel() {
	switch logLevel {
	case "ERROR":
		log.SetLevel(log.ErrorLevel)
	case "WARN":
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

func setCheckLevel() {
    switch checkLevel {
    case "BASE":
       ins.SetLevel(ins.Base)
    case "DETAIL":
        ins.SetLevel(ins.Detail)
    default:
        log.WithField("check-level", checkLevel).Warning("Wrong check level set. Falling back to BASE")
        ins.SetLevel(ins.Base)
    }
}

func main() {
	rootCmd.Execute()
}

func run(cmd *cobra.Command, args []string) {

	// set loglevel
	setLogLevel()
    setCheckLevel()

	log.WithField("checks", checkFilter).Info("received list of checks")
	result := worker.Handle(checkFilter)
	worker.Print(result)
}
