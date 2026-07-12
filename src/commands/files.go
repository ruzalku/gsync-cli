package commands

import (
	"context"

	"github.com/urfave/cli/v3"
	_ "google.golang.org/api/docs/v1"
	_ "google.golang.org/api/slides/v1"
	_ "google.golang.org/api/sheets/v4"
)


func SaveFile(ctx context.Context, cmd *cli.Command) error {
	return nil
}


func GetListFiles(ctx context.Context, cmd *cli.Command) error {
	return nil
}