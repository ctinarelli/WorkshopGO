//+build goroutine.go

package main

import (
	"fmt"
	"sync"
)

const (
	numPrint = 10
	toPrint  = "coucou"
)

func main() {
	// `make` permet de construire certaine variable comme les `[]` ou les `chan`.
	c := make(chan string)
	// Les `chan` ou channels, sont des conteneurs qui permettent de transmettre des valeurs
	// d'une goroutine à une autre.

	var wg sync.WaitGroup

	for i := 0; i < numPrint; i++ {
		// Le mot clé `go` permet de lancer une fonction dans une nouvelle goroutine.
		go ConcurrentPrint(c, &wg)
	}

	for i := 0; i < numPrint; i++ {
		// On utilise l'opérateur `<-` pour envoyer une valeur dans la channel.
		c <- fmt.Sprintf("%d %s", i, toPrint)
	}

	// L'utilisation du `WaitGroup` est necéssaire pour éviter que la goroutine
	// principale se finisse avant les autres.
	// Commentez le, vous verrez le résultat est totalement différent.
	wg.Wait()
}

// Une channel typé `<-chan` est une channel depuis laquelle il n'est possible que
// de recevoir des valeurs.
// A l'inverse, une channel typé `chan<-` est une channel avec laquelle il n'est possible
// que d'envoyer des valeurs.
func ConcurrentPrint(str <-chan string, wg *sync.WaitGroup) {
	wg.Add(1)

	// On utilise l'opérateur `<-` pour recevoir une valeur depuis la channel.
	fmt.Println(<-str)
	wg.Done()
}
