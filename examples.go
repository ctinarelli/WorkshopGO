//+build examples

// Chaque fichier doit avoir une déclaration de package
// Le fichier contenant la fonction main doit être dans le package main
// Un seul fichier peut être dans le package main
package main

// Les packages importés
// Cette variante permet de n'utiliser qu'un seul mot clé import pour plusieurs packages
// mais il est aussi possible de voir `import packageName`
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
)

// Constantes globales
// Elles peuvent être soit des string soit numérique
// Précisé le type est optionel
// La valeur est donné avec `=`
const (
	apiUrl1           = "https://api.chucknorris.io/jokes/random"
	apiUrl2    string = "https://api.chucknorris.io/jokes/random"
	printCount        = 3
)

// Déclaration de structure
// Le mot clé type est obligatoire
// Les membres se déclarent dans l'ordre `nom` `type` `tag`
// Les tags sont des componsants optionels qui ne sont pas importants dans ce cadre
type ChuckNorrisFact struct {
	// Déclaration de `slice`
	// Les slices sont des tablaux à taille variable
	Category []string `json:"category"`
	IconURL  string   `json:"icon_url"`
	ID       string   `json:"id"`
	URL      string   `json:"url"`
	Value    string   `json:"value"`
}

// Déclaration de fonction
// Le mot clé func est obligatoire
// C'est une fonction simple qui ne prend pas d'argument et qui en return un
// qui est une string
func getChuckNorrisFact() ChuckNorrisFact {
	// Déclaration de variable
	// Le mot clé var est obligatoire
	// `f` est le nom de la variable et `ChuckNorrisFact` le type
	// Il n'est pas possible de donner une valeur à cette variable
	// lors de la déclaration
	var f ChuckNorrisFact

	// Une autre manière de déclarer des variables est de les assigner directement
	// avec `:=`
	// Le type de ces variables est géré automatiquement
	resp, _ := http.Get("https://api.chucknorris.io/jokes/random")

	// Une fonction peut return plusieurs arguements en GO
	// `_` permet d'ignorer certaines valeurs de retour
	body, _ := ioutil.ReadAll(resp.Body)

	// Cette fonction permet de remplir la structure f à partir du json
	// récupéré avec la requête http faite plus haut
	json.Unmarshal(body, &f)

	// Le mot clé `return` se comporte comme dans la plupart des autres langages
	// Comme le C, le python ou le java
	return f

	// Cette fonction fait une rêquete http à l'api chuck norris fact, parse la
	// réponse http puis le json de l'api et return la fact
}

// Une autre déclaration de fonction
// Celle ci prend deux arguments et n'en return aucun
// En GO les arguments sont déclarés dans l'ordre `nom` `type`
func LauchPrint(arg1 int, arg2 ChuckNorrisFact) {
	// En GO il existe deux fonctions standard qui permettent "d'allouer" des
	// variables.
	// La gestion de la mémoire est automatique mais certain type nécéssite
	// une "construction" qui se fait avec make
	// Cet appel à make donne une channel pour le type ChuckNorrisFact
	// Les channels sont des tableaux qui peuvent être passé à différentes
	// goroutine pour communiqué
	c := make(chan ChuckNorrisFact)

	// Le mot clé go permet de lancé une fonction parallèlement à l'exécution
	// actuel
	// Dans ce cas l'exécution de la fonction PrintFact va se dérouler en même
	// temps que la fonction LauchPrint
	go PrintFact(c)

	// Une boucle for classique, la différence avec le C est l'absence de
	// parenthèse `()` et l'obligation d'utilisé des acolades `{}` dans tous
	// les cas
	for i := 0; i < arg1; i++ {
		// L'utilisation de la flèche `<-` permet de "push" une valeur dans
		// une channel, ici `c`
		// Si la channel est pleinne alors la goroutine actuel bloque
		// jusqu'à que la channel ait la place nécéssaire
		c <- arg2
	}

	// Il est utile d'appellé la fonction standard `close` sur une channel
	// pour la rendre inutilisable
	close(c)

	// Cette fonction créée une channel, lance une goroutine avec la fonction
	// PrintFact en lui pasant la channel créée et "push" arg1 fois la variable
	// arg2 avant de `close` cette channel
}

