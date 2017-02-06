package dread

import (
	"bufio"
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

// ReadRune reads a single rune from stdin.
func ReadRune() (rune, error) {
	state, err := terminal.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return 0, err
	}

	r := bufio.NewReader(os.Stdin)
	ru, _, err := r.ReadRune()
	terminal.Restore(0, state)
	return ru, err
}
