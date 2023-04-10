package handlers

import (
	"fmt"
	"math/rand"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// OnMemberJoined handles the event when a member joins a group chat. It retrieves
// the chat ID and the username of the new member, sends a welcome message to
// the chat, and deletes the message after it is sent.
func (h *Handlers) OnMemberJoined(ctx *tgbotapi.Message) {
	for _, member := range ctx.NewChatMembers {
		chatID := ctx.Chat.ID
		username := GetUsername(&member)

		h.logger.Infof("New member joined chat: %s (%d)", username, member.ID)
		msg := tgbotapi.NewMessage(chatID, getNewMemberWelcomeText(username))

		sentMsg, err := h.bot.Send(msg)
		if err != nil {
			h.logger.Errorf("Error sending message: %v", err)
			return
		}

		go h.DeleteMessage(chatID, sentMsg.MessageID)
	}
}

func getNewMemberWelcomeText(username string) string {
	var messages = []string{
		"Wubba Lubba Dub Dub, what's up, %s?",
		"%s, Boom! Big reveal! I turned myself into a pickle!",
		"%s, Nobody exists on purpose. Nobody belongs anywhere. We’re all going to die. Come watch TV.",
		"Weddings are basically funerals with a cake, so what you want? %s",
		"Don’t move, %s. Gonorrhea can’t see us if we don’t move. Wait! I was wrong! I was thinking of a T. Rex",
		"You gotta do it for Grandpa, %s. You gotta put these seeds inside your butt.",
		"%s, your parents are a bag of dicks.",
		"Bitch, my generation gets traumatized for breakfast! So what you want, %s?",
		"%s, life is tekno and I’ll stop when I die!",
		"%s, what are you doing? No time for school let's goo to boobworld and get high on milk",
		"Greet extra or to change with other one that I don't remember, %s",
		"Bluuuurpp!! SUUUP %s where you've been? On the moon? Ahhhrr fuck it",
		"%s you don't  know me!! I can be a robot, I can be a hologram, I can be a robot hologram shagging your sister. Bluuurp, but for now I'm a pickle",
		"%s, life is meaningless but you can listen to Advena's tekno, kicks like the galactic cocaine.",
		"Hello, %s. What, did you get lost on the way to the unemployment office?",
		"Well, well, well, look who it is, %s. Hello, Principal Vagina. Still struggling to assert your dominance over teenagers, I see.",
		"Oh, hello, Summer %s. Did you finally come to terms with the fact that you'll never be as smart as your grandpa?",
		"Hey there, %s. I see you're still living in the shadow of your father's intelligence. It must be tough to be a constant disappointment.",
		"Well, hello there, Squanchy. Do you ever get tired of being a walking stereotype, %s?",
		"Ah, Morty's teacher. Hello, Mr. %s. I hope you're ready to be psychologically scarred for life again.",
		"Oh, it's you, Jerry's parents. Hello, I guess. I hope you're not here to bore me with your mundane lives, %s!",
	}

	rand.Seed(time.Now().UnixNano())
	randMsg := messages[rand.Intn(len(messages))]

	return fmt.Sprintf(randMsg, username)
}
