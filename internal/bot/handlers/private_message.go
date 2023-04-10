package handlers

import (
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const WaitToReplyTime = time.Second * 2

// Private message handler that sends a predefined message sequence to the user who sent the private message.
func (h *Handlers) OnPrivateMessage(ctx *tgbotapi.Message) {
	h.logger.Infof("New private message from %s (%d)", ctx.From.UserName, ctx.From.ID)
	var msgText string

	msgText = "listen up. Just go ahead and add me to this chat, then make me the grand poobah, and give me the power to vaporize any spam-happy losers dumb enough to show their faces around here. Got it, "
	sendMsg := tgbotapi.NewMessage(ctx.Chat.ID, "Hey, "+ctx.From.UserName+", "+msgText+" "+ctx.From.UserName+"?")
	sendMsg.ReplyToMessageID = ctx.MessageID

	go h.sendPrivateMessage(sendMsg)

	time.Sleep(WaitToReplyTime)
	msgText = "I wasn't just born, I was engineered specifically for @advenadiem. 😉"
	secMsg := tgbotapi.NewMessage(ctx.Chat.ID, msgText)

	go h.sendPrivateMessage(secMsg)

	time.Sleep(WaitToReplyTime)
	msgText = "If you're in need of a bot for your group, just hit up @elicodes. He's got what you need. Wubba lubba dub dub! 🖕"
	thirdMsg := tgbotapi.NewMessage(ctx.Chat.ID, msgText)

	go h.sendPrivateMessage(thirdMsg)
}

func (h *Handlers) sendPrivateMessage(msgConfig tgbotapi.MessageConfig) {
	_, err := h.bot.Send(msgConfig)
	if err != nil {
		h.logger.Errorf("Error sending message: %v", err)
		return
	}
}
