package main

import (
	"log"
	"net/http"
	"os"
	"wolfbot/domain/model"
	"wolfbot/handler"
	"wolfbot/initializer"

	"github.com/joho/godotenv"
	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {
	env := os.Getenv("APP_ENV")

	if env == "local" {
		godotenv.Load(".env.local")
	}

	channelSecret := mustGetenv("CHANNEL_SECRET")
	channelAccessToken := mustGetenv("CHANNEL_ACCESS_TOKEN")
	port := mustGetenv("PORT")
	dbURL := mustGetenv("DATABASE_URL")

	_, service := initializer.Initialize(dbURL)

	messageHandler := handler.NewMessageHandler(service.VillageService)

	bot, err := linebot.New(channelSecret, channelAccessToken)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/callback", func(w http.ResponseWriter, req *http.Request) {
		events, err := bot.ParseRequest(req)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				w.WriteHeader(400)
			} else {
				w.WriteHeader(500)
			}
			return
		}

		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					reply, err := messageHandler.HandleGroupMessage(
						message.Text,
						model.UserID(event.Source.UserID),
						model.GroupID(event.Source.GroupID),
					)
					if err != nil {
						log.Println(err)
					}

					if _, err = bot.ReplyMessage(
						event.ReplyToken,
						linebot.NewTextMessage(reply),
					).
						Do(); err != nil {
						log.Print(err)
					}
				}
			}
		}
	})

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func mustGetenv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatal("failed to read environment variable: " + key)
	}
	return val
}
