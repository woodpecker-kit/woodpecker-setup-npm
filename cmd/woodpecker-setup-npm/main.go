//go:build !test

package main

import (
	"github.com/gookit/color"
	"github.com/joho/godotenv"
	"github.com/woodpecker-kit/woodpecker-setup-npm"
	"github.com/woodpecker-kit/woodpecker-setup-npm/cmd/cli"
	"github.com/woodpecker-kit/woodpecker-setup-npm/internal/pkg_kit"
	"github.com/woodpecker-kit/woodpecker-tools/wd_log"
	os "os"
)

func main() {
	wd_log.SetLogLineDeep(wd_log.DefaultExtLogLineMaxDeep)
	pkg_kit.InitPkgJsonContent(woodpecker_setup_npm.PackageJson)

	// register helpers once
	//wd_template.RegisterSettings(wd_template.DefaultHelpers)

	// kubernetes runner patch
	if _, err := os.Stat("/run/drone/env"); err == nil {
		errDotEnv := godotenv.Overload("/run/drone/env")
		if errDotEnv != nil {
			wd_log.Fatalf("load /run/drone/env err: %v", errDotEnv)
		}
	}

	// load env file by env `PLUGIN_ENV_FILE`
	if envFile, set := os.LookupEnv("PLUGIN_ENV_FILE"); set {
		errLoadEnvFile := godotenv.Overload(envFile)
		if errLoadEnvFile != nil {
			wd_log.Fatalf("load env file %s err: %v", envFile, errLoadEnvFile)
		}
	}

	app := cli.NewCliApp()

	args := os.Args
	if err := app.Run(args); nil != err {
		color.Redf("cli err at %v\n", err)
	}
}
