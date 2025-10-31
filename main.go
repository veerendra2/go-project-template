package main

import (
	"log/slog"
	"os"

	"github.com/alecthomas/kong"
	"github.com/veerendra2/gopackages/slogger"
	"github.com/veerendra2/gopackages/version"
)

var cli struct {
	Log slogger.Config `embed:"" prefix:"log." envprefix:"LOG_"`

	Version bool `help:"Show app version and exit"`
}

func main() {
	kongCtx := kong.Parse(&cli,
		kong.Description("My app."),
	)

	if cli.Version {
		println(version.Version)
		os.Exit(0)
	}

	kongCtx.FatalIfErrorf(kongCtx.Error)

	slog.SetDefault(slogger.New(cli.Log))

	slog.Info("Version information", version.Info())
	slog.Info("Build context", version.BuildContext())
}
