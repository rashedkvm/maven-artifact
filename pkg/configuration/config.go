package configuration

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type RepositoryType struct {
	Name     string `yaml:"name"`
	URL      string `yaml:"url"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type ConfigType struct {
	ActiveRepository string           `yaml:"activeRepository"`
	Registry         []RepositoryType `yaml:"registry"`
}

type ConfigurationType struct {
	Configuration ConfigType `yaml:configuration`
}

// func Load(configurationFilePath string) (*ConfigType, error) {
// 	var config ConfigType
// 	var configReader io.Reader
// 	f, err := os.Open(configurationFilePath)
// 	if err != nil {
// 		return nil, err
// 	}
// 	configReader = f
// 	defer f.Close()

// 	d := yaml.NewYAMLOrJSONDecoder(configReader, 4096)
// 	for {
// 		if err := d.Decode(config); err != nil {
// 			if err == io.EOF {
// 				return &config, nil
// 			}
// 			return nil, err
// 		}
// 	}

// 	// return &config, nil
// }

func LoadNew(fileName string) (*ConfigurationType, error) {
	yamlFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	var yamlConfig ConfigurationType
	err = yaml.Unmarshal(yamlFile, &yamlConfig)
	if err != nil {
		return nil, err
	}
	return &yamlConfig, nil
}
