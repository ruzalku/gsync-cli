package utils

import (
	"net"
	_ "log"
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

	defer client.Close()

	_, err = client.Write([]byte(command))

	if err != nil {
		return err
	}

	return ProcessDaemonMessage(client)
}
