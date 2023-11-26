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
	tgSalt       string
}

func New(appName string, verbose bool, cfg Config, tgSalt string) (*App, error) {
	a := &App{
		appName:      appName,
		cfg:          cfg,
		apisrvClient: apisrv.NewDefaultClient(cfg.Server.Api),
		tgSalt:       tgSalt,
	}
	a.SetStdLoggers(verbose)

	a.bs = handler.NewBotService(a.Logger, a.apisrvClient, a.cfg.Server.IsDevel, a.tgSalt)
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
		bot.WithCallbackQueryDataHandler("add_income_", bot.MatchTypePrefix, a.bs.AddIncomeHandler),
		bot.WithCallbackQueryDataHandler("add_outcome_", bot.MatchTypePrefix, a.bs.AddOutcomeHandler),
		bot.WithCallbackQueryDataHandler("add_savings_", bot.MatchTypePrefix, a.bs.AddSavingsHandler),
		bot.WithCallbackQueryDataHandler("add_plan_income_", bot.MatchTypePrefix, a.bs.AddPlanIncomeHandler),
		bot.WithCallbackQueryDataHandler("add_plan_outcome_", bot.MatchTypePrefix, a.bs.AddPlanOutcomeHandler),
		bot.WithCallbackQueryDataHandler("add_plan_savings_", bot.MatchTypePrefix, a.bs.AddPlanSavingsHandler),
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
