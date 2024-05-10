package setup_npm_plugin_test

import (
	"fmt"
	"github.com/sinlov-go/unittest-kit/unittest_file_kit"
	"github.com/woodpecker-kit/woodpecker-setup-npm/npm_cfg"
	"github.com/woodpecker-kit/woodpecker-setup-npm/setup_npm_plugin"
	"github.com/woodpecker-kit/woodpecker-tools/wd_info"
	"github.com/woodpecker-kit/woodpecker-tools/wd_log"
	"github.com/woodpecker-kit/woodpecker-tools/wd_mock"
	"github.com/woodpecker-kit/woodpecker-tools/wd_short_info"
	"path/filepath"
	"strings"
	"testing"
)

func TestCheckArgsPlugin(t *testing.T) {
	t.Log("mock Plugin")

	testDataPathRoot, errTestDataPathRoot := testGoldenKit.GetOrCreateTestDataFullPath("check_args")
	if errTestDataPathRoot != nil {
		t.Fatal(errTestDataPathRoot)
	}

	// successArgs
	successArgsWoodpeckerInfo := *wd_mock.NewWoodpeckerInfo(
		wd_mock.FastWorkSpace(filepath.Join(testDataPathRoot, "successArgs")),
		wd_mock.FastCurrentStatus(wd_info.BuildStatusSuccess),
	)
	successArgsSettings := mockPluginSettings()
	successArgsSettings.Username = "foo"
	successArgsSettings.UserPassword = "bar"

	// emptySetupMode
	emptySetupModeWoodpeckerInfo := *wd_mock.NewWoodpeckerInfo(
		wd_mock.FastWorkSpace(filepath.Join(testDataPathRoot, "emptySetupMode")),
		wd_mock.FastCurrentStatus(wd_info.BuildStatusSuccess),
	)
	emptySetupModeSettings := mockPluginSettings()
	emptySetupModeSettings.SetupMode = ""

	// noArgsUsername
	noArgsUsernameWoodpeckerInfo := *wd_mock.NewWoodpeckerInfo(
		wd_mock.FastWorkSpace(filepath.Join(testDataPathRoot, "noArgsUsername")),
		wd_mock.FastCurrentStatus(wd_info.BuildStatusSuccess),
	)
	noArgsUsernameSettings := mockPluginSettings()
	noArgsUsernameSettings.Username = ""

	tests := []struct {
		name           string
		woodpeckerInfo wd_info.WoodpeckerInfo
		settings       setup_npm_plugin.Settings
		workRoot       string

		isDryRun          bool
		wantArgFlagNotErr bool
	}{
		{
			name:              "successArgs",
			woodpeckerInfo:    successArgsWoodpeckerInfo,
			settings:          successArgsSettings,
			wantArgFlagNotErr: true,
		},
		{
			name:           "emptySetupMode",
			woodpeckerInfo: emptySetupModeWoodpeckerInfo,
			settings:       emptySetupModeSettings,
		},
		{
			name:           "noArgsUsername",
			woodpeckerInfo: noArgsUsernameWoodpeckerInfo,
			settings:       noArgsUsernameSettings,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			p := mockPluginWithSettings(t, tc.woodpeckerInfo, tc.settings)
			p.OnlyArgsCheck()
			errPluginRun := p.Exec()
			if tc.wantArgFlagNotErr {
				if errPluginRun != nil {
					wdShotInfo := wd_short_info.ParseWoodpeckerInfo2Short(p.GetWoodPeckerInfo())
					wd_log.VerboseJsonf(wdShotInfo, "print WoodpeckerInfoShort")
					wd_log.VerboseJsonf(p.Settings, "print Settings")
					t.Fatalf("wantArgFlagNotErr %v\np.Exec() error:\n%v", tc.wantArgFlagNotErr, errPluginRun)
					return
				}
				infoShot := p.ShortInfo()
				wd_log.VerboseJsonf(infoShot, "print WoodpeckerInfoShort")
			} else {
				if errPluginRun == nil {
					t.Fatalf("test case [ %s ], wantArgFlagNotErr %v, but p.Exec() not error", tc.name, tc.wantArgFlagNotErr)
				}
				t.Logf("check args error: %v", errPluginRun)
			}
		})
	}
}

