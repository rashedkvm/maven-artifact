package main

import (
	"fmt"
	"github/rashedkvm/maven-artifact/pkg/repository"
	"os"
)

func main() {
	var url, isSet = os.LookupEnv("REPOSITORY_URL")
	if !isSet {
		fmt.Fprintf(os.Stderr, "Warning: Environment variable REPOSITORY_URL missing. Using default.\n")
		url = `https://repo1.maven.org/maven2/HTTPClient/HTTPClient/maven-metadata.xml`
	}

	cl := repository.Client()
	resp, err := cl.Get(url)

	if err != nil {
		fmt.Printf("error from %q %v \n", url, err)
	}

	if resp != nil {
		defer resp.Body.Close()
	}

	fmt.Printf("url: %q \nhttp status: %q", url, resp.Status)
}
