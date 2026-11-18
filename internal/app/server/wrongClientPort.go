package server

import (
	"net"

	"gitlab.univ-nantes.fr/iutna.info2.r305/proj/internal/pkg/proto"
)

// Renvoi au client un message l'informant qu'il est sur le mauvais port pour effectuer une action.
func wrongClientPort(conn net.Conn) {
	err := proto.Write(timeout, conn, "WrongClientPort")

	if err != nil {
		proto.SlogError(err)
		return
	}
}
