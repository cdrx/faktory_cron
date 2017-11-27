package main

import (
	"flag"

	"github.com/robfig/cron"
)

var (
	version     = "master"
	config_path string
	config      *Config
	scheduler   *cron.Cron
)

func init() {
	flag.StringVar(&config_path, "config", "./crontab.yaml", "path to the configuration file")
	flag.Parse()
}

func main() {
	configureLogging()
	log.Infof("Cron for Faktory (version %s)", version)
	log.Debugf("Reading config from file: %v", config_path)

	config = NewConfig(config_path)
	err := config.Update()
	if err != nil {
		log.Fatalf("Error in config: %v", err)
	}

	scheduler = cron.New()
	for _, t := range config.Jobs {
		log.Infof("Running %v every %v", t.Name, t.Schedule)
		t.AddToScheduler()
	}
	log.Infof("Loaded %d scheduled tasks from %v", len(config.Jobs), config_path)

	scheduler.Start()
	defer scheduler.Stop()

	// run forever
	select {}
}
