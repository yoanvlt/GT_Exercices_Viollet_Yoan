package main

import (
	"fmt"
)

func choix_nb_allumette() int {
	var nb_allumette int
	fmt.Println("Choisissez un nombre d'allumette superieur a 4 : ")
	fmt.Scan(&nb_allumette)
	for nb_allumette < 4 {
		fmt.Println("Choisissez un nombre d'allumette superieur a 4 : ")
		fmt.Scan(&nb_allumette)
	}
	return nb_allumette
}

func affiche_allumette(nb_allumette int) {
	fmt.Println("Il reste ", nb_allumette, " allumettes")
}

func choix_nb_allumette_retire() int {
	var nb_allumette_retire int
	fmt.Println("Choisissez un nombre d'allumette a retirer entre 1 et 3 : ")
	fmt.Scan(&nb_allumette_retire)
	for nb_allumette_retire < 1 || nb_allumette_retire > 3 {
		fmt.Println("Choisissez un nombre d'allumette a retirer entre 1 et 3 : ")
		fmt.Scan(&nb_allumette_retire)
	}
	return nb_allumette_retire
}

func defaite(nb_allumette int) bool {
	if nb_allumette == 0 {
		fmt.Println("Vous avez perdu")
		return true
	}
	return false
}

func compteur_allumette(nb_allumette int, nb_allumette_retire int) int {
	nb_allumette -= nb_allumette_retire
	if nb_allumette < 0 {
		nb_allumette = 0
	}
	return nb_allumette
}

func joue() {
	nb_allumette := choix_nb_allumette()
	for !defaite(nb_allumette) {
		affiche_allumette(nb_allumette)
		nb_allumette_retire := choix_nb_allumette_retire()
		nb_allumette = compteur_allumette(nb_allumette, nb_allumette_retire)
	}
}

func main() {
	joue()
}
