package handler

import (
	"context"
	"fmt"
	"strconv"
	"strings"

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
	tgSalt  string
	isDevel bool
}

func NewBotService(logger embedlog.Logger, ac *apisrv.Client, isDevel bool, tgSalt string) *BotService {
	return &BotService{Logger: logger, Ac: ac, isDevel: isDevel, tgSalt: tgSalt}
}

func (bs *BotService) StartHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	r := update.Message.Text
	parts := strings.SplitN(r, " ", 2)
	if len(parts) > 1 {
		res, err := bs.Ac.Plans.RegisterTgID(ctx, strconv.Itoa(int(update.Message.From.ID))+bs.tgSalt, parts[1])
		if err != nil {
			return
		} else if res {
			bs.SendMessage(ctx, b, &bot.SendMessageParams{
				ChatID: update.Message.Chat.ID,
				Text:   "Добро пожаловать в бот Finansee, вы успешно зарегистрированы!",
			})
		}
	} else {
		bs.SendMessage(ctx, b, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Добро пожаловать в бот Finansee! Зарегистрируйтесь через синюю кнопку FinanSee снизу слева, чтобы получить доступ к функциям бота!",
		})
	}
}

func (bs *BotService) InlineQueryHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.InlineQuery == nil || update.InlineQuery.Query == "" {
		return
	}
	if _, err := strconv.Atoi(update.InlineQuery.Query); err != nil {
		return
	}

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
							CallbackData: "add_outcome_" + update.InlineQuery.Query,
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

func (bs *BotService) AddIncomeHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	q := strings.Split(update.CallbackQuery.Data, "_")
	// TODO: add errors handling
	plans, err := bs.Ac.Plans.GetPlans(ctx, strconv.Itoa(int(update.CallbackQuery.Sender.ID))+bs.tgSalt)
	if err != nil {
		bs.Errorf("%v", err)
		return
	}
	kb := [][]models.InlineKeyboardButton{[]models.InlineKeyboardButton{}}
	for _, plan := range plans {
		kb = append(kb, []models.InlineKeyboardButton{
			{
				Text:         plan.Name,
				CallbackData: "add_plan_income_" + plan.ID + "_" + q[2]},
		})
	}
	_, err = b.EditMessageText(ctx, &bot.EditMessageTextParams{
		InlineMessageID: update.CallbackQuery.InlineMessageID,
		Text:            "Выберите план из списка ниже:",
		ReplyMarkup:     models.InlineKeyboardMarkup{InlineKeyboard: kb},
	})
}

func (bs *BotService) AddOutcomeHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	q := strings.Split(update.CallbackQuery.Data, "_")
	// TODO: add errors handling
	plans, err := bs.Ac.Plans.GetPlans(ctx, strconv.Itoa(int(update.CallbackQuery.Sender.ID))+bs.tgSalt)
	if err != nil {
		bs.Errorf("%v", err)
		return
	}
	kb := [][]models.InlineKeyboardButton{[]models.InlineKeyboardButton{}}
	for _, plan := range plans {
		kb = append(kb, []models.InlineKeyboardButton{
			{
				Text:         plan.Name,
				CallbackData: "add_plan_outcome_" + plan.ID + "_" + q[2]},
		})
	}
	_, err = b.EditMessageText(ctx, &bot.EditMessageTextParams{
		InlineMessageID: update.CallbackQuery.InlineMessageID,
		Text:            "Выберите план из списка ниже:",
		ReplyMarkup:     models.InlineKeyboardMarkup{InlineKeyboard: kb},
	})
	if err != nil {
		bs.Errorf("%v", err)
		return
	}
	//bs.EditMessage(ctx, b, &bot.EditMessageTextParams{
	//	InlineMessageID: update.CallbackQuery.InlineMessageID,
	//	Text:            "Выберите план из списка ниже:",
	//	ReplyMarkup:     models.InlineKeyboardMarkup{InlineKeyboard: kb},
	//})
}
func (bs *BotService) AddSavingsHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	q := strings.Split(update.CallbackQuery.Data, "_")
	// TODO: add errors handling
	plans, err := bs.Ac.Plans.GetPlans(ctx, strconv.Itoa(int(update.CallbackQuery.Sender.ID))+bs.tgSalt)
	if err != nil {
		bs.Errorf("%v", err)
		return
	}
	kb := [][]models.InlineKeyboardButton{[]models.InlineKeyboardButton{}}
	for _, plan := range plans {
		kb = append(kb, []models.InlineKeyboardButton{
			{
				Text:         plan.Name,
				CallbackData: "add_plan_savings_" + plan.ID + "_" + q[2]},
		})
	}
	_, err = b.EditMessageText(ctx, &bot.EditMessageTextParams{
		InlineMessageID: update.CallbackQuery.InlineMessageID,
		Text:            "Выберите план из списка ниже:",
		ReplyMarkup:     models.InlineKeyboardMarkup{InlineKeyboard: kb},
	})
}

