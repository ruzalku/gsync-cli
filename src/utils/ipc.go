package utils

import (
	"net"
)

func GetListener(path string) (net.Listener, error) {
	return newListener(path)
}

func GetClient(path string) (net.Conn, error) {
	return newClient(path)
}


func AddCommand(path string, command string) error {
	client, err := GetClient(path)
	if err != nil {
		return err
	}

	_, err = client.Write([]byte(command))

	return err
}
