package lib

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"
	"time"

	tele "gopkg.in/telebot.v3"
	//https://github.com/go-telebot/telebot/tree/v3.3.6
)

var bot *tele.Bot
var err error

func TelegramBod(token string) {
	token = strings.TrimSuffix(token, "\n")
	pref := tele.Settings{
		Token:  token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	bot, err = tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}
	//bot.Use(loggingMiddleware)

	bot.Handle("/hello", func(c tele.Context) error {
		var user = c.Sender()
		//user.ID
		fmt.Printf("images = %v is of type %T \n", user.ID, user.ID)
		fmt.Printf("images = %v is of type %T \n", user.Username, user.Username)

		return c.Send("Hello!")
	})

	// Handle all text messages

	bot.Handle(tele.OnText, func(c tele.Context) error {
		// All the text messages that weren't
		// captured by existing handlers.

		var (
			user = c.Sender()
			text = c.Text()
		)
		fmt.Printf("id = %v is of type %T \n", user.ID, user.ID)
		fmt.Printf("Usernames = %v is of type %T \n", user.Usernames, user.Usernames)

		WriteToFile(text)

		_, err := bot.Send(user, "Understood !")
		if err != nil {
			return err
		}
		fmt.Printf("To Log:  %v \n", text)

		// Instead, prefer a context short-hand:
		return nil
	})
	bot.Handle(tele.OnAudio, func(c tele.Context) error {
		// Check if the message contains an audio file
		m := c.Message()
		if m.Audio != nil {
			fileID := m.Audio.FileID
			filePath := filepath.Join("./", fileID+".ogg")

			// Download the audio file
			//err = tele.Download(fileID, filePath)
			//if err != nil {
			//	log.Printf("Failed to download audio file: %v", err)
			//	return
			//}

			log.Printf("Audio file downloaded and saved to: %s", filePath)

		}
		return nil
	})

	//bot.Handle(tele.OnText, func(m *tele.Message) {
	//	// Process or respond to the message here
	//	log.Printf("Received message: %s", m.Text)
	//})

	//userID := int64(6065413981)
	//localUser := &tele.User{ID: userID}
	//_, err = bot.Send(localUser, "test")
	//if err != nil {
	//	log.Printf("Failed to send message to user %d: %v", userID, err)
	//}

	//bot.Use(middleware.Logger())
	//bot.Use(middleware.AutoRespond())
	go bot.Start()

	//bot.Send(kasper, "hoi ik ben online")

}

func SendMessageToUser(userID int64, message string) {
	user := &tele.User{ID: userID}
	_, err := bot.Send(user, message)
	if err != nil {
		log.Printf("Failed to send message to user %d: %v", userID, err)
	}
}

func loggingMiddleware(next tele.HandlerFunc) tele.HandlerFunc {
	return func(context tele.Context) error {
		fmt.Printf("context = %v is of type %T \n", context, context)

		if context.Callback() != nil {
			defer context.Respond()
		}
		return next(context) // continue execution chain
	}
}
