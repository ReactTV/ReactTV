package main

import (
	"server/cmd/commands"

	"github.com/alecthomas/kong"
)

func main() {
	cli := commands.DataMangementCLI{
		Globals: commands.Globals{},
	}

	ctx := kong.Parse(&cli,
		kong.Name(commands.Name),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{Compact: true}))
	err := ctx.Run(&cli.Globals)
	ctx.FatalIfErrorf(err)
}