func (bs *BotService) AddPlanIncomeHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	q := strings.Split(update.CallbackQuery.Data, "_")
	// TODO: add errors handling
	amount, _ := strconv.Atoi(q[4])
	err := bs.Ac.Plans.AddIncome(ctx, strconv.Itoa(int(update.CallbackQuery.Sender.ID))+bs.tgSalt, q[3], float64(amount))
	if err != nil {
		bs.Errorf("%v", err)
		return
	}
	bs.EditMessage(ctx, b, &bot.EditMessageTextParams{
		InlineMessageID: update.CallbackQuery.InlineMessageID,
		Text:            fmt.Sprintf("income to plan with summ %d successfully added", amount),
	})
}

func (bs *BotService) AddPlanOutcomeHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	q := strings.Split(update.CallbackQuery.Data, "_")
	// TODO: add errors handling
	amount, _ := strconv.Atoi(q[4])
	err := bs.Ac.Plans.AddOutcome(ctx, strconv.Itoa(int(update.CallbackQuery.Sender.ID))+bs.tgSalt, q[3], float64(amount))
	if err != nil {
		bs.Errorf("%v", err)
		return
	}
	bs.EditMessage(ctx, b, &bot.EditMessageTextParams{
		InlineMessageID: update.CallbackQuery.InlineMessageID,
		Text:            fmt.Sprintf("outcome to plan with summ %d successfully added", amount),
	})
}

func (bs *BotService) AddPlanSavingsHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	q := strings.Split(update.CallbackQuery.Data, "_")
	// TODO: add errors handling
	amount, _ := strconv.Atoi(q[4])
	err := bs.Ac.Plans.AddSavings(ctx, strconv.Itoa(int(update.CallbackQuery.Sender.ID))+bs.tgSalt, q[3], float64(amount))
	if err != nil {
		bs.Errorf("%v", err)
		return
	}
	bs.EditMessage(ctx, b, &bot.EditMessageTextParams{
		InlineMessageID: update.CallbackQuery.InlineMessageID,
		Text:            fmt.Sprintf("savings to plan with summ %d successfully added", amount),
	})
}

func (bs *BotService) PlansHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	plans, err := bs.Ac.Plans.GetPlans(ctx, strconv.Itoa(int(update.Message.From.ID))+bs.tgSalt)
	if err != nil {
		bs.Errorf("%v", err)
		return
	}
	kb := [][]models.InlineKeyboardButton{[]models.InlineKeyboardButton{}}
	for _, plan := range plans {
		kb = append(kb, []models.InlineKeyboardButton{
			{
				Text:         plan.Name,
				CallbackData: "get_plan_" + plan.ID},
		})
	}
	bs.SendMessage(ctx, b, &bot.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "Выберите план из списка ниже:",
		ReplyMarkup: models.InlineKeyboardMarkup{InlineKeyboard: kb},
	})
}

func (bs *BotService) GetPlanHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	q := strings.Split(update.CallbackQuery.Data, "_")
	plan, err := bs.Ac.Plans.GetPlan(ctx, strconv.Itoa(int(update.CallbackQuery.Sender.ID))+bs.tgSalt, q[2])
	if err != nil {
		return
	}
	text := fmt.Sprintf("Название: %s\nДата начала: %s\nДата окончания: %s\nИзначальный баланс: %d\nТекущий баланс: %d\nАктуальный: %v", plan.Name, plan.Start_date, plan.End_date, plan.Initial_balance, plan.Current_balance, plan.Is_actual)
	bs.SendMessage(ctx, b, &bot.SendMessageParams{
		ChatID: update.CallbackQuery.Message.Chat.ID,
		Text:   text,
	})
}
