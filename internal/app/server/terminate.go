package server

import (
	"net"
	"os"
	"time"

	"gitlab.univ-nantes.fr/iutna.info2.r305/proj/internal/pkg/proto"
)

// Attend que toutes les tâches soient terminées avant de
// déconnecter tous les clients et fermer le serveur
func terminate(conn net.Conn) {

	// Tant qu'il y'a des requêtes Get et List on attend
	for tasksCounter.Load() > 0 {
		time.Sleep(time.Millisecond * 100)
	}

	// Déconnecte chaque client
	clients.Range(func(key any, value any) bool {
		client, ok := key.(net.Conn)

		if !ok {
			return true
		}

		if conn == client {
			return true
		}

		clients.Delete(client)

		err := client.Close()

		if err != nil {
			proto.SlogError(err)
			return true
		}

		return true
	})

	err := proto.Write(timeout, conn, "OK")

	if err != nil {
		proto.SlogError(err)
		return
	}

	err = conn.Close()

	if err != nil {
		proto.SlogError(err)
		return
	}

	os.Exit(0)
}
