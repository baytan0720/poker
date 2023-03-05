package pty

import (
	"io"
	"net"
	"os"
	"os/exec"

	"github.com/creack/pty"
	"golang.org/x/term"
)

// NewTty tty is true terminal
func NewTty(l net.Listener, cmd *exec.Cmd) error {
	// important
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return err
	}
	defer func() { _ = term.Restore(int(os.Stdin.Fd()), oldState) }()

	// start cmd
	ptmx, err := pty.Start(cmd)
	if err != nil {
		return err
	}

	go func() {
		defer ptmx.Close()
		conn, err := l.Accept()
		if err != nil {
			return
		}
		defer conn.Close()

		go func() {
			_, _ = io.Copy(ptmx, conn)
			return
		}()
		_, _ = io.Copy(conn, ptmx)
	}()

	return nil
}
