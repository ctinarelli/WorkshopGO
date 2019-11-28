//+build hexshell

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Un exemple de boucle de `shell` qui parse des nombres
// et qui les affiche en hexadécimal.
// Pour la partie 2.
func main() {
	input := bufio.NewReader(os.Stdin)

	fmt.Print("> ")
	for in, err := input.ReadString('\n'); err == nil || err == io.EOF; in, err = input.ReadString('\n') {
		in = strings.TrimRight(in, "\n")
		if err != nil {
			fmt.Println("\nGoodbye :)")
			break
		}
		if in == "end" || in == "quit" || in == "exit" {
			fmt.Println("Goodbye :)")
			break
		}

		if n, parseErr := strconv.ParseInt(in, 10, 64); parseErr == nil {
			fmt.Printf("0x%x\n", n)
		} else {
			fmt.Printf("Can't parse `%s` as an integer\n", in)
		}

		fmt.Print("> ")
	}
}
