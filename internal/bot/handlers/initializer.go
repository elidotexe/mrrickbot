package handlers

import (
	"time"

	"github.com/elidotexe/mrrickbot/internal/logger"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const DeleteMessageTime = time.Minute * 5

type Handlers struct {
	bot    *tgbotapi.BotAPI
	logger *logger.Logger
}

// Initialize initializes a new instance of the Handlers struct with the provided BotAPI and Logger, 
// and returns it along with a nil error.
func Initialize(b *tgbotapi.BotAPI, logger *logger.Logger) (*Handlers, error) {
	return &Handlers{
		bot:    b,
		logger: logger,
	}, nil
}
