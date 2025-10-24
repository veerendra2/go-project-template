package main

import (
	"log/slog"

	"github.com/alecthomas/kong"
	"github.com/veerendra2/gopackages/slogger"
)

var cli struct {
	Log slogger.Config `embed:"" prefix:"log." envprefix:"LOG_"`
}

func main() {
	kongCtx := kong.Parse(&cli)
	kongCtx.FatalIfErrorf(kongCtx.Error)

	slog.SetDefault(slogger.New(cli.Log))
	slog.Info("Info log")
	slog.Warn("Warning log")
	slog.Debug("Debug log")
	slog.Error("Error log")
}
