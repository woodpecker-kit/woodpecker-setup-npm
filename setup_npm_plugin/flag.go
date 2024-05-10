package setup_npm_plugin

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/woodpecker-kit/woodpecker-tools/wd_flag"
	"github.com/woodpecker-kit/woodpecker-tools/wd_info"
	"github.com/woodpecker-kit/woodpecker-tools/wd_log"
	"github.com/woodpecker-kit/woodpecker-tools/wd_short_info"
	"strings"
)

const (
	CliNameNpmSetupMode = "settings.npm-setup-mode"
	EnvNameNpmSetupMode = "PLUGIN_NPM_SETUP_MODE"

	CliNameNpmRegistry = "settings.npm-registry"
	EnvNameNpmRegistry = "PLUGIN_NPM_REGISTRY"

	CliNameNpmUsername = "settings.npm-username"
	EnvNameNpmUsername = "PLUGIN_NPM_USERNAME"

	CliNameNpmPassword = "settings.npm-password"
	EnvNameNpmPassword = "PLUGIN_NPM_PASSWORD"

	CliNameNpmToken = "settings.npm-token"
	EnvNameNpmToken = "PLUGIN_NPM_TOKEN"

	CliNameNpmScopedList = "settings.npm-scoped-list"
	EnvNameNpmScopedList = "PLUGIN_NPM_SCOPED_LIST"

	CliNameSkipWhoAmi = "settings.npm-skip-whoami"
	EnvNameSkipWhoAmi = "PLUGIN_SKIP_WHOAMI"

	CliNameNpmFolder = "settings.npm-folder"
	EnvNameNpmFolder = "PLUGIN_NPM_FOLDER"

	CLiNameNpmDryRun = "settings.npm-dry-run"
	EnvNameNpmDryRun = "PLUGIN_NPM_DRY_RUN"
)

// GlobalFlag
// Other modules also have flags
func GlobalFlag() []cli.Flag {
	return []cli.Flag{

		&cli.StringFlag{
			Name:    CliNameNpmSetupMode,
			Usage:   fmt.Sprintf("setup mode, support: %v", strings.Join(setupModeSupport, ", ")),
			Value:   SetupModeNpmRc,
			EnvVars: []string{EnvNameNpmSetupMode},
		},

		&cli.StringFlag{
			Name:    CliNameNpmRegistry,
			Usage:   fmt.Sprintf("NPM registry to use when install packages. if empty will use %s", globalRegistry),
			EnvVars: []string{EnvNameNpmRegistry},
		},
		&cli.StringFlag{
			Name:    CliNameNpmUsername,
			Usage:   "NPM username to use when install packages.",
			EnvVars: []string{EnvNameNpmUsername},
		},
		&cli.StringFlag{
			Name:    CliNameNpmPassword,
			Usage:   "NPM password to use when install packages.",
			EnvVars: []string{EnvNameNpmPassword},
		},
		&cli.StringFlag{
			Name:    CliNameNpmToken,
			Usage:   "NPM token to use when install packages. if token is set, username and password will be ignored.",
			EnvVars: []string{EnvNameNpmToken},
		},

		&cli.StringFlag{
			Name:    CliNameNpmFolder,
			Usage:   "NPM folder to use when publishing packages which must containing package.json. default will use workspace",
			EnvVars: []string{EnvNameNpmFolder},
		},

		&cli.StringSliceFlag{
			Name:    CliNameNpmScopedList,
			Usage:   "NPM scoped list to use when install packages.",
			EnvVars: []string{EnvNameNpmScopedList},
		},

		&cli.BoolFlag{
			Name:    CliNameSkipWhoAmi,
			Usage:   "Skip npm whoami check",
			EnvVars: []string{EnvNameSkipWhoAmi},
		},

		&cli.BoolFlag{
			Name:    CLiNameNpmDryRun,
			Usage:   "dry run mode, will add NPM registry config but will print the command only",
			EnvVars: []string{EnvNameNpmDryRun},
		},
	}
}

func HideGlobalFlag() []cli.Flag {
	return []cli.Flag{}
}

func BindCliFlags(c *cli.Context,
	debug bool,
	cliName, cliVersion string,
	wdInfo *wd_info.WoodpeckerInfo,
	rootPath,
	stepsTransferPath string, stepsOutDisable bool,
) (*Plugin, error) {

	config := Settings{
		Debug:             debug,
		TimeoutSecond:     c.Uint(wd_flag.NameCliPluginTimeoutSecond),
		StepsTransferPath: stepsTransferPath,
		StepsOutDisable:   stepsOutDisable,
		RootPath:          rootPath,

		SetupMode: c.String(CliNameNpmSetupMode),

		Registry:   c.String(CliNameNpmRegistry),
		Username:   c.String(CliNameNpmUsername),
		Password:   c.String(CliNameNpmPassword),
		Token:      c.String(CliNameNpmToken),
		ScopedList: c.StringSlice(CliNameNpmScopedList),
		Folder:     c.String(CliNameNpmFolder),

		SkipWhoAmI: c.Bool(CliNameSkipWhoAmi),
		NpmDryRun:  c.Bool(CLiNameNpmDryRun),
	}

	wd_log.Debugf("args %s: %v", wd_flag.NameCliPluginTimeoutSecond, config.TimeoutSecond)

	infoShort := wd_short_info.ParseWoodpeckerInfo2Short(*wdInfo)

	p := Plugin{
		Name:           cliName,
		Version:        cliVersion,
		woodpeckerInfo: wdInfo,
		wdShortInfo:    &infoShort,
		Settings:       config,
	}

	return &p, nil
}
