package config

import (
	"embed"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

const (
	developmentEnv = "development"
	ymlExtension   = ".yaml"
)

//go:embed *.yaml
var cfgFiles embed.FS

var configs map[string][]byte

func Initialize() error {
	return LoadConfigFromFile(Scope())
}

func Scope() string {
	env := os.Getenv("GO_ENVIRONMENT")
	if len(env) == 0 {
		env = "development" // default environment
	} else if env == "production" { // aka Fury
		env = os.Getenv("SCOPE") // use the scope name
	}
	return env
}

func LoadConfigFromFile(name string) error {
	filePath := getConfigFileName()
	file, err := cfgFiles.Open(filePath)
	if err != nil {
		return err
	}
	return LoadConfigFromReader(file)
}

func LoadConfigFromReader(reader io.Reader) error {
	configs = make(map[string][]byte)
	data := make(map[string]interface{})

	content, err := ioutil.ReadAll(reader)
	if err != nil {
		return errors.Wrap(err, "failed to read the config file")
	}

	if err := yaml.Unmarshal(content, data); err != nil {
		return errors.Wrap(err, "failed to load the config")
	}

	for section, config := range data {
		bytes, err := yaml.Marshal(config)
		if err != nil {
			return errors.WithStack(err)
		}
		configs[section] = bytes
	}
	return nil
}

func LoadConfigSection(section string, pointer interface{}) error {
	if len(configs) == 0 {
		if err := Initialize(); err != nil {
			return fmt.Errorf(`error initializing config "%s"`, err.Error())
		}
	}

	bytes, found := configs[section]
	if !found {
		return fmt.Errorf(`config section "%s" not found`, section)
	}
	return errors.Wrap(yaml.Unmarshal(bytes, pointer), "failed to load the config section")
}

func getConfigFileName() string {
	fileName, err := getScopeFromEnv()
	if err != nil {
		fileName = getEnvironment()
	}
	return fmt.Sprintf("%v%v", fileName, ymlExtension)
}

func getEnvironment() string {
	env := os.Getenv("GO_ENVIRONMENT")
	if env == "" {
		env = developmentEnv
	}
	return env
}

func getScopeFromEnv() (string, error) {
	scope := os.Getenv("SCOPE")
	if scope == "" {
		return "", errors.New("empty scope")
	}
	return scope, nil
}
