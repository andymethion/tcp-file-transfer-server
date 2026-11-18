package server

import (
	"net"

	"gitlab.univ-nantes.fr/iutna.info2.r305/proj/internal/pkg/proto"
)

// Ajoute du fichier s'il existe dans la variable globale hiddenFiles
func hide(conn net.Conn, fileName string) {
	// On récupère la liste de fichiers.
	files, err := getFiles("data")

	if err != nil {
		proto.SlogError(err)
		return
	}

	// On regarde pour tous les fichiers si il y'en a un qui correspond.
	for _, file := range files {
		// Si oui, on l'ajoute à la sync.Map hiddenFiles et on envoie OK au client.
		if fileName == file.Name {
			hiddenFiles.Store(file.Name, nil)

			err = proto.Write(timeout, conn, "OK")

			if err != nil {
				proto.SlogError(err)
				return
			}

			return
		}
	}

	// Sinon on envoi UnknownFile
	err = proto.Write(timeout, conn, "UnknownFile")

	if err != nil {
		proto.SlogError(err)
		return
	}
}
