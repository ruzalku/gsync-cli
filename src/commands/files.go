package commands

import (
	"context"
	"log"

	"github.com/urfave/cli/v3"
	_ "google.golang.org/api/docs/v1"
	_ "google.golang.org/api/sheets/v4"
	_ "google.golang.org/api/slides/v1"

	"gsynccli/src/utils"
)


func SaveFile(ctx context.Context, cmd *cli.Command) error {
	err := utils.AddCommand("/tmp/gsynccli.sock", cmd.StringArg("path"))
	if err != nil {
		return err
	}
	
	log.Println("Sucessfully saved")
	return nil
}


func GetListFiles(ctx context.Context, cmd *cli.Command) error {
	return nil
}