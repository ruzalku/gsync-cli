package utils

import (
	"net"
	"log"
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

	buf := make([]byte, 1024)
	n, err := client.Read(buf)

	if err != nil {
		return err
	}

	log.Println(string(buf[:n]))

	return err
}
