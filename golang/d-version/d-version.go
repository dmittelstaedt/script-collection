// Package main implements updating current version in index.html of an
// overview site hosted in Apache Tomcat
package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/spf13/viper"

	"github.com/gocolly/colly"
)

const configurationFileName = "config"

type configuration struct {
	RemoteVersionURL  string
	CurrentVersionURL string
	IndexHTMLFile     string
	SearchString      string
	UseLocal          bool
}

func readConfig(configurationFilename string) configuration {
	var configuration configuration

	viper.SetConfigName(configurationFilename)
	viper.AddConfigPath(".")
	viper.ReadInConfig()

	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatal(err)
	}

	return configuration
}

func getCurrentVersion(currentVersionURL string) string {
	var currentVersion string

	collector := colly.NewCollector()
	collector.OnHTML("p[class=\"small text-center\"]", func(e *colly.HTMLElement) {
		currentVersion = e.Text
	})

	collector.Visit(currentVersionURL)

	return strings.TrimSpace(currentVersion)
}

func getCurrentVersionLocal() string {
	var currentVersion string

	fileTransport := &http.Transport{}
	fileTransport.RegisterProtocol("file", http.NewFileTransport(http.Dir("./")))

	collector := colly.NewCollector()
	collector.WithTransport(fileTransport)
	collector.OnHTML("p[class=\"small text-center\"]", func(e *colly.HTMLElement) {
		currentVersion = e.Text
	})

	err := collector.Visit("file://./index.html")
	if err != nil {
		log.Fatal(err)
	}

	return strings.TrimSpace(currentVersion)
}

func getRemoteVersion() {

}

func updateCurrentVersion() {

}

func main() {
	configuration := readConfig(configurationFileName)
	fmt.Printf("Current Version URL: %v\n", configuration.CurrentVersionURL)
	currentVersion := getCurrentVersion(configuration.CurrentVersionURL)
	fmt.Printf("Current Version: %v\n", currentVersion)
}
