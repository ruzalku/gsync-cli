package commands

import (
	"context"
	"errors"
	"io"
	"log"
	"net"
	"os"
	"strings"

	"github.com/urfave/cli/v3"
	"google.golang.org/api/drive/v3"

	"gsynccli/src/utils"
	"gsynccli/src/utils/ipc"
)


func Start(ctx context.Context, cmd *cli.Command) error {
	if cmd.Bool("quiet") {
		log.SetOutput(io.Discard)
	}
	_ = os.Remove(ipc.ConnectPath)
	l, err := net.Listen("unix", ipc.ConnectPath)
	if err != nil {
		return err
	}

	defer l.Close()
	service, err := drive.NewService(ctx)
	if err != nil {
		return err
	}
	
	log.Print("START")

	for {
		conn, err := l.Accept()
		if err != nil {
			return err
		}

		go func(c net.Conn) {
			defer c.Close()
			buf := make([]byte, 1024)
			n, err := c.Read(buf)

			if err != nil {
				log.Fatalln(err)
				return
			}

			err = ExecuteCommand(string(buf[:n]), service)
			msg := "1"
			if err != nil {
				msg = "0 " + string(err.Error())
			}

			_, err = c.Write([]byte(msg))

			if err != nil {
				log.Fatalln(err)
				return
			}
		}(conn)
	}
}

func ExecuteCommand(c string, service *drive.Service) error {
	arrCommands := strings.Split(c, ";")
	if strings.TrimSpace(arrCommands[0]) == "save" {
		autosave := false
		if arrCommands[2] == "1" {
			autosave = true
		}
		file := utils.GFile{
			Path: arrCommands[1],
			Autosave: autosave,
		}
		err := file.SetGoogleFileID()
		if err != nil {
			return err
		}

		err = file.SaveFile(service)
		return err
	} else {
		log.Println(arrCommands)
		return errors.New("Error of executing command")
	}
}

func Exit(ctx context.Context, cmd *cli.Command) error {
	return nil
}