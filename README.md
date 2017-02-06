# dread
Go package for directly reading a rune (char) from stdin.

Normally reading from Stdin waits for the user to press enter to enable the user to correct errors in the line. If you want to read the next character directly when it is typed, you can use this package.

Please be aware that enabling raw mode for the terminal disables all convenience functions of your terminal, including:
- Local echo
- Pasting text

## Performance
The performance is terrible. Do not use this to process large amounts of data. If you want to do that, please use this package: https://godoc.org/golang.org/x/crypto/ssh/terminal

## Example
```
package main

import (
	"log"

	"github.com/noxer/dread"
)

func main() {
  // Read 10 runes from and print them on the command line.
	for i := 0; i < 10; i++ {
		r, err := dread.ReadRune()
		if err != nil {
			log.Fatalf("%d: Error: %s\n", i, err)
		}
		log.Printf("%d: Rune: %c (%[2]U)", i, r)
	}

	log.Println("Done.")
}
```
