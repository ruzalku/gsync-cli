package main

import (
	"context"
	"log"
	"os"

	"gsynccli/src"
)

func main() {
	cmd := src.GetCMD()

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

