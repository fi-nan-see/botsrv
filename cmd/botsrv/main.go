package main

import (
	"io"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"botsrv/pkg/app"

	"github.com/BurntSushi/toml"
	"github.com/namsral/flag"
)

var (
	fs           = flag.NewFlagSetWithEnvPrefix(os.Args[0], "BOTSRV", 0)
	flConfigPath = fs.String("config", "cfg/local.toml", "Path to config file")
	flVerbose    = fs.Bool("verbose", true, "enable debug output")

	cfg app.Config
)

func main() {
	rand.NewSource(time.Now().UnixNano())
	flag.DefaultConfigFlagname = "config.flag"
	exitOnError(fs.Parse(os.Args[1:]))
	fixStdLog(*flVerbose)

	_, err := toml.DecodeFile(*flConfigPath, &cfg)
	exitOnError(err)
	//TODO: cfg.Bot.Token = *token

	application, err := app.New("botsrv", *flVerbose, cfg, os.Getenv("tgSalt"))
	exitOnError(err)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	// Run
	go func() {
		if err := application.Run(); err != nil {
			exitOnError(err)
		}
	}()
	<-quit
}

func exitOnError(err error) {
	if err != nil {
		log.SetOutput(os.Stderr)
		log.Fatal(err)
	}
}

// fixStdLog sets additional params to std logger (prefix D, filename & line).
func fixStdLog(verbose bool) {
	log.SetPrefix("D")
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	if verbose {
		log.SetOutput(os.Stdout)
	} else {
		log.SetOutput(io.Discard)
	}
}
