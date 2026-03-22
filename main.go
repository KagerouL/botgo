package main

import (
	"context"
	"database/sql"
	"os"

	qrterminal "github.com/mdp/qrterminal/v3"
	_ "github.com/mattn/go-sqlite3"

	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store/sqlstore"
	"go.mau.fi/whatsmeow/types/events"
	waLog "go.mau.fi/whatsmeow/util/log"
)

var client *whatsmeow.Client
var db *sql.DB

func main() {

	ctx := context.Background()

	// WhatsApp session database
	dbLog := waLog.Stdout("Database", "INFO", true)
	container, err := sqlstore.New(
		ctx,
		"sqlite3",
		"file:store.db?_foreign_keys=on",
		dbLog,
	)

	if err != nil {
		panic(err)
	}

	deviceStore, err := container.GetFirstDevice(ctx)
	if err != nil {
		panic(err)
	}

	clientLog := waLog.Stdout("Client", "INFO", true)
	client = whatsmeow.NewClient(deviceStore, clientLog)

	client.AddEventHandler(eventHandler)

	// QR login
	if client.Store.ID == nil {

		qrChan, _ := client.GetQRChannel(ctx)

		go func() {
			for evt := range qrChan {
				if evt.Event == "code" {
					qrterminal.GenerateHalfBlock(evt.Code, qrterminal.L, os.Stdout)
				}
			}
		}()
	}

	// connect WhatsApp
	err = client.Connect()
	if err != nil {
		panic(err)
	}

	// user database
	db, err = sql.Open("sqlite3", "users.db")
	if err != nil {
		panic(err)
	}

	initLogger()

	initUserDB()

	// start scheduler
	startScheduler()

	select {}
}

func eventHandler(evt interface{}) {

	switch v := evt.(type) {

	case *events.Message:
		handleCommand(v)

	case *events.Disconnected:
		autoReconnect()

	}
}