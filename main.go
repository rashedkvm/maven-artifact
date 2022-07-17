package main

import (
	"fmt"
	"github/rashedkvm/maven-artifact/pkg/mavenresolver"
)

func main() {

	repo := mavenresolver.Repository{
		// URL: `https://repo1.maven.org/maven2`,
		URL:      `https://maven.pkg.github.com/rashedkvm/tanzu-java-web-app`,
		Username: ``,
		Password: ``,
	}

	artifact := mavenresolver.Artifact{
		Id:      "demo",
		GroupId: "com.example",
		Version: "0.0.1-SNAPSHOT",
	}

	if err := artifact.Resolve(&repo); err != nil {
		fmt.Println(err)
	}

	fmt.Println(artifact.MetaXML)
	fmt.Println(artifact.ResolvedURL)
}
