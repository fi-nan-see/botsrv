package handler

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"

	"botsrv/pkg/finansee"

	"botsrv/pkg/client/apisrv"
	"botsrv/pkg/embedlog"
)

type SendMsgParams struct {
	IsDocument      bool
	IsContact       bool
	IsMenuButton    bool
	FirstName       string
	PhoneNumber     string
	ChatID          int64
	Text            string
	BotName         string
	InputFileString []string
	MenuButton      finansee.MenuButtonParams
	ReplyMarkup     finansee.ReplyMarkupParams
}

type BotService struct {
	embedlog.Logger
	Ac      *apisrv.Client
	isDevel bool
}

func NewBotService(logger embedlog.Logger, ac *apisrv.Client, isDevel bool) *BotService {
	return &BotService{Logger: logger, Ac: ac, isDevel: isDevel}
}

func (bs *BotService) StartHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	bs.SendMessage(ctx, b, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Добро пожаловать в бот Finansee! Здесь вы можете осуществлять финансовое планирование, не выходя из Telegram!",
	})
}
