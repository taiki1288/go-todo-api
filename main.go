package main

import (
	"os"
)

const defaultport = "8080"

func port() string {
	p := os.Getenv("PORT")
	if p != "" {
		return ":" + p
	}
	return ":" + defaultport
}