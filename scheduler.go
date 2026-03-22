package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
	"go.mau.fi/whatsmeow/types"
)

func startScheduler() {

	c := cron.New()

	c.AddFunc("0 9 * * *", func() {

		jid, _ := types.ParseJID("123456789@s.whatsapp.net")

		sendMessage(jid, "🌞 Good morning!")

	})

	c.Start()
}

func autoReconnect() {

	fmt.Println("Disconnected. Reconnecting...")

	go func() {

		for {

			err := client.Connect()

			if err == nil {
				fmt.Println("Reconnected!")
				break
			}

			time.Sleep(5 * time.Second)
		}

	}()
}