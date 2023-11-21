package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"botsrv/pkg/client/apisrv"
	"botsrv/pkg/embedlog"
	"botsrv/pkg/handler"

	"github.com/go-telegram/bot"
)

type Config struct {
	Server struct {
		Api     string
		IsDevel bool
	}
	Bot struct {
		Token string
	}
}

type App struct {
	embedlog.Logger
	appName      string
	cfg          Config
	apisrvClient *apisrv.Client
	bot          *bot.Bot
	bs           *handler.BotService
}

func New(appName string, verbose bool, cfg Config) (*App, error) {
	a := &App{
		appName:      appName,
		cfg:          cfg,
		apisrvClient: apisrv.NewDefaultClient(cfg.Server.Api),
	}
	a.SetStdLoggers(verbose)

	a.bs = handler.NewBotService(a.Logger, a.apisrvClient, a.cfg.Server.IsDevel)
	a.bs.SetLoggers(a.Loggers())

	// register bot
	if err := a.registerBot(); err != nil {
		return nil, fmt.Errorf("register bot failed: %w", err)
	}

	return a, nil
}

func (a *App) registerBot() error {
	opts := []bot.Option{
		bot.WithMessageTextHandler("/start", bot.MatchTypeExact, a.bs.StartHandler),
		bot.WithDefaultHandler(a.bs.InlineQueryHandler),
	}

	b, err := bot.New(a.cfg.Bot.Token, opts...)
	if err != nil {
		return err
	}
	a.bot = b
	return nil
}

func (a *App) Run() error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	a.bot.Start(ctx)
	return nil
}
