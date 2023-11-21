package handler

import (
	"context"
	"strconv"

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

func (bs *BotService) InlineQueryHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.InlineQuery == nil || update.InlineQuery.Query == "" {
		return
	}
	if _, err := strconv.Atoi(update.InlineQuery.Query); err != nil {
		return
	}
	//query, err := b.AnswerInlineQuery(ctx)
	//if err != nil {
	//	return
	//}
	//bs.SendMessage(ctx, b, &bot.SendMessageParams{
	//	ChatID: update.InlineQuery.From.ID,
	//	Text:   "Вы послали: " + update.InlineQuery.Query,
	//})

	_, err := b.AnswerInlineQuery(ctx, &bot.AnswerInlineQueryParams{
		InlineQueryID: update.InlineQuery.ID,
		Results: []models.InlineQueryResult{
			&models.InlineQueryResultArticle{
				ID:           "1",
				Title:        "добавить доход: " + update.InlineQuery.Query,
				ThumbnailURL: "https://data.nalog.ru/cdn/image/2283847/original.jpg",
				InputMessageContent: models.InputTextMessageContent{
					MessageText: "Вы уверены, что хотите добавить доход " + update.InlineQuery.Query,
				},
				ReplyMarkup: models.InlineKeyboardMarkup{
					InlineKeyboard: [][]models.InlineKeyboardButton{{
						models.InlineKeyboardButton{
							Text:         "Да",
							CallbackData: "add_income_" + update.InlineQuery.Query,
						},
						models.InlineKeyboardButton{
							Text:         "Нет",
							CallbackData: "cancel",
						},
					},
					},
				},
			},
			&models.InlineQueryResultArticle{
				ID:           "2",
				Title:        "добавить расход: " + update.InlineQuery.Query,
				ThumbnailURL: "https://telegra.ph/file/6171cd885a7f04704bf16.jpg",
				InputMessageContent: models.InputTextMessageContent{
					MessageText: "Вы уверены, что хотите добавить расход " + update.InlineQuery.Query,
				},
				ReplyMarkup: models.InlineKeyboardMarkup{
					InlineKeyboard: [][]models.InlineKeyboardButton{{
						models.InlineKeyboardButton{
							Text:         "Да",
							CallbackData: "add_expense_" + update.InlineQuery.Query,
						},
						models.InlineKeyboardButton{
							Text:         "Нет",
							CallbackData: "cancel",
						},
					},
					},
				},
			},
			&models.InlineQueryResultArticle{
				ID:           "3",
				Title:        "добавить сбережения: " + update.InlineQuery.Query,
				ThumbnailURL: "https://goodtrading.ru/wp-content/uploads/2023/10/a1e414a6518db29da6c70a985a253bd9.jpeg",
				InputMessageContent: models.InputTextMessageContent{
					MessageText: "Вы уверены, что хотите добавить сбережения " + update.InlineQuery.Query,
				},
				ReplyMarkup: models.InlineKeyboardMarkup{
					InlineKeyboard: [][]models.InlineKeyboardButton{{
						models.InlineKeyboardButton{
							Text:         "Да",
							CallbackData: "add_savings_" + update.InlineQuery.Query,
						},
						models.InlineKeyboardButton{
							Text:         "Нет",
							CallbackData: "cancel",
						},
					},
					},
				},
			},
		},
		CacheTime:  0,
		IsPersonal: false,
		NextOffset: "",
		Button:     nil,
	})
	if err != nil {
		return
	}
}
