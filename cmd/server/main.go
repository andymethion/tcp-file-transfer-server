package main

import (
	"flag"
	"log/slog"
	"time"

	"gitlab.univ-nantes.fr/iutna.info2.r305/proj/internal/app/server"
)

// Prend les attributs mis dans les flags
func parseArgs() (port int, timeout time.Duration) {
	portFlag := flag.Int("port", 3333, "Server Port (default: 3333)")
	timeoutFlag := flag.Int("timeout", 60, "Connection Timeout (default: 60)")
	debugFlag := flag.Bool("debug", false, "Enable Debug Mode (default: false)")

	flag.Parse()

	timeout = time.Duration(*timeoutFlag) * time.Second

	if *debugFlag {
		slog.SetLogLoggerLevel(slog.LevelDebug)
	}

	return *portFlag, timeout
}

func main() {
	port, timeout := parseArgs()

	go server.Run(true, port, timeout)
	go server.Run(false, port+1, timeout)

	select {}
}
