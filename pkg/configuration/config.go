package configuration

type RepositoryType struct {
	Name     string
	URL      string
	Username string
	Password string
}

type ConfigType struct {
	ActiveRepository string
	Registry         []RepositoryType
}

func LoadConfiguration(configurationFilePath string) (ConfigType, error) {

	return ConfigType{}, nil
}
