//+build print

// Chaque fichier doit avoir une déclaration de package.
// Un fichier ayant la fonction main doit être dans le package main.
package main

// Les packages importés.
import "fmt"

// En go la déclaration de fonction se fait avec le mot clé func, le nom de la fonction,
// les parenthèses avec les arguments, le type de retour et enfin les accolades
// avec le corps de la fonction à l'interieur.
func main() {
	// Les différents types de print.
	fmt.Print("fmt.Print: Un print basic\n")
	fmt.Println("fmt.Println: Pas besoin de \\n")
	fmt.Printf("fmt.Printf: Utile pour `%s`\n", "formater des strings")
	fmt.Print(fmt.Sprintf("fmt.Sprintf: Comme un `%s` mais qui return la string\n", "fmt.Printf"))
}
