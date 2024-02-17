package api

import (
	"fmt"
	"github.com/atadzan/cobra-layout/internal/lib"
)

type CliArgs struct {
	SomeAdditionalArgForApi string
	ConfigPath              string
}

func RunApplication(cliArgs *lib.CliArgs, extCliArgs *CliArgs) {
	startServer(cliArgs.HTTPPort)
}

func startServer(port uint16) {
	fmt.Printf("Server started on %d port", port)
}
