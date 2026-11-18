package server

import (
	"fmt"
	"log/slog"
	"net"
	"sync"
	"time"

	"gitlab.univ-nantes.fr/iutna.info2.r305/proj/internal/pkg/proto"
)

var timeout time.Duration
var clients sync.Map           // Liste des clients connectés
var tasksCounter proto.Counter // Nombre de tâches (list, get) en cours
var hiddenFiles sync.Map       // Liste des fichiers cachés

func Run(main bool, port int, _timeout time.Duration) {
	// Commence à écouter sur le port donner en argmuent
	timeout = _timeout
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

	if err != nil {
		proto.SlogError(err)
		return
	}

	slog.Info("Server opened", "port", port)

	defer func() {
		listener.Close()
		slog.Info("Server closed", "port", port)
	}()

	for {
		conn, err := listener.Accept()

		if err != nil {
			proto.SlogError(err)
			return
		}

		// En fonction du port du serveur on accepte différemment le client
		if main {
			go accept(main, port, conn)
		} else {
			accept(main, port, conn)
		}
	}
}
