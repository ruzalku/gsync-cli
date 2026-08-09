package commands

import (
	"log"
	"context"
	"io"
	"net"
	"os"

	"github.com/urfave/cli/v3"
	"gsynccli/src/utils"
)


func Start(ctx context.Context, cmd *cli.Command) error {
	if cmd.Bool("quiet") {
		log.SetOutput(io.Discard)
	}
	_ = os.Remove(utils.ConnectPath)
	l, err := net.Listen("unix", utils.ConnectPath)
	if err != nil {
		return err
	}

	defer l.Close()
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

			msg := "1 " + string(buf[:n])
			_, err = c.Write([]byte(msg))

			if err != nil {
				log.Fatalln(err)
				return
			}
		}(conn)
		
	}
}

func Login(ctx context.Context, cmd *cli.Command) error {
	return nil
}

func Exit(ctx context.Context, cmd *cli.Command) error {
	return nil
}