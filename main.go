package main

import (
	"fmt"
	"github/rashedkvm/maven-artifact/pkg/configuration"
	"github/rashedkvm/maven-artifact/pkg/mavenresolver"
	"os"
)

func main() {

	var configFile, useprovidedConfig = os.LookupEnv("MVN_CONFIG")

	if !useprovidedConfig {
		configFile = "config.yaml"
	}
	config, err := configuration.LoadNew(configFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if config.ActiveRepository == "" {
		fmt.Println("no active repo in configuration")
		os.Exit(1)
	}

	if config.ActiveRepo() == nil {
		fmt.Println("configured active repo not defined")
		os.Exit(1)
	}

	repo := mavenresolver.Repository{
		URL:      config.ActiveRepo().URL,
		Username: config.ActiveRepo().Username,
		Password: config.ActiveRepo().Password,
	}

	artifact := mavenresolver.Artifact{
		Id:      "java-hello-world",
		GroupId: "com.maventest.app",
		Version: "1.0.6-SNAPSHOT",
	}

	if err := artifact.Resolve(&repo); err != nil {
		fmt.Println(err)
	}

	fmt.Println(artifact.MetaXML)
	fmt.Println(artifact.ResolvedURL)
}
