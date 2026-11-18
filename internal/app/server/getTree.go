package server

import (
	"fmt"
	"strings"

	"gitlab.univ-nantes.fr/iutna.info2.r305/proj/internal/pkg/proto"
)

// Crée une représentation arborescente des fichiers et répertoires.
func getTree(files []File, arbo string, espacement string) (tree string, err error) {
	leng := len(files)

	// Récupérer le nom du répertoire courant
	fileName := strings.Split(arbo, "/")[len(strings.Split(arbo, "/"))-1]

	// Initialiser la liste avec le nom du répertoire courant
	tree = fmt.Sprintf(fileName)

	var style string
	// Parcourir les fichiers et répertoires
	for index, file := range files {

		// Déterminer le style de l'arborescence
		if index == leng-1 {
			style = "└── "
		} else {
			style = "├── "
		}

		// Si c'est un répertoire, appeler récursivement createArbo
		if file.IsDir {
			space := ""
			if style == "└── " {
				space = "    "
			} else {
				space = "│   "
			}

			newArbo := arbo + "/" + file.Name
			newFile, err := getFiles(newArbo)
			if err != nil {
				proto.SlogError(err)
				return "", err
			}
			newTree, err := getTree(newFile, newArbo, espacement+space)
			if err != nil {
				proto.SlogError(err)
				return "", err
			}
			tree += "\n" + espacement + style + newTree

			// Sinon, ajouter le fichier à la liste
		} else {
			tree += fmt.Sprintf("\n%s%s %s  [%d]", espacement, style, file.Name, file.Size)
		}

	}
	return tree, nil
}
