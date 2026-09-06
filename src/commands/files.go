package commands

import (
	"context"
	"log"

	"github.com/urfave/cli/v3"
	_ "google.golang.org/api/docs/v1"
	_ "google.golang.org/api/sheets/v4"
	_ "google.golang.org/api/slides/v1"

	"gsynccli/src/utils/ipc"
	"gsynccli/src/utils"
)


func SaveFile(ctx context.Context, cmd *cli.Command) error {
	normalizedPath, err := utils.NormalizePath(cmd.StringArg("path"))
	if err != nil {
		return err
	}
	command, err := ArgsToCommand(
		"save",
		normalizedPath,
		cmd.Bool("autosave"),
	)

	if err != nil {
		return err
	}

	log.Println(string(command))

	err = ipc.AddCommand(ipc.ConnectPath, command)
	if err != nil {
		return err
	}
	
	log.Println("Sucessfully saved")
	return nil
}


func GetListFiles(ctx context.Context, cmd *cli.Command) error {
	return nil
}