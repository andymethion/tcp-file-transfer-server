package client

import (
	"io"
	"log/slog"
	"net"
	"os"
	"strconv"

	"gitlab.univ-nantes.fr/iutna.info2.r305/proj/internal/pkg/proto"
)

func get(conn net.Conn, argument string) {
	b := make([]byte, 1024)
	// Récupère la résonse du serveur
	str, err := proto.Read(timeout, conn, b)

	if err != nil {
		proto.SlogError(err)
		return
	}

	// On gère au cas par cas le message du serveur
	switch str {
	case "WrongClientPort":
		slog.Info("WrongClientPort")
		return
	case "UnknownFile":
		slog.Info("UnknownFile")
	// Si le message est start on commence la procèdure Get.
	case "Start":
		// On attend le contenu du fichier.
		b = make([]byte, 1024)
		str, err = proto.Read(timeout, conn, b)

		if err != nil {
			proto.SlogError(err)
			return
		}
		// On convertie le message en int
		fileSize, err := strconv.Atoi(str)

		if err != nil {
			proto.SlogError(err)
			return
		}

		// On crée le fichier correspondant.
		file, err := os.Create(argument)

		if err != nil {
			proto.SlogError(err)
			return
		}

		defer file.Close()

		// On copie le message du serveur dans le fichier.
		_, err = io.CopyN(file, conn, int64(fileSize))

		if err != nil {
			proto.SlogError(err)
			return
		}
	}

	err = proto.Write(timeout, conn, "OK")

	if err != nil {
		proto.SlogError(err)
		return
	}
}
