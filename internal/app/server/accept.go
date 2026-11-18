package server

import (
	"errors"
	"io"
	"log/slog"
	"net"
	"os"

	"gitlab.univ-nantes.fr/iutna.info2.r305/proj/internal/pkg/proto"
)

func accept(main bool, port int, conn net.Conn) {
	slog.Info("Connection accepted", "port", port, "remoteAddr", conn.RemoteAddr())

	clients.Store(conn, "data")
	slog.Info("Clients refreshed", "count", proto.CountMap(&clients))

	defer func() {
		clients.Delete(conn)
		slog.Info("Clients refreshed", "count", proto.CountMap(&clients))

		conn.Close()
		slog.Info("Connection closed", "port", port, "remoteAddr", conn.RemoteAddr())
	}()

	for {
		b := make([]byte, 1024)
		// On attend une commande du client.
		str, err := proto.Read(timeout, conn, b)

		if err != nil {
			if err == io.EOF {
				return
			}

			if errors.Is(err, net.ErrClosed) {
				slog.Info("Connection terminated", "port", port, "remoteAddr", conn.RemoteAddr())
				return
			}

			if os.IsTimeout(err) {
				slog.Info("Connection timeout", "port", port, "remoteAddr", conn.RemoteAddr())
				return
			}

			proto.SlogError(err)
			return
		}

		request := str
		// On récupère la commande
		command := proto.GetCommand(request)
		// On récupère les arguments
		argument := proto.GetArgument(request)

		// En fonction de la commande on exécute une fonction.
		switch command {
		case "List":
			tasksCounter.Increment()
			list(conn)
			tasksCounter.Decrement()
		case "Get":
			if main {
				tasksCounter.Increment()
				get(conn, argument)
				tasksCounter.Decrement()
			} else {
				wrongClientPort(conn)
			}
		case "End":
			end(conn)
			return
		case "Hide":
			if !main {
				hide(conn, argument)
			} else {
				wrongClientPort(conn)
			}
		case "Reveal":
			if !main {
				reveal(conn, argument)
			} else {
				wrongClientPort(conn)
			}
		case "Terminate":
			if !main {
				terminate(conn)
			} else {
				wrongClientPort(conn)
			}
		case "Cd":
			cd(conn, argument)
		default:
			defaultCase(conn)
		}
	}
}
