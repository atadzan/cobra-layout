package eventHandler

import (
	"fmt"
	"github.com/atadzan/cobra-layout/internal/lib"
)

type CliArgs struct {
	ConfigPath string
}

func RunApplication(cliArgs *lib.CliArgs, extArgs *CliArgs) {
	startEventHandler(cliArgs.ConfigPath)
}
func startEventHandler(_ string) {
	fmt.Print("Event handler started ...")
}
