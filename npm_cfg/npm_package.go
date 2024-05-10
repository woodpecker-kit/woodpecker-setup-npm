package npm_cfg

const (
	// NpmJsRegistry defines the default NPM registry.
	NpmJsRegistry = "https://registry.npmjs.org/"

	NpmRcFileName = ".npmrc"
)

type (
	NpmPackageJson struct {
		// Name
		//
		Name string `json:"name,omitempty"`

		// Version
		//
		Version string `json:"version,omitempty"`

		// Registries
		//
		Registries map[string]string `json:"registries,omitempty"`
	}
)
