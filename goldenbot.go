package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type Plugin interface {
	Setup() error
	Start()
}

// Setup is a tiny helper to run database setups for the specified plugins
func Setup(plugins ...Plugin) error {
	for p := range plugins {
		err := p.Setup()
		if err != nil {
			return err
		}
	}
	return nil
}
