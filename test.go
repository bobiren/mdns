package main

import (
	"log"
	"os"
	"os/signal"
	"time"

	"bonjour"
)

func main() {
	// Run registration (blocking call)
	s, err := bonjour.Register("airocov agentservice", "_airocov_agentserver._tcp", "local", 9515, "airocovserver.local", "192.168.22.76", nil, nil)
	if err != nil {
		log.Fatalln(err.Error())
	}

	// Ctrl+C handling
	handler := make(chan os.Signal, 1)
	signal.Notify(handler, os.Interrupt)
	for sig := range handler {
		if sig == os.Interrupt {
			s.Shutdown()
			time.Sleep(1e9)
			break
		}
	}
}

