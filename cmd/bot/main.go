package main

import (
	"pills-bot/internal/client/tg"
	"pills-bot/internal/config"
	"pills-bot/internal/pkg/logger/logger"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		logger.Fatal("Ошибка получения файла конфигурации:", "err", err)
	}

	tgClient, err := tg.New(cfg.GetToken(), tg.ProcessMessage)
	if err != nil {
		logger.Fatal("Ошибка инициализации ТГ-клиента:", "err", err)
	}

	tgClient.ListenUpdates()
}
