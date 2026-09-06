//go:build windows

package ipc

import (
	"net"
	"github.com/Microsoft/go-winio"
)

const ConnectPath string = `\\.\pipe\gsynccli`

func newListener(path string) (net.Listener, error) {
	return winio.ListenPipe(path, nil)
}

func newClient(path string) (net.Conn, error) {
	return winio.DialPipe(path, nil)
}
