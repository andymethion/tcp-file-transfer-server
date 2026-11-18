package server

import (
	"fmt"
	"net"
	"os"
	"strings"

	"gitlab.univ-nantes.fr/iutna.info2.r305/proj/internal/pkg/proto"
)

func cd(conn net.Conn, argument string) {
	// On récupère le répertoire courant du clienty.
	value, _ := clients.Load(conn)
	path := value.(string)

	if !strings.Contains(argument, "..") {
		path = fmt.Sprintf("%s/%s", path, argument)
		fileInfo, err := os.Stat(path)

		if err != nil && !os.IsNotExist(err) {
			proto.SlogError(err)
			return
		}

		if os.IsNotExist(err) || !fileInfo.IsDir() {
			err = proto.Write(timeout, conn, "UnknownDir")

			if err != nil {
				proto.SlogError(err)
				return
			}

			return
		}

		clients.Store(conn, path)
	}

	// Si l'argument est .. on remonte l'arborescence.
	// On vérifie que l'on ne "sort" pas du réportoire data pour des raison de sécurité.
	if argument == ".." {
		if path != "data" {
			directories := strings.Split(path, "/")
			directories = directories[:len(directories)-1]
			clients.Store(conn, strings.Join(directories, "/"))
		}
	}

	// On envoi OK au client.
	err := proto.Write(timeout, conn, "OK")

	if err != nil {
		proto.SlogError(err)
		return
	}
}
