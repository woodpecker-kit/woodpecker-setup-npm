package npm_cfg

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/user"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

func (n *NpmRcConfig) CheckFolder() error {

	if n.folderFullPath == "" {
		return errors.New("npm config check package.json path is empty")
	}

	npmPkg, err := readPackageFile(n.folderFullPath)
	if err != nil {
		return err
	}
	n.npmPkg = npmPkg

	return nil
}

// / readPackageFile reads the package file at the given path.
func readPackageFile(folder string) (*NpmPackageJson, error) {
	// Verify package.json file exists
	packagePath := filepath.Join(folder, "package.json")
	info, err := os.Stat(packagePath)

	if os.IsNotExist(err) {
		return nil, fmt.Errorf("no package.json at %s: %w", packagePath, err)
	}
	if info.IsDir() {
		return nil, fmt.Errorf("the package.json at %s is a directory", packagePath)
	}

	// Read the file
	file, err := os.ReadFile(packagePath)
	if err != nil {
		return nil, fmt.Errorf("could not read package.json at %s: %w", packagePath, err)
	}

	// Unmarshal the json data
	npm := NpmPackageJson{}
	err = json.Unmarshal(file, &npm)
	if err != nil {
		return nil, err
	}

	// Make sure values are present
	if npm.Name == "" {
		return nil, fmt.Errorf("no package name present")
	}
	if npm.Version == "" {
		return nil, fmt.Errorf("no package version present")
	}

	return &npm, nil
}

