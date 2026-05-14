package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/diniboy1123/usque/config"
	"github.com/diniboy1123/usque/tunnel"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	var (
		configFile  = flag.String("config", "config.json", "Path to configuration file")
		showVersion = flag.Bool("version", false, "Print version information and exit")
		verbose     = flag.Bool("verbose", false, "Enable verbose/debug logging")
	)
	flag.Parse()

	if *showVersion {
		fmt.Printf("usque %s (commit: %s, built: %s)\n", version, commit, date)
		os.Exit(0)
	}

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	if *verbose {
		log.Println("Verbose logging enabled")
	}

	// Load configuration from file
	cfg, err := config.Load(*configFile)
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	if *verbose {
		log.Printf("Configuration loaded from %s", *configFile)
	}

	// Start the tunnel with the provided configuration
	t, err := tunnel.New(cfg)
	if err != nil {
		log.Fatalf("Failed to initialise tunnel: %v", err)
	}

	if err := t.Run(); err != nil {
		log.Fatalf("Tunnel exited with error: %v", err)
	}
}
