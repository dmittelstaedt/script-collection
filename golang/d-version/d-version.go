// Package main implements updating current version in index.html of an
// overview site hosted in Apache Tomcat
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type configuration struct {
	CurrentVersionURL string
	IndexHTMLFile     string
	SearchString      string
}

func readConfig() configuration {
	var configuration configuration

	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(data, &configuration)
	if err != nil {
		log.Fatal(err)
	}

	return configuration
}

func getCurrentVersion() {

}

func getRemoteVersion() {

}

func updateCurrentVersion() {

}

func main() {
	configuration := readConfig()
	fmt.Printf("currentVersionURL: %v\n", configuration.CurrentVersionURL)
}
