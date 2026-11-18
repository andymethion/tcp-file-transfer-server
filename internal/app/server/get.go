package server

import (
	"io"
	"net"
	"os"
	"strconv"

	"gitlab.univ-nantes.fr/iutna.info2.r305/proj/internal/pkg/proto"
)

// Permet d'envoyer le fichier demandé par le client.
// Le fichier est envoyé si les conditions suivantes sont respectées
// 1. Le fichier ne doit pas être caché
// 2. Le fichier doit exister dans l'arborescence du client
func get(conn net.Conn, argument string) {
	// Vérification que le fichier n'est pas "hidden".
	_, ok := hiddenFiles.Load(argument)

	// S'assure que le fichier demandé n'est pas dans la liste des fichiers cachés

	if ok {
		// Si il est hidden alors on envoi "UnknownFile"
		err := proto.Write(timeout, conn, "UnknownFile")

		if err != nil {
			proto.SlogError(err)
			return
		}

		b := make([]byte, 1024)
		// On attend la réponse du client.
		_, err = proto.Read(timeout, conn, b)

		if err != nil {
			proto.SlogError(err)
			return
		}

		// On ferme le procéssus
		return
	}

	// Vérifie que le fichier demandé existe
	// Prend en compte la position du client dans l'arborescence

	value, _ := clients.Load(conn)
	path := value.(string)

	// On récupère le fichier.
	fileInfo, err := os.Stat(path + "/" + argument)

	if err != nil && os.IsExist(err) {
		proto.SlogError(err)
		return
	}

	// Si le fichier n'existe pas ou que c'est un répertoire on renvoie "UnknownFile"
	if os.IsNotExist(err) || fileInfo.IsDir() {
		err = proto.Write(timeout, conn, "UnknownFile")

		if err != nil {
			proto.SlogError(err)
			return
		}

		b := make([]byte, 1024)
		_, err := proto.Read(timeout, conn, b)

		if err != nil {
			proto.SlogError(err)
			return
		}

		return
	}

	// On envoi au client "start" pour commencer la procédure.

	err = proto.Write(timeout, conn, "Start")

	if err != nil {
		proto.SlogError(err)
		return
	}

	// On récupère le contenu du fichier.
	file, err := os.Open("data/" + argument)

	if err != nil {
		proto.SlogError(err)
		return
	}

	defer file.Close()

	// Envoie de la taille du fichier

	fileSize := strconv.FormatInt(fileInfo.Size(), 10)
	err = proto.Write(timeout, conn, fileSize)

	if err != nil {
		proto.SlogError(err)
		return
	}

	// Envoie du fichier au client

	_, err = io.CopyN(conn, file, fileInfo.Size())

	if err != nil {
		proto.SlogError(err)
		return
	}

	// Attend que le client répond par OK

	b := make([]byte, 1024)
	_, err = proto.Read(timeout, conn, b)

	if err != nil {
		proto.SlogError(err)
		return
	}
}