func TestPlugin(t *testing.T) {
	t.Log("do Plugin")
	if envCheck(t) {
		return
	}
	if envMustArgsCheck(t) {
		return
	}
	t.Log("mock Plugin args")

	testDataPathRoot, errTestDataPathRoot := testGoldenKit.GetOrCreateTestDataFullPath("setup_npm_plugin")
	if errTestDataPathRoot != nil {
		t.Fatal(errTestDataPathRoot)
	}

	// statusSuccess
	statusSuccessWoodpeckerInfo := *wd_mock.NewWoodpeckerInfo(
		wd_mock.FastWorkSpace(filepath.Join(testDataPathRoot, "statusSuccess")),
		wd_mock.FastCurrentStatus(wd_info.BuildStatusSuccess),
	)
	statusSuccessSettings := mockPluginSettings()

	// withScopedList
	withScopedListWoodpeckerInfo := *wd_mock.NewWoodpeckerInfo(
		wd_mock.FastWorkSpace(filepath.Join(testDataPathRoot, "withScopedList")),
		wd_mock.FastCurrentStatus(wd_info.BuildStatusSuccess),
	)
	withScopedListSettings := mockPluginSettings()
	withScopedListSettings.ScopedList = []string{
		"npm.foo.com",
		"npm.bar.com",
		"npm.baz.com",
	}

	tests := []struct {
		name           string
		woodpeckerInfo wd_info.WoodpeckerInfo
		settings       setup_npm_plugin.Settings
		workRoot       string

		mockScopedList []string
		mockVersion    string

		ossTransferKey  string
		ossTransferData interface{}

		isDryRun bool
		wantErr  bool
	}{
		{
			name:           "statusSuccess",
			woodpeckerInfo: statusSuccessWoodpeckerInfo,
			settings:       statusSuccessSettings,
			mockVersion:    "1.0.0",
			mockScopedList: []string{
				"npm.foo.com",
			},
		},
		{
			name:           "withScopedList",
			woodpeckerInfo: withScopedListWoodpeckerInfo,
			settings:       withScopedListSettings,
			mockVersion:    "1.0.0",
			mockScopedList: []string{
				"npm.foo.com",
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			p := mockPluginWithSettings(t, tc.woodpeckerInfo, tc.settings)
			p.Settings.DryRun = tc.isDryRun
			if tc.ossTransferKey != "" {
				errGenTransferData := generateTransferStepsOut(
					p,
					tc.ossTransferKey,
					tc.ossTransferData,
				)
				if errGenTransferData != nil {
					t.Fatal(errGenTransferData)
				}
			}

			errMockPackageJson := mockPackageJsonFile(p.Settings.RootPath, tc.name, tc.mockVersion, p.Settings.Registry, tc.mockScopedList)
			if errMockPackageJson != nil {
				t.Fatal(errMockPackageJson)
			}

			err := p.Exec()
			if (err != nil) != tc.wantErr {
				t.Errorf("FeishuPlugin.Exec() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
		})
	}
}

func mockPackageJsonFile(root, pkgName, version string, registry string, scopedList []string) error {
	pkgData := npm_cfg.NpmPackageJson{
		Name:    strings.ToLower(pkgName),
		Version: version,
	}

	if registry != "" && len(scopedList) > 0 {
		registriesMap := make(map[string]string)
		for _, scoped := range scopedList {
			newKey := fmt.Sprintf("@%s/*", scoped)
			registriesMap[newKey] = registry
		}
		pkgData.Registries = registriesMap
	}

	pkgJsonPath := filepath.Join(root, "package.json")
	return unittest_file_kit.WriteFileAsJsonBeauty(pkgJsonPath, pkgData, true)
}
