package configuration

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// type RepositoryType struct {
// 	Name     string `yaml:"name"`
// 	URL      string `yaml:"url"`
// 	Username string `yaml:"username"`
// 	Password string `yaml:"password"`
// }

// type ConfigType struct {
// 	ActiveRepository string           `yaml:"activeRepository"`
// 	Registry         []RepositoryType `yaml:"registry"`
// }

type ConfigurationType struct {
	Config struct {
		ActiveRepository string `yaml:"activeRepository"`
		Registry         []struct {
			Repository struct {
				Name     string `yaml:"name"`
				URL      string `yaml:"url"`
				Username string `yaml:"username"`
				Password string `yaml:"password"`
			} `yaml:"repository"`
		} `yaml:"registry"`
	} `yaml:"config"`
}

func LoadNew(fileName string) (*ConfigurationType, error) {
	yamlFile, err := ioutil.ReadFile(fileName)

	if err != nil {
		return nil, err
	}

	// fmt.Println(string(yamlFile))

	var yamlConfig ConfigurationType
	err = yaml.Unmarshal(yamlFile, &yamlConfig)
	if err != nil {
		return nil, err
	}
	return &yamlConfig, nil
}
