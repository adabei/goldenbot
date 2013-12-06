package main

import (
  "database/sql"
  _ "github.com/mattn/go-sqlite3"
	"encoding/json"
	"flag"
	"github.com/adabei/goldenbot/events"
	"github.com/adabei/goldenbot/events/cod"
  "github.com/adabei/goldenbot/rcon"
	_ "github.com/adabei/goldenbot/rcon/q3"
	"github.com/adabei/goldenbot/tails"
	"io/ioutil"
	"log"
	"os"
	_ "strings"
)

type GoldenConfig struct {
	Address      string
	RCONPassword string
	LogfilePath  string
	SayPrefix    string
}

func main() {
	// Parse command line flags
	configPath := *flag.String("config", "golden.cfg", "the config file to use")
	flag.Parse()
	cfg := LoadConfig(configPath)

	// Initialize EventAggregator
	ea := events.NewAggregator()

	// Setup RCON connection
	rch := make(chan rcon.RCONQuery, 10)
	go rcon.Relay("q3", cfg.Address, cfg.RCONPassword, rch)

	// Database
  db, err := sql.Open("sqlite3", "./golden.sqlite3")
  if err != nil {
    log.Fatal(err)
  }

  defer db.Close()

  // If new??? setup???

  query := "create table players(id text not null primary key);"
  _, err = db.Exec(query)
  if err != nil {
    log.Printf("%q: %s\n", err, query)
    return
  }

	logchan := make(chan string)
	go tails.Tail(cfg.LogfilePath, logchan, false)
	for {
		line := <-logchan
		join := cod.Join{"GUIDDD123", 1, line}
		ea.Publish(join)
	}
}

func LoadConfig(path string) GoldenConfig {
	// Read config
	fi, err := os.Open(path)
	if err != nil {
		log.Fatal("Couldn't open config file: ", err)
	}
	defer fi.Close()

	b, err := ioutil.ReadAll(fi)
	if err != nil {
		log.Fatal("Couldn't read config file: ", err)
	}

	var cfg GoldenConfig
	json.Unmarshal(b, &cfg)
	return cfg
}

type Plugin interface {
	Start()
}
