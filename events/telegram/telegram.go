package telegram

import "mybot/clients/telegram"

type TelegramProcessor struct {
	tg     *telegram.Client
	offset int
	// storage
}

func New(client *telegram.Client)
