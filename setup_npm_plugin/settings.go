package setup_npm_plugin

const (
	TimeoutSecondMinimum = 10

	// globalRegistry defines the default NPM registry.
	globalRegistry = "https://registry.npmjs.org/"

	SetupModeNpmRc = "npmrc"
)

type (
	// Settings setup_npm_plugin private config
	Settings struct {
		Debug             bool
		TimeoutSecond     uint
		StepsTransferPath string
		StepsOutDisable   bool
		RootPath          string

		DryRun bool

		SetupMode string

		Registry       string
		Username       string
		UserPassword   string
		Token          string
		ScopedList     []string
		Folder         string
		folderFullPath string // this will append root path as Folder

		VerdaccioUserTokenSupport bool

		NpmDryRun bool
	}
)

var (
	setupModeSupport = []string{
		SetupModeNpmRc,
	}
)