// En GO l'ordre des déclarations des fonctions n'a pas d'importance,
// la fonction d'haut dessus arrive à trouver celle-ci
// L'arguement `c` est une channel dont il est seulement possible de "pull"
// puisque la flèche `<-` est spécifié devant
// Les channels pour lequels ils seulement possible de "push" sont noté `chan<-`
// Une channel sans restriction se note simplement `chan`
func PrintFact(c <-chan ChuckNorrisFact) {
	// Une for utilisant un range sur une channel
	// Le mot clé range permet de parcourir chaque valeur d'un conteneur
	// Ici c'est une channel
	// Utilisé `range` sur une channel a une particularité, c'est qu'à
	// chaque itération la plus vielle valeur présente dans la channel va
	// être récupérer, comme avec la flèche "pull" `<-` qui s'utilise ainsi:
	// `valeur <- c`
	// Comme pour la flèche "push", "pull" est une opération qui bloque
	// tant que les ressources nécéssaires ne sont pas disponibles,
	// et dans se cas tant qu'il n'y a pas de valeur dans la channel
	// qui ont été placé avec la flèche "push"
	// Chaque itération de cette bloque va donc se faire à chaque fois
	// qu'il y à une valeur qui est "push" dans la channel `c`
	// De plus, lorsque `range` est utilisé sur une channel, la boucle se
	// termine lorsque que `close` est utilisé sur la channel
	// Un equivalent en C pourait ressemblé à celà
	// for (Type v; is_open(v); v = pull(c)) {}
	for f := range c {

		// Par défault tous les types peuvent être print
		// Par exemple, ici il n'y a pas besoin d'avoir une fonction
		// qui transforme f qui de type `ChuckNorrisFact` en string
		// pour le print
		fmt.Println(f)
	}

	// Cette fonction va récupérer les ChuckNorrisFact et les print
	// tant que la channel `c` est ouverte
}

// Encore une autre déclaration de fonction
// Celle ci prend un argument de type `string` et return deux `string`
// Les arguments de retour sont nommés ce qui permet de s'y référer comme
// si c'était des variables déclarées au préalable
// De plus le type `string` n'est spécifié qu'une fois dans les arguments de retour,
// c'est une syntax qui permet de donner ce type à tous les arguments précédents
// dont le type n'a pas été spécifié
func CutString(str string) (str1, str2 string) {
	// Les if n'ont pas besoin de parenthèse `()` mais les acolades `{}`
	// sont obligatoires dans tous les cas
	// `len` est une fonction standard de GO qui return la longueur d'un type
	// qui comporte plusieurs éléments
	if len(str) <= 1 {
		str1 = str
	} else {
		// Deux slicing
		// D'un point de vue simplifié les strings fonctionnent comme
		// des arrays de char en GO
		// En GO il est possible de découper un array avec la notation
		// `[:]`
		// Ici str1 va prendre la moitié si la longueur et pair ou la
		// moitié + 1 si la longueur est impaire.
		// str2 va prendre la moitié arondie à l'inferieur dans tout les
		// cas
		str1 = str[:int(math.Ceil(float64(len(str))/2.0))]
		str2 = str[int(math.Ceil(float64(len(str))/2.0)):]
	}

	// Dans le cas où les valeurs de retour sont nommées il n'est pas nécéssaire
	// de spécifier des valeurs de retour mais il est toujours nécéssaire
	// d'utiliser le mot clé `return`
	return

	// Cette fonction prend une string et la découpe en deux avec la moitié
	// la plus longue dans la première valeur de retour et la plus courte dans
	// la seconde
}

// La fonction main
// Elle ne prend aucun arguments et n'en return aucun
// Pour récupérer les arguements du programme il faut utilisé `os.Argv` et pour avoir
// un code sortie autre que 0 il faut utilisé `os.Exit`
func main() {
	// Les différents print de base
	fmt.Println("Pas besoin de \\n")
	fmt.Print("Un print basic\n")
	fmt.Printf("Utilie pour %s\n", "formater des strings")

	fmt.Println(CutString("sourcil"))
	fmt.Println(CutString("saucecice"))

	LauchPrint(printCount, getChuckNorrisFact())

	// Il est possible de lancé rapidement un programme GO sans le recompilé
	// avec la commande go run monDuFichier sur un source du package main et
	// qui possède une fonction main
}
