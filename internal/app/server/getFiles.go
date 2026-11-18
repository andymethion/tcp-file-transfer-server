package server

import (
	"os"

	"gitlab.univ-nantes.fr/iutna.info2.r305/proj/internal/pkg/proto"
)

type File struct {
	Name  string
	Size  int64
	IsDir bool
}

// La fonction getFiles retourne la liste des fichiers **visibles** qui se trouvent dans le répertoire data.
func getFiles(name string) (files []File, err error) {
	// On récupère les différent fichier et dossier d'un dossier.
	dirEntries, err := os.ReadDir(name)

	if err != nil {
		proto.SlogError(err)
		return []File{}, err
	}

	// On ajoute à la struct File les différentes informations des fichiers et on les ajoutes à la liste files.
	for _, dirEntry := range dirEntries {
		fileInfo, err := dirEntry.Info()

		if err != nil {
			proto.SlogError(err)
			return []File{}, err
		}

		_, ok := hiddenFiles.Load(fileInfo.Name())

		if ok {
			continue
		}

		files = append(files, File{
			fileInfo.Name(),
			fileInfo.Size(),
			fileInfo.IsDir(),
		})
	}

	return files, nil
}
