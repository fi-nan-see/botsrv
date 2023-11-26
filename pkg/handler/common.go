package handler

import (
	"context"
	"strings"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// SetChatMenuButton is a function that sends chat menu button to TG API
func (bs *BotService) SetChatMenuButton(ctx context.Context, b *bot.Bot, chatId int64, params *bot.SetChatMenuButtonParams) {
	_, err := b.SetChatMenuButton(ctx, params)
	if err != nil {
		bs.Errorf("[SetChatMenuButton] failed to send chat menu button bot=%v err=%v", b, err)
		return
	}
	if bs.isDevel {
		bs.Printf("Send chat menu button with params %v", params)
	}
}

// SendContact is a function that sends contact to TG API
func (bs *BotService) SendContact(ctx context.Context, b *bot.Bot, chatId int64, params *bot.SendContactParams) {
	_, err := b.SendContact(ctx, params)
	if err != nil {
		bs.Errorf("[SendContact] failed to send contact bot=%v err=%v", b, err)
		return
	}
	if bs.isDevel {
		bs.Printf("Sent contact with params %v", params)
	}
}

// SendMessage is a function that sends message to TG API. It also supports testMode
func (bs *BotService) SendMessage(ctx context.Context, b *bot.Bot, params *bot.SendMessageParams) {
	params.ParseMode = models.ParseModeHTML
	_, err := b.SendMessage(ctx, params)
	if err != nil {
		bs.Errorf("[SendMessage] failed to send message bot=%v message=%v err=%v", b, params.Text, err)
		return
	}
	if bs.isDevel {
		bs.Printf("Sent message with params %v", params)
	}

}

// SendDocument is a function that sends document to TG API.
func (bs *BotService) SendDocument(ctx context.Context, b *bot.Bot, chatId int64, params *bot.SendDocumentParams) {
	params.ParseMode = models.ParseModeHTML
	_, err := b.SendDocument(ctx, params)
	if err != nil {
		bs.Errorf("[SendDocument] failed to send document bot=%v document=%v err=%v", b, params.Document, err)
		return
	}
}

// SendMediaGroup is a function that sends document to TG API.
func (bs *BotService) SendMediaGroup(ctx context.Context, b *bot.Bot, chatId int64, inputFileData []string) {
	var media []models.InputMedia
	var text []string
	for _, drawing := range inputFileData {
		if strings.HasSuffix(drawing, ".doc") || strings.HasSuffix(drawing, ".docx") || strings.HasSuffix(drawing, ".xlsx") {
			text = append(text, drawing)
		} else {
			media = append(media, &models.InputMediaDocument{Media: drawing, ParseMode: models.ParseModeHTML})
		}
	}
	if text != nil {
		bs.SendMessage(ctx, b, &bot.SendMessageParams{
			ChatID: chatId,
			Text:   strings.Join(text, "\n"),
		})
	}
	if media != nil {
		_, err := b.SendMediaGroup(ctx, &bot.SendMediaGroupParams{
			ChatID: chatId,
			Media:  media,
		})
		if err != nil {
			bs.Errorf("[SendMediaGroup] failed to send media bot=%v document=%v err=%v", b, media, err)
			return
		}
	}
}

// DeleteMessage is a function that sends delete message request to TG API.
func (bs *BotService) DeleteMessage(ctx context.Context, b *bot.Bot, chatId int64, params *bot.DeleteMessageParams) {
	_, err := b.DeleteMessage(ctx, params)
	if err != nil {
		bs.Errorf("[DeleteMessage] failed to delete message bot=%v messageId=%v err=%v", b, params.MessageID, err)
		return
	}
}

// EditMessage is a function that sends edit message request to TG API.
func (bs *BotService) EditMessage(ctx context.Context, b *bot.Bot, params *bot.EditMessageTextParams) {
	_, err := b.EditMessageText(ctx, params)
	params.ParseMode = models.ParseModeHTML
	if err != nil {
		bs.Errorf("[EditMessage] failed to edit message bot=%v messageId=%v messageText=%v inlineMessageId=%v err=%v", b, params.MessageID, params.Text, params.InlineMessageID, err)
		return
	}
	if bs.isDevel {
		bs.Printf("Edited message with params %v", params)
	}
}
