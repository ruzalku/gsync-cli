package ipc

import (
	"net"
	_ "log"

	"gsynccli/src/utils"
)

func GetListener(path string) (net.Listener, error) {
	return newListener(path)
}

func GetClient(path string) (net.Conn, error) {
	return newClient(path)
}


func AddCommand(path string, command []byte) error {
	client, err := GetClient(path)

	if err != nil {
		return err
	}

	defer client.Close()

	_, err = client.Write(command)

	if err != nil {
		return err
	}

	return utils.ProcessDaemonMessage(client)
}
