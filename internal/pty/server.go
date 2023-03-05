package pty

import (
	"github.com/creack/pty"
	"golang.org/x/term"
	"io"
	"net"
	"os"
	"os/exec"
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
		conn, err := l.Accept()
		if err != nil {
			return
		}
		go func() { _, _ = io.Copy(ptmx, conn) }()
		_, _ = io.Copy(conn, ptmx)
	}()

	return nil
}
