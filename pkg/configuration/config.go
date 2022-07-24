package configuration

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type RepositoryType struct {
	URL      string `yaml:"url"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type RegistryType struct {
	Name       string         `yaml:"name"`
	Repository RepositoryType `yaml:"repository"`
}

type Configuration struct {
	ActiveRepository string         `yaml:"activeRepository"`
	Registry         []RegistryType `yaml:"registry"`
}

func LoadNew(fileName string) (*Configuration, error) {
	yamlFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	var yamlConfig Configuration
	err = yaml.Unmarshal(yamlFile, &yamlConfig)
	if err != nil {
		return nil, err
	}
	return &yamlConfig, nil
}

func (c *Configuration) ActiveRepo() *RepositoryType {
	for i := 0; i < len(c.Registry); i++ {
		if c.Registry[i].Name == c.ActiveRepository {
			return &c.Registry[i].Repository
		}
	}
	return nil
}
