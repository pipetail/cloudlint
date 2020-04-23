package main

import (
	"fmt"

	"github.com/pipetail/cloudlint/internal/app/worker"
)

func main() {

	fmt.Println("hello pipetail")

	worker.Printhello()

	// defer klog.Flush()

	// baseName := filepath.Base(os.Args[0])

	// err := velero.NewCommand(baseName).Execute()
	// cmd.CheckError(err)
}
