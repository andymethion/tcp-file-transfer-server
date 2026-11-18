package client

import (
	"log/slog"
	"net"
	"time"

	"gitlab.univ-nantes.fr/iutna.info2.r305/proj/internal/pkg/proto"
)

var timeout time.Duration

func Run(address string, _timeout time.Duration) {
	// Établi la connexion avec le serveur
	timeout = _timeout
	conn, err := net.DialTimeout("tcp", address, timeout)

	// Retourne l'erreur s'il y'en a une
	if err != nil {
		proto.SlogError(err)
		return
	}

	defer conn.Close()

	for {
		// On récupère la commande donné par l'utilisateur
		request, err := proto.Input("request: ")

		if err != nil {
			proto.SlogError(err)
			return
		}

		// On récupère la commande et ses arguments
		command := proto.GetCommand(request)
		argument := proto.GetArgument(request)

		err = proto.Write(timeout, conn, request)

		if err != nil {
			proto.SlogError(err)
			return
		}

		if command == "" {
			slog.Info("Requested " + request)
		}

		// On exécute la fonction correspondante à la commande
		switch command {
		case "":
		case "List":
			list(conn)
		case "Get":
			get(conn, argument)
		case "End":
			return
		default:
			defaultCase(conn)
		}
	}
}
