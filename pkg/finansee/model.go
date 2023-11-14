package finansee

import (
	"github.com/go-telegram/bot/models"
)

//type CallbackDataParams struct {
//	ID                 string `json:"ID"`
//	Text               string `json:"-"`
//	OrderDeclineReason string `json:"OrderDeclineReason,omitempty"`
//	OrderID            int    `json:"OrderID,omitempty"`
//	OfferID            int    `json:"OfferID,omitempty"`
//	TgID               int    `json:"TgID,omitempty"`
//	ProductionID       int    `json:"ProductionID,omitempty"`
//	NewPage            int    `json:"NewPage,omitempty"`
//	CountProviders     int    `json:"CountProviders,omitempty"`
//	CountProds         int    `json:"CountProds,omitempty"`
//}

type WebAppButtonParams struct {
	Text      string
	WebAppUrl string
}
type MenuButtonParams struct {
	Type   string
	Text   string
	WebApp string
}
type ReplyMarkupParams struct {
	IsRemove     bool
	IsInline     bool
	IsWebApp     bool
	WebAppParams WebAppButtonParams
	//InlineParams []CallbackDataParams
}

//func NewCallbackDataParams(s string) (CallbackDataParams, error) {
//	var b CallbackDataParams
//	err := json.Unmarshal([]byte(s), &b)
//	return b, err
//}
//
//func (b CallbackDataParams) String() (string, error) {
//	s, err := json.Marshal(b)
//	return string(s), err
//}
//
//func (b CallbackDataParams) MustString() string {
//	s, err := b.String()
//	if err != nil {
//		panic(err)
//	}
//
//	return s
//}

func NewMenuButton(params MenuButtonParams) models.MenuButtonWebApp {
	return models.MenuButtonWebApp{
		Type:   params.Type,
		Text:   params.Text,
		WebApp: models.WebAppInfo{URL: params.WebApp},
	}
}
func NewReplyMarkup(params ReplyMarkupParams) models.ReplyMarkup {
	if params.IsInline {
		if params.IsWebApp {
			return models.InlineKeyboardMarkup{
				InlineKeyboard: [][]models.InlineKeyboardButton{{models.InlineKeyboardButton{
					Text:   params.WebAppParams.Text,
					WebApp: &models.WebAppInfo{URL: params.WebAppParams.WebAppUrl},
				}}},
			}
		} else {
			//markup := make([][]models.InlineKeyboardButton, len(params.InlineParams))
			//for i := range markup {
			//	markup[i] = make([]models.InlineKeyboardButton, 1)
			//}
			//for i, button := range params.InlineParams {
			//	markup[i][0] = models.InlineKeyboardButton{
			//		Text:         button.Text,
			//		CallbackData: button.MustString(),
			//	}
			//}
			//return models.InlineKeyboardMarkup{
			//	InlineKeyboard: markup,
			//}
		}
	} else {
		if params.IsWebApp {
			return models.ReplyKeyboardMarkup{
				Keyboard: [][]models.KeyboardButton{
					{models.KeyboardButton{
						Text:   params.WebAppParams.Text,
						WebApp: &models.WebAppInfo{URL: params.WebAppParams.WebAppUrl},
					}},
				},
				ResizeKeyboard: true,
			}
		} //else {
		//
		//}
	}
	if params.IsRemove {
		return models.ReplyKeyboardRemove{
			RemoveKeyboard: true,
		}
	}
	return models.ReplyKeyboardRemove{
		RemoveKeyboard: false,
	}
}
