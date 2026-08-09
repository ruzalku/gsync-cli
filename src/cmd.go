package src


import (
	"github.com/urfave/cli/v3"
	"gsynccli/src/commands"
)


func GetCMD() *cli.Command {
	cmd := &cli.Command{
		Commands: []*cli.Command{
			{
				Name: "start",
				Usage: "Starts programm",
				Action: commands.Start,
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name: "quiet",
						Aliases: []string{"q"},
						Usage: "Quiet mode",
					},
				},
			},
			{
				Name: "save",
				Usage: "Save file to Google Drive",
				Action: commands.SaveFile,
				Arguments: []cli.Argument{
					&cli.StringArg{
						Name: "path",
						UsageText: "Path of file",
					},
				},
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name: "autosave",
						Aliases: []string{"a"},
						Usage: "Enable autosaving file to Google Drive",
					},
				},
			},
			{
				Name: "exit",
				Usage: "Close programm",
				Action: commands.Exit,
			},
			{
				Name: "list",
				Usage: "List of connected to Google Drive files",
				Action: commands.GetListFiles,
			},
		},
	}

	return cmd
}
