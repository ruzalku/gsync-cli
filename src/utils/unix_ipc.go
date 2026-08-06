//go:build !windows

package utils

import (
	"net"
)

func newListener(path string) (net.Listener, error) {
	return net.Listen("unix", path)
}

func newClient(path string) (net.Conn, error) {
	return net.Dial("unix", path)
}
