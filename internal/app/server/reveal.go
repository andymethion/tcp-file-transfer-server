package server

import (
	"net"

	"gitlab.univ-nantes.fr/iutna.info2.r305/proj/internal/pkg/proto"
)

// Retire du fichier s'il existe dans la variable globale hiddenFiles
func reveal(conn net.Conn, fileName string) {

	// On cherche dans les fichiers cachés si le fichier voulue est présent.
	_, ok := hiddenFiles.Load(fileName)

	// Si la valeur n'est pas présente on envoie UnknownFile au client.
	if !ok {
		err := proto.Write(timeout, conn, "UnknownFile")

		if err != nil {
			proto.SlogError(err)
			return
		}
	}

	// Suppression du fichier dans les fichiers cachés
	hiddenFiles.Delete(fileName)

	// Confrmation au client de la procédure
	err := proto.Write(timeout, conn, "OK")

	if err != nil {
		proto.SlogError(err)
		return
	}
}
