package main

import (
	"fmt"

	"github.com/pipetail/cloudlint/internal/app/worker"
	"github.com/spf13/cobra"
)

var (
	cfgFile string
	rootCmd = &cobra.Command{
		Use:   "cloudlint",
		Short: "",
		Long:  "",
	}
)

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/cloudlint.yaml)")
}

func main() {
	rootCmd.Execute()
	fmt.Printf("configfile: %s\n", cfgFile)

	worker.Printhello()

	worker.Handle()

	// defer klog.Flush()

	// baseName := filepath.Base(os.Args[0])

	// err := velero.NewCommand(baseName).Execute()
	// cmd.CheckError(err)
}

func initConfig() {

}
