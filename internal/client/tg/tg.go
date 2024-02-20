package tg

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"

	"pills-bot/internal/pkg/logger/logger"
)

type HandlerFunc func(tgUpdate tgbotapi.Update, c *Client)

type Client struct {
	client         *tgbotapi.BotAPI
	processMessage HandlerFunc // Функция обработки входящих сообщений.
}

func New(token string, handlerProcessingFunc HandlerFunc) (*Client, error) {
	client, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, errors.Wrap(err, "Ошибка NewBotAPI")
	}

	return &Client{
		client:         client,
		processMessage: handlerProcessingFunc,
	}, nil
}

func (c *Client) ListenUpdates() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := c.client.GetUpdatesChan(u)

	logger.Info("Start listening for tg messages")

	for update := range updates {
		// Функция обработки сообщений (обернутая в middleware).
		c.processMessage(update, c)
	}
}

func ProcessMessage(tgUpdate tgbotapi.Update, c *Client) {
	logger.Info(tgUpdate.Message.Text)
}
