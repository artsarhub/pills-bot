package config

import (
	"os"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"

	"pills-bot/internal/pkg/logger/logger"
)

const configFile = "data/config.yaml"

type Config struct {
	Token string `yaml:"token"`              // Токен бота в телеграме.
	PgDsn string `yaml:"ConnectionStringDB"` // Строка подключения в базе данных.
	//KafkaTopic  string   `yaml:"KafkaTopic"`         // Наименование топика Kafka.
	//BrokersList []string `yaml:"BrokersList"`        // Список адресов брокеров сообщений (адрес Kafka).
}

type Service struct {
	config Config
}

func New() (*Service, error) {
	s := &Service{}

	rawYAML, err := os.ReadFile(configFile)
	if err != nil {
		logger.Error("Ошибка reading config file", "err", err)
		return nil, errors.Wrap(err, "reading config file")
	}

	err = yaml.Unmarshal(rawYAML, &s.config)
	if err != nil {
		logger.Error("Ошибка parsing yaml", "err", err)
		return nil, errors.Wrap(err, "parsing yaml")
	}

	return s, nil
}

func (s *Service) GetToken() string {
	return s.config.Token
}
