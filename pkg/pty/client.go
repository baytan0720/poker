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
			_, _ = io.Copy(conn, os.Stdin)
		}()
	}
	_, _ = io.Copy(os.Stdout, conn)
}
