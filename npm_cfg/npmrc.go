package npm_cfg

type NpmRcConfig struct {
	NpmRcConfigInterface NpmRcConfigInterface `json:"-"`

	folderFullPath string

	npmToken        string
	npmUsername     string
	npmUserPassword string

	npmRcUserHomeEnable bool

	npmPkg       *NpmPackageJson
	mockUserHome string
	nowRegistry  string
}

type NpmRcConfigInterface interface {
	CheckFolder() error

	WriteNpmRcFile(registry string, scopedList []string) error
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
	return &NpmRcConfig{}
}

type NpmRcConfigOption func(*NpmRcConfig)

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

func WithNpmRcUserHomeEnable(npmRcUserHomeEnable bool) NpmRcConfigOption {
	return func(o *NpmRcConfig) {
		o.npmRcUserHomeEnable = npmRcUserHomeEnable
	}
}
