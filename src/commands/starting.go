package commands

import (
	"log"
	"time"
	"context"
	"io"

	"github.com/urfave/cli/v3"
)


func Start(ctx context.Context, cmd *cli.Command) error {
	if cmd.Bool("quiet") {
		log.SetOutput(io.Discard)
	}
	go func() {
		for {
			time.Sleep(time.Second * 2)
			log.Print("gsynccli started")
		}
	}()
	log.Print("START")
	select{}
}

func Exit(ctx context.Context, cmd *cli.Command) error {
	return nil
}