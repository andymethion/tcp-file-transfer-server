package server

import (
	"log/slog"
	"net"

	"gitlab.univ-nantes.fr/iutna.info2.r305/proj/internal/pkg/proto"
)

// Renvoie l'arborescence des fichiers du serveur
func list(conn net.Conn) {
	slog.Info("Received List", "remoteAddr", conn.RemoteAddr().String())

	// On récupère le dossier courent du client.
	dir, _ := clients.Load(conn)

	// Appel de la fonction getTree qui renvoie un arbre représentant l'arborécense du dossier data.
	filearbo, err := getFiles(dir.(string))
	if err != nil {
		proto.SlogError(err)
		return
	}
	tree, err := getTree(filearbo, dir.(string), "")

	if err != nil {
		proto.SlogError(err)
		return
	}

	// Envoi de l'arbre au client.
	err = proto.Write(timeout, conn, tree)

	if err != nil {
		proto.SlogError(err)
		return
	}

	slog.Info("Responded List", "remoteAddr", conn.RemoteAddr().String())

	b := make([]byte, 1024)

	// On attend la réponse du Client.

	_, err = proto.Read(timeout, conn, b)

	if err != nil {
		proto.SlogError(err)
		return
	}

	slog.Info("Received OK", "remoteAddr", conn.RemoteAddr().String())
}
