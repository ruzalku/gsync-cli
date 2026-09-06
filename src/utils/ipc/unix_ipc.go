//go:build !windows

package ipc

import (
	"net"
)

const ConnectPath string = "/tmp/gsynccli.sock"

func newListener(path string) (net.Listener, error) {
	return net.Listen("unix", path)
}

func newClient(path string) (net.Conn, error) {
	return net.Dial("unix", path)
}
