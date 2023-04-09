package handlers

import (
	"math/rand"
	"regexp"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// OnGroupMessage which handles group messages received by the Telegram bot. The function first 
// checks if the user sending the message is an admin or not. If the user is not an admin and 
// the message contains certain keywords related to drugs, the function deletes the message 
// and sends an anti-spam message in response.
func (h *Handlers) OnGroupMessage(ctx *tgbotapi.Message) {
	chatID := ctx.Chat.ID

	admins, _ := h.getChatAdmins(ctx)

	userIsAdmin := false
	for _, admin := range admins {
		if ctx.From.ID == admin.User.ID {
			userIsAdmin = true
			break
		}
	}

	pattern := regexp.MustCompile(`(?i)\b(cocaine|pills|meth|ecstasy|plugs|pingers|k|shrooms|mushrooms|dealer|dealers|f2f|deliveries|2cb|lsd|weed|md|mdma|xans|xanax|bud|ketamine|speed|adderall|acid|coke|hash|ket|tabs|dmt|xtc)\b`)

	if !userIsAdmin && pattern.MatchString(ctx.Text) {
		deleteMsg := tgbotapi.NewDeleteMessage(chatID, ctx.MessageID)

		_, err := h.bot.Request(deleteMsg)
		if err != nil {
			h.logger.Errorf("Error deleting message: %v", err)
			return
		}

		msg := tgbotapi.NewMessage(chatID, h.getAntiSpamMessageText())

		sentMsg, err := h.bot.Send(msg)
		if err != nil {
			h.logger.Errorf("Error sending message: %v", err)
			return
		}

	  go h.DeleteMessage(chatID, sentMsg.MessageID)
	}
}

func (h *Handlers) getChatAdmins(message *tgbotapi.Message) ([]tgbotapi.ChatMember, error) {
	adminConfig := tgbotapi.ChatAdministratorsConfig{
		ChatConfig: tgbotapi.ChatConfig{
			ChatID:             message.Chat.ID,
			SuperGroupUsername: message.Chat.ChatConfig().SuperGroupUsername,
		},
	}

	admins, err := h.bot.GetChatAdministrators(adminConfig)
	if err != nil {
		h.logger.Errorf("Error getting chat admins: %v", err)
		return nil, err
	}

	return admins, nil
}

func (h *Handlers) getAntiSpamMessageText() string {
	var messages = []string{
		"Hey, pal, do you have any idea how much trouble you're causing by promoting drugs? You're messing with people's lives, and that's not cool.",
		"Listen, pal, drugs are bad news. They'll mess you up worse than a faulty portal gun. Quit promoting that garbage before someone gets hurt.",
		"Wubba lubba dub dub! I don't care how much you like your drugs, promoting them is a recipe for disaster. Knock it off.",
		"Look, pal, I don't want to hear any more about your drug promotion scheme. You're playing with fire, and someone's gonna get burned.",
		"Summer, you need to stop promoting drugs right now. They're not a joke, they're not cool, and they're not something we want to be associated with.",
		"Hey, Tammy, how about you put a cork in that drug promotion nonsense? It's not worth the trouble it's causing",
		"Geez, pal, promoting drugs? That's lower than a Mr. Meeseeks stuck in a never-ending cycle. Cut it out",
		"Listen, Advena Diem fans, we don't tolerate drug promotion in this community. Keep it up and you'll be booted faster than a Plutonian pimple.",
		"Pal, you're promoting drugs? You're about as smart as a Cromulon trying to take on a Level 9 planet. Knock it off.",
		"Hey, Summer, don't be a Glorzo. Promoting drugs is dangerous and irresponsible. Quit it before it's too late.",
	}

	rand.Seed(time.Now().UnixNano())
	randMsg := messages[rand.Intn(len(messages))]

	return randMsg
}
