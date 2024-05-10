package npm_cfg

const (
	minimumApiTimeoutSecond = 10
)

type NpmRcConfig struct {
	NpmRcConfigInterface NpmRcConfigInterface `json:"-"`

	dryRun bool

	folderFullPath string

	npmToken        string
	npmUsername     string
	npmUserPassword string

	npmRcUserHomeEnable bool

	apiTimeoutSecond uint

	npmPkg         *NpmPackageJson
	mockUserHome   string
	nowRegistry    string
	writeNpmRcPath string
}

type NpmRcConfigInterface interface {
	CheckFolder() error

	FetchVerdaccioTokenByUserPass(verdaccioUrl string) error

	WriteNpmRcFile(registry string, scopedList []string) (string, error)

	GetNpmRcWritePath() string
}

// NewNpmRcConfig
//
//	use as
//
//	changeNpmRcConfig := NewNpmRcConfig(
//	    npm_cfg.WithFolderFullPath(""),
//	)
func NewNpmRcConfig(opts ...NpmRcConfigOption) (opt *NpmRcConfig) {
	opt = defaultOptionNpmRcConfig
	for _, o := range opts {
		o(opt)
	}
	defaultOptionNpmRcConfig = setDefaultOptionNpmRcConfig()

	return
}

var (
	defaultOptionNpmRcConfig = setDefaultOptionNpmRcConfig()
)

func setDefaultOptionNpmRcConfig() *NpmRcConfig {
	return &NpmRcConfig{
		apiTimeoutSecond: minimumApiTimeoutSecond,
	}
}

type NpmRcConfigOption func(*NpmRcConfig)

func WithDryRun(dryRun bool) NpmRcConfigOption {
	return func(o *NpmRcConfig) {
		o.dryRun = dryRun
	}
}

func WithFolderFullPath(folderFullPath string) NpmRcConfigOption {
	return func(o *NpmRcConfig) {
		o.folderFullPath = folderFullPath
	}
}

func WithNpmToken(npmToken string) NpmRcConfigOption {
	return func(o *NpmRcConfig) {
		o.npmToken = npmToken
	}
}

func WithNpmUsername(npmUsername string) NpmRcConfigOption {
	return func(o *NpmRcConfig) {
		o.npmUsername = npmUsername
	}
}

func WithNpmUserPassword(npmUserPassword string) NpmRcConfigOption {
	return func(o *NpmRcConfig) {
		o.npmUserPassword = npmUserPassword
	}
}

func WithApiTimeoutSecond(apiTimeoutSecond uint) NpmRcConfigOption {
	return func(o *NpmRcConfig) {
		if apiTimeoutSecond > minimumApiTimeoutSecond {
			o.apiTimeoutSecond = apiTimeoutSecond
		}
	}
}

func WithNpmRcUserHomeEnable(npmRcUserHomeEnable bool) NpmRcConfigOption {
	return func(o *NpmRcConfig) {
		o.npmRcUserHomeEnable = npmRcUserHomeEnable
	}
}
