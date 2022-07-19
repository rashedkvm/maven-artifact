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

	if config.Configuration.ActiveRepository == "" {
		fmt.Println("no active repo in configuration")
		os.Exit(1)
	}

	// Add code starting here

	repo := mavenresolver.Repository{
		URL:      config.Configuration.Registry[0].URL,
		Username: ``,
		Password: ``,
	}

	artifact := mavenresolver.Artifact{
		Id:      "java-hello-world",
		GroupId: "com.maventest.app",
		Version: "1.0.4-SNAPSHOT",
	}

	if err := artifact.Resolve(&repo); err != nil {
		fmt.Println(err)
	}

	fmt.Println(artifact.MetaXML)
	fmt.Println(artifact.ResolvedURL)
}
