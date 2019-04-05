// Package main implements updating current version in index.html of an
// overview site hosted in Apache Tomcat.
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"

	"github.com/gocolly/colly"
)

// Constants for the configuration file name, search string in the overview
// site and search string for the comand site
const (
	configurationFileName = "config"
	overviewSearchElement = "p[class=\"small text-center\"]"
	comandSearchElement   = "span[class=\"releaseInformation\"]"
)

type configuration struct {
	RemoteVersionURL  string
	CurrentVersionURL string
	IndexHTMLFile     string
	SearchString      string
	UseLocal          bool
}

// ReadConfig parses the configuration file and returns a configuration
// struct.
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

func getVersionFromURL(currentVersionURL, searchElement, searchString string) string {
	var currentVersion string

	collector := colly.NewCollector()
	collector.OnHTML(searchElement, func(e *colly.HTMLElement) {
		currentVersion = e.Text
	})

	collector.Visit(currentVersionURL)

	return strings.TrimSpace(currentVersion)

}

func getVersionFromFile(fileName, searchElement, searchString string) string {
	var currentVersion string

	dir, _ := filepath.Split(fileName)

	fileTransport := &http.Transport{}
	fileTransport.RegisterProtocol("file", http.NewFileTransport(http.Dir(dir)))

	collector := colly.NewCollector()
	collector.WithTransport(fileTransport)
	collector.OnHTML(searchElement, func(e *colly.HTMLElement) {
		currentVersion = e.Text
	})

	err := collector.Visit("file://" + fileName)
	if err != nil {
		log.Fatal(err)
	}

	return strings.TrimSpace(currentVersion)
}

func updateCurrentVersion(fileName, oldVersion, newVersion string) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	fileContent := string(bytes)
	updatedFileContent := strings.Replace(fileContent, oldVersion, newVersion, 1)
	fmt.Println(updatedFileContent)

	err = ioutil.WriteFile(fileName, []byte(updatedFileContent), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	//TODO: Test go routines with time
	configuration := readConfig(configurationFileName)
	fmt.Printf("Current Version URL: %v\n", configuration.CurrentVersionURL)
	// currentVersion := getVersionFromURL(configuration.CurrentVersionURL, overviewSearchElement, configuration.SearchString)
	// fmt.Printf("Current Version: %v\n", currentVersion)
	currentVersion := getVersionFromFile(configuration.IndexHTMLFile, overviewSearchElement, configuration.SearchString)
	fmt.Printf("Current Version: %v\n", currentVersion)
}
