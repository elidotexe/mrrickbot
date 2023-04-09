package bot

import (
	"github.com/elidotexe/mrrickbot/internal/bot/handlers"
	"github.com/elidotexe/mrrickbot/internal/logger"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	zaplog "go.uber.org/zap"
)

type Bot struct {
	bot      *tgbotapi.BotAPI
	logger   *logger.Logger
	updates  tgbotapi.UpdatesChannel
	handlers *handlers.Handlers
}

// NewBot creates a new instance of the bot, initializes it with the provided token,
// logger and message handlers, and returns the bot along with an error (if any).
func NewBot(token string, logger *logger.Logger) (*Bot, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	tgbotapi.NewUpdate(0)

	logger.Info("Authorized on Telegram", zaplog.String("bot", bot.Self.UserName))

	bot.Debug = false

	updates := bot.GetUpdatesChan(tgbotapi.UpdateConfig{
		Offset:  0,
		Timeout: 10,
	})

	h, err := handlers.Initialize(bot, logger)
	if err != nil {
		logger.Error("Error initializing handlers", zaplog.Error(err))
		return nil, err
	}

	b := &Bot{
		bot:      bot,
		logger:   logger,
		updates:  updates,
		handlers: h,
	}

	return b, nil
}

// Start starts the bot and listens for incoming messages. It uses the bot's
// handlers to route the incoming messages to their appropriate functions. If the
// message is of an unknown type, it logs an error message.
func (b *Bot) Start() error {
	b.logger.Info("Bot has been successfully started...")

	for u := range b.updates {
		switch {
		case u.Message == nil:
			continue
		case u.Message.NewChatMembers != nil:
			b.handlers.OnMemberJoined(u.Message)
		case u.Message.LeftChatMember != nil:
			b.handlers.OnMemberLeft(u.Message)
		case u.Message.Chat.IsPrivate():
			b.handlers.OnPrivateMessage(u.Message)
		case u.Message.Text != "":
			b.handlers.OnGroupMessage(u.Message)
		default:
			b.logger.Infof("Unknown message type: %T", u.Message)
		}
	}

	return nil
}
