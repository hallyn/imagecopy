package main

import (
	"fmt"
	"os"

	"github.com/apex/log"
	"github.com/urfave/cli"
	"stackerbuild.io/stacker/pkg/lib"
)

func main() {
	app := cli.NewApp()
	app.Name = "image"
	app.Version = "1"
	app.Commands = []cli.Command{
		copyCmd,
	}
	app.Flags = []cli.Flag{
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatalf("%v\n", err)
	}
}

var copyCmd = cli.Command{
	Name:   "copy",
	Usage:  "copy an oci image",
	Action: doCopy,
}

func doCopy(ctx *cli.Context) error {
	args := ctx.Args()
	if len(args) != 2 {
		return fmt.Errorf("Usage: image copy src dest")
	}
	copyOpts := lib.ImageCopyOpts{
		Src:         args[0],
		Dest:        args[1],
		Progress:    os.Stdout,
		SrcSkipTLS:  true,
		DestSkipTLS: true,
	}
	if err := lib.ImageCopy(copyOpts); err != nil {
		return fmt.Errorf("Failed copying %s to %s: %w", args[0], args[1], err)
	}
	return nil
}