func (n *NpmRcConfig) FetchVerdaccioTokenByUserPass(verdaccioUrl string) error {
	if verdaccioUrl == "" {
		return fmt.Errorf("verdaccio url is empty")
	}
	verdaccioUrl = strings.TrimRight(verdaccioUrl, "/")

	if n.npmUsername == "" {
		return fmt.Errorf("verdaccio username is empty")
	}

	if n.npmUserPassword == "" {
		return fmt.Errorf("verdaccio password is empty")
	}

	url := fmt.Sprintf("%s/-/verdaccio/sec/login", verdaccioUrl)

	loginReq := VerdaccioLoginRequest{
		Name:     n.npmUsername,
		Password: n.npmUserPassword,
	}
	jsonReq, err := json.Marshal(loginReq)
	if err != nil {
		return fmt.Errorf("verdaccio error preparing request: %v", err)
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonReq))
	if err != nil {
		return fmt.Errorf("verdaccio error on request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{
		Timeout: time.Duration(n.apiTimeoutSecond) * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("verdaccio error on response: %v\n", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var loginResp VerdaccioLoginResponse
	err = json.Unmarshal(body, &loginResp)
	if err != nil {
		return fmt.Errorf("verdaccio error decoding response: %v", err)
	}

	n.npmToken = loginResp.Token

	return nil
}

func (n *NpmRcConfig) GetNpmRcWritePath() string {
	return n.writeNpmRcPath
}

// WriteNpmRcFile
// if registry is empty, use default NpmJsRegistry
func (n *NpmRcConfig) WriteNpmRcFile(registry string, scopedList []string) (string, error) {
	if n.npmPkg == nil {
		return "", fmt.Errorf("please use CheckFolder to check package.json first")
	}

	if registry == "" {
		registry = NpmJsRegistry
	}
	n.nowRegistry = registry

	var authContentFunc func(n *NpmRcConfig) string
	if n.npmToken == "" {
		authContentFunc = npmrcContentsUsernamePassword
	} else {
		authContentFunc = npmrcContentsToken
	}

	npmrcPath := path.Join(n.folderFullPath, NpmRcFileName)

	if n.npmRcUserHomeEnable {
		// write npmrc file
		home := "/root"
		if n.mockUserHome == "" {
			currentUser, err := user.Current()
			if err == nil {
				home = currentUser.HomeDir
			}
		} else {
			home = n.mockUserHome
		}
		npmrcPath = path.Join(home, NpmRcFileName)
	}
	n.writeNpmRcPath = npmrcPath

	authContents := authContentFunc(n)

	scopedContentList := parseScoped4NpmRc(n, scopedList)
	if len(scopedContentList) > 0 {
		authContents += "\n" + strings.Join(scopedContentList, "\n")
	}

	if n.dryRun {

		return authContents, nil
	}

	return authContents, os.WriteFile(npmrcPath, []byte(authContents), 0644)
}

const regPackageJsonRegistries = `^(@)([\w\.\-].*)(\/)(.*)$`

func parseScoped4NpmRc(n *NpmRcConfig, scopeList []string) []string {
	var scopedKeys []string
	if len(n.npmPkg.Registries) > 0 {
		compile := regexp.MustCompile(regPackageJsonRegistries)

		for k, v := range n.npmPkg.Registries {
			if v == n.nowRegistry {
				subRes := compile.FindStringSubmatch(k)
				if len(subRes) > 4 {
					scopedKey := subRes[2]
					scopedKeys = append(scopedKeys, scopedKey)
				}
			}
		}
	}

	if len(scopeList) > 0 {
		scopedKeys = append(scopedKeys, scopeList...)
	}

	if len(scopedKeys) == 0 {
		return nil
	}

	scopedKeys = stringArrRemoveDuplicates(scopedKeys)

	var scopedContentList []string
	for _, v := range scopedKeys {
		scopedContentList = append(scopedContentList, fmt.Sprintf("@%s:registry=%s", v, n.nowRegistry))
	}

	return scopedContentList
}

func npmrcContentsUsernamePassword(n *NpmRcConfig) string {
	// get the base64 encoded string
	authString := fmt.Sprintf("%s:%s", n.npmUsername, n.npmUserPassword)
	encoded := base64.StdEncoding.EncodeToString([]byte(authString))

	// create the file contents
	if n.nowRegistry == NpmJsRegistry {
		return fmt.Sprintf("_auth = %s", encoded)
	}
	registryString := parseRegistryString(n)

	return fmt.Sprintf("%s:_auth = %s", registryString, encoded)
}

func parseRegistryString(n *NpmRcConfig) interface{} {
	registry, _ := url.Parse(n.nowRegistry)
	registry.Scheme = "" // Reset the scheme to empty. This makes it so we will get a protocol relative URL.
	host, port, _ := net.SplitHostPort(registry.Host)
	if port == "80" || port == "443" {
		registry.Host = host // Remove standard ports as they're not supported in authToken since NPM 7.
	}
	registryString := registry.String()

	if !strings.HasSuffix(registryString, "/") {
		registryString += "/"
	}
	return registryString
}

func npmrcContentsToken(n *NpmRcConfig) string {
	registryString := parseRegistryString(n)
	return fmt.Sprintf("%s:_authToken=%s", registryString, n.npmToken)
}

func stringArrRemoveDuplicates(slc []string) []string {
	if len(slc) == 0 {
		return slc
	}
	if len(slc) < 1024 {
		return strRemoveDuplicatesByLoop(slc)
	} else {
		return strRemoveDuplicatesByMap(slc)
	}
}

func strRemoveDuplicatesByLoop(slc []string) []string {
	var result []string
	for i := range slc {
		flag := true
		for j := range result {
			if slc[i] == result[j] {
				flag = false
				break
			}
		}
		if flag {
			result = append(result, slc[i])
		}
	}
	return result
}

// strRemoveDuplicatesByMap
// must use slc size gather than 1024
func strRemoveDuplicatesByMap(slc []string) []string {
	var result []string
	tempMap := make(map[string]byte, 1024)
	for _, e := range slc {
		l := len(tempMap)
		tempMap[e] = 0
		if len(tempMap) != l { // If map length changes after map is added, elements do not duplicate
			result = append(result, e)
		}
	}
	return result
}
