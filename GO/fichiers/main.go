package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func recup_texte(nom_fichier string) (string, error) {
	fichier, err := os.Open(nom_fichier)
	if err != nil {
		return "", err
	}
	defer fichier.Close()
	contenu, err := ioutil.ReadAll(fichier)
	if err != nil {
		return "", err
	}
	texte := string(contenu)
	return texte, nil
}

func lancement_recup_texte() {
	texte, err := recup_texte("document.txt")
	if err != nil {
		panic(err)
	}
	println(texte)
}

func ajout_texte(nom_fichier string, texte_a_ajouter string) error {
	fichier, err := os.OpenFile(nom_fichier, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer fichier.Close()
	_, err = fichier.WriteString(texte_a_ajouter)
	if err != nil {
		return err
	}
	return nil
}

func lancement_ajout_texte() {
	err := ajout_texte("document.txt", "Texte ajouté")
	if err != nil {
		panic(err)
	}
}

func supprimer_contenu_fichier(nom_fichier string) error {
	fichier, err := os.OpenFile(nom_fichier, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer fichier.Close()
	return nil
}

func lancement_supprimer_contenu_fichier() {
	err := supprimer_contenu_fichier("document.txt")
	if err != nil {
		panic(err)
	}
}

func remplacer_contenu_fichier(nom_fichier string, texte_a_remplacer string) error {
	err := supprimer_contenu_fichier(nom_fichier)
	if err != nil {
		return err
	}
	err = ajout_texte(nom_fichier, texte_a_remplacer)
	if err != nil {
		return err
	}
	return nil
}

func lancement_remplacer_contenu_fichier() {
	err := remplacer_contenu_fichier("document.txt", "Texte remplacé")
	if err != nil {
		panic(err)
	}
}

// Faire un menu pour choisir entre les fonctions et relancer le menu apres chaque choix rajouter un choix quitter le programme
func menu() {
	println("1 - Afficher le contenu d'un fichier")
	println("2 - Ajouter du texte")
	println("3 - Supprimer le contenu d'un fichier")
	println("4 - Remplacer le contenu d'un fichier")
	println("5 - Quitter")
}

func lancement_menu() {
	menu()
	var choix int
	println("Entrez votre choix")
	fmt.Scan(&choix)
	switch choix {
	case 1:
		lancement_recup_texte()
		lancement_menu()
	case 2:
		lancement_ajout_texte()
		lancement_menu()
	case 3:
		lancement_supprimer_contenu_fichier()
		lancement_menu()
	case 4:
		lancement_remplacer_contenu_fichier()
		lancement_menu()
	case 5:
		println("Au revoir")
		os.Exit(0)
	default:
		println("Choix invalide")
	}
}

func main() {
	lancement_menu()
}
