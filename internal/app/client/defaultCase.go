package client

import (
	"log/slog"
	"net"

	"gitlab.univ-nantes.fr/iutna.info2.r305/proj/internal/pkg/proto"
)

func defaultCase(conn net.Conn) {
	b := make([]byte, 1024)
	// On lit le message du serveur
	s, err := proto.Read(timeout, conn, b)

	if err != nil {
		proto.SlogError(err)
		return
	}

	slog.Info(s)
}
