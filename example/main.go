package main

import (
	"log"

	"github.com/noxer/dread"
)

func main() {
	for i := 0; i < 10; i++ {
		r, err := dread.ReadRune()
		if err != nil {
			log.Fatalf("%d: Error: %s\n", i, err)
		}
		log.Printf("%d: Rune: %c (%[2]U)", i, r)
	}

	log.Println("Done.")
}
