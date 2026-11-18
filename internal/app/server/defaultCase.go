package server

import (
	"net"

	"gitlab.univ-nantes.fr/iutna.info2.r305/proj/internal/pkg/proto"
)

func defaultCase(conn net.Conn) {
	// Renvoi UnknownCommand.
	err := proto.Write(timeout, conn, "UnknownCommand")

	if err != nil {
		proto.SlogError(err)
		return
	}
}
