package main

import (
	"context"
	"strings"

	"go.mau.fi/whatsmeow/types"

	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types/events"
)

func handleCommand(v *events.Message) {

	text := strings.ToLower(strings.TrimSpace(v.Message.GetConversation()))

	if text == "" {
		return
	}

	sender := v.Info.Chat

	if isFirstTime(sender.String()) {
		sendMessage(sender, "Ketik /manual untuk list command")
	}

	switch {

	case text == "ping":
		sendMessage(sender, "🏓 Pong!")

	case text == "menu" || text == "Menu":
		sendMessage(sender,
			"Terminal O' WORLD\n\n"+
				"/\n"+
				"menu",
		)

	}

	messageLogger.Println(sender.String(), "|", text)
}

func sendMessage(jid types.JID, text string) {

	msg := &waProto.Message{
		Conversation: &text,
	}

	client.SendMessage(
		context.Background(),
		jid,
		msg,
	)

	messageLogger.Println("BOT ->", jid.String(), "|", text)
}