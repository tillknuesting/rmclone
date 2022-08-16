package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"os/exec"
	"strings"
)

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) != 1 {
		log.Fatal("Usage: ./rmc <URL>")
	}

	URL := argsWithoutProg[0]

	_, err := url.ParseRequestURI(URL)
	if err != nil {
		log.Fatal("Usage: ./rmc <URL>")
	}

	name, err := extractName(URL)
	if err != nil {
		log.Fatal(err)
	}

	err = os.RemoveAll(name)
	if err != nil {
		log.Println(err)
	}

	err = gitClone(URL)
	if err != nil {
		log.Fatal(err)
	}
}

func extractName(url string) (string, error) {
	url = strings.TrimPrefix(url, "https://")

	urlSplit := strings.Split(url, "/")

	if len(urlSplit) == 0 {
		return "", fmt.Errorf("no / in URL")
	}

	if len(urlSplit) < 3 {
		return "", fmt.Errorf("no enough / in URL")
	}

	return strings.TrimSuffix(urlSplit[len(urlSplit)-1], ".git"), nil
}

func gitClone(url string) error {
	cmd := exec.Command("git", "clone", url)

	_, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("git clone: %w", err)
	}

	return nil
}
