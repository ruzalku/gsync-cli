package utils

import (
	"errors"
	"strings"
	"net"
)

func ProcessDaemonMessage(client net.Conn) error {
	buf := make([]byte, 1024)
	n, err := client.Read(buf)

	if err != nil {
		return err
	}

	msgArr := strings.Split(string(buf[:n]), " ")
	status := msgArr[0]
	text := msgArr[1]

	if status == "1" {
		return errors.New(text)
	}

	return nil
}
