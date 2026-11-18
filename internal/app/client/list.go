package client

import (
	"log/slog"
	"net"

	"gitlab.univ-nantes.fr/iutna.info2.r305/proj/internal/pkg/proto"
)

// Demande l'arborescence des fichiers au serveur
func list(conn net.Conn) {
	b := make([]byte, 1024)
	str, err := proto.Read(timeout, conn, b)

	if err != nil {
		proto.SlogError(err)
		return
	}

	slog.Info("Received List\n" + str)

	// On envoi OK lorsque la liste est bien reçue
	err = proto.Write(timeout, conn, "OK")

	if err != nil {
		proto.SlogError(err)
		return
	}

	slog.Info("Sent OK")
}
