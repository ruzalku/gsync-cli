package utils

import (
	"errors"
	"net"
)

func ProcessDaemonMessage(client net.Conn) error {
	buf := make([]byte, 1024)
	n, err := client.Read(buf)

	if err != nil {
		return err
	}

	status := string(buf[0])
	text := string(buf[1:n])

	if status == "0" {
		return errors.New(text)
	}

	return nil
}
