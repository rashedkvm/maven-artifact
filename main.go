package main

import (
	"fmt"
	"github/rashedkvm/maven-artifact/pkg/mavenresolver"
	"io"
	"os"
)

func main() {

	var configFile, useprovidedConfig = os.LookupEnv("MVN_CONFIG")

	if useprovidedConfig == false {
		configFile = "config.yaml"
	}
	var configReader io.Reader
	f, err := os.Open(configFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	configReader = f
	defer f.Close()

	// Add code starting here

	repo := mavenresolver.Repository{
		URL:      `https://repo1.maven.org/maven2`,
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
