package handlers

import (
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Common helper functions that are used by the handlers.

// GetUsername returns the username of a Telegram user. If the user has a username,
// it is returned with a "@" prefix. If not, the user's first and last name are
// concatenated and returned as the username.
func GetUsername(m *tgbotapi.User) string {
	if m.UserName != "" {
		return "@" + m.UserName
	}

	username := ""
	username = m.FirstName
	if m.LastName != "" {
		username = username + " " + m.LastName
	}

	return username
}

// DeleteMessage deletes a message by sending a request to the Telegram API
// with a DeleteMessage command. It sleeps for DeleteMessageTime before making the request
// to allow for a delay in message delivery. Returns an error if the request fails.
func (h *Handlers) DeleteMessage(chatID int64, messageID int) error {
	time.Sleep(DeleteMessageTime)

	delMsg := tgbotapi.NewDeleteMessage(chatID, messageID)
	_, err := h.bot.Request(delMsg)
	if err != nil {
		h.logger.Errorf("Error deleting message: %v", err)
		return err
	}

	return nil
}
