package main

import (
	"flag"
	"fmt"
	"log/slog"
	"time"

	"gitlab.univ-nantes.fr/iutna.info2.r305/proj/internal/app/client"
)

// Prend les attributs mis dans les flags
func parseArgs() (address string, timeout time.Duration) {
	hostFlag := flag.String("host", "127.0.0.1", "Server Host (default: 127.0.0.1)")
	portFlag := flag.Int("port", 3333, "Server Port (default: 3333)")
	timeoutFlag := flag.Int("timeout", 60, "Connection Timeout (default: 60)")
	debugFlag := flag.Bool("debug", false, "Enable Debug Mode (default: false)")

	flag.Parse()

	address = fmt.Sprintf("%s:%d", *hostFlag, *portFlag)
	timeout = time.Duration(*timeoutFlag) * time.Second

	if *debugFlag {
		slog.SetLogLoggerLevel(slog.LevelDebug)
	}

	return address, timeout
}

func main() {
	address, timeout := parseArgs()
	client.Run(address, timeout)
}
