//+build http

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Décalration de constante.
// Elles peuvent être soit des strings, soit des valeurs numériques.
const apiUrl = "https://api.chucknorris.io/jokes/random"

// Déclaration de structure.
// Les champs de structures sont déclarés dans l'ordre: nom type tag.
// Les tags sont des éléments qui sont utilisés pour la séréalisation.
type ChuckNorrisFact struct {
	Category []string `json:"category"`
	IconURL  string   `json:"icon_url"`
	ID       string   `json:"id"`
	URL      string   `json:"url"`
	Value    string   `json:"value"`
}

func main() {
	// Déclaration de variable.
	// Avec cette notation elles sont initialisées avec leur valeur par défault,
	// 0 pour les nombres, "" pour les strings, nil pour les pointeurs...
	var f ChuckNorrisFact

	// Une autre façon de déclarer une variable.
	// Avec `:=` le type est déterminé automatiquement.
	// Dans ce cas il ne faut pas utiliser le mot clé `var`.
	b := doHTTPGet(apiUrl)

	json.Unmarshal(b, &f)
	fmt.Println(f)
}

// Il est possible de nommer les variables de retour.
// Dans ce cas c'est comme si `var body []byte` était placé
// au début du corps de la fonction.
func doHTTPGet(url string) (body []byte) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}

	body, _ = ioutil.ReadAll(resp.Body)

	// Lorsque les variables sont nommées, elles sont automatiquement
	// retournées avec `return`.
	return
}
