package main

import (
	"encoding/json"
	"fmt"
	"github.com/genkiroid/cert"
	"log"
	"os"
	"time"
)

type Request struct {
	Source  Source  `json:"source"`
	Version Version `json:"version"`
}

type Source struct {
	Domain    string `json:"domain"`
	ExpiresIn int    `json:"expires_in"`
}

type Version struct {
	Time time.Time `json:"time"`
}

func main() {
	var request Request

	err := json.NewDecoder(os.Stdin).Decode(&request)
	if err != nil {
		log.Print("request parse error: ", err.Error())
		os.Exit(1)
	}

	tlsCert := cert.NewCert(fmt.Sprintf("%s:443", request.Source.Domain))
	expiresAt, err := time.Parse("2006-01-02 15:04:05 -0700 MST", tlsCert.NotAfter)
	if err != nil {
		log.Print("tls cert parse error: ", err.Error())
		os.Exit(1)
	}

	var versions []Version

	previousTime := request.Version.Time
	if !previousTime.IsZero() {
		versions = append(versions, Version{Time: previousTime})
	}

	currentTime := time.Now()
	daysLeft := expiresAt.Sub(currentTime).Hours() / 24
	if int(daysLeft) <= request.Source.ExpiresIn {
		versions = append(versions, Version{Time: time.Now()})
	}

	json.NewEncoder(os.Stdout).Encode(versions)
}
