package handlers

import (
	"fmt"
	"math/rand"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// OnMemberLeft handles the event when a member leaves the chat. It gets the chat ID
// and the username of the member who left, logs the event, sends a message to the chat
// to notify other members of the departure, and then deletes the notification message.
func (h *Handlers) OnMemberLeft(ctx *tgbotapi.Message) {
	chatID := ctx.Chat.ID
	username := GetUsername(ctx.LeftChatMember)

	h.logger.Infof("Member left chat: %s (%d)", username, ctx.LeftChatMember.ID)
	msg := tgbotapi.NewMessage(chatID, getMemberLeftText(username))

	sentMsg, err := h.bot.Send(msg)
	if err != nil {
		h.logger.Errorf("Error sending message: %v", err)
		return
	}

	go h.DeleteMessage(chatID, sentMsg.MessageID)
}

func getMemberLeftText(username string) string {
	var messages = []string{
		"Goodbye, %s. Don't be a Jerry. You'll just end up with a boring life and a loveless marriage.",
		"I'll miss you, %s. But I'll miss those Szechuan nugget sauce even more. Damn, those were good.",
		"Farewell, %s. Don't forget to bury your alternate selves properly. We wouldn't want any more Cronenbergs running around.",
		"Goodbye, %s. You better not clone yourself without my permission again. It's getting old, and we don't need any more Ricks in this universe.",
		"See you later, %s. Hope you find your way back from the Cronenberg dimension. That place is a real nightmare.",
		"Adios, %s. Keep your portal gun safe, and don't let Summer use it again. We don't need any more portal gun accidents.",
		"Goodbye, %s. Remember, don't trust aliens with sexy butts. They're usually up to no good.",
		"So long, %s. May your adventures never be boring. I mean, who wants to live a mundane life?",
		"Sayonara, %s. Don't let the Meeseeks take over your life. Trust me, it's not worth it.",
		"Peace out, %s. Try not to create any more apocalyptic timelines, okay? It's a real pain to clean up after those.",
		"Goodbye, %s. Don't forget to feed your sentient poop before you flush it down the toilet. It's a living creature, %s!",
		"Farewell, %s. Don't forget to pay your interdimensional cable bill. I don't want to lose my subscription again.",
		"Adieu, %s. Remember, you're not allowed to join any cults without my permission. I mean, come on, %s. Use your head.",
		"Goodbye, %s. Try not to turn into a pickle again. That was a weird one, even for me.",
		"See you later, %s. Don't let your math teacher turn you into a giant jellybean. That's just messed up.",
		"Sayonara, %s. May your alcoholism never get the best of you. But let's be real, it probably will.",
		"Goodbye, %s. Remember, you're not allowed to use the Meeseeks box for personal gain. That's just selfish, %s.",
		"So long, %s. Hope you never have to fight the Council of Ricks again. Those guys are a real pain in the ass.",
		"Peace out, %s. Don't let the Cromulons judge you too harshly. They're a fickle bunch.",
		"Goodbye, %s. Don't forget to bring back some cool souvenirs from your next interdimensional trip. I'm expecting something really weird, %s.",
	}

	rand.Seed(time.Now().UnixNano())
	randMsg := messages[rand.Intn(len(messages))]

	return fmt.Sprintf(randMsg, username)
}
