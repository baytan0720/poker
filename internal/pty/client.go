package pty

import (
	"io"
	"net"
	"os"

	"golang.org/x/term"
)

// NewPty pty is pseudo terminal
func NewPty(conn net.Conn, interactive bool) {
	// important
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	defer func() { _ = term.Restore(int(os.Stdin.Fd()), oldState) }()

	if interactive {
		go func() {
			_, err := io.Copy(conn, os.Stdin)
			if err != nil {
				panic(err)
			}
		}()
	}
	_, err = io.Copy(os.Stdout, conn)
	if err != nil {
		panic(err)
	}
}
