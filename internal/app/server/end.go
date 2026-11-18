package server

import (
	"net"

	"gitlab.univ-nantes.fr/iutna.info2.r305/proj/internal/pkg/proto"
)

// Met un terme a la connection du client
func end(conn net.Conn) {
	// Ferme la connexion.
	err := conn.Close()

	if err != nil {
		proto.SlogError(err)
		return
	}
}
