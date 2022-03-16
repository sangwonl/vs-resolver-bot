package main

import (
		tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
		"os"
		"fmt"
		"log"
		"math/rand"
		"strings"
)

type candidate struct {
	text string
	note string
}

type selection struct {
	text string
	note string
	idx int
}

func main() {
	botToken := os.Getenv("botToken")

	if botToken == "" {
		fmt.Println("No bot token")
		os.Exit(1)
	}

	bot, err := tgbotapi.NewBotAPI(botToken)

	if err != nil {
		log.Panic(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message.IsCommand() {
			switch update.Message.Command() {

			case "vs":
				candidates, err := textIntoCandidates(update.Message.CommandArguments())
				choice := chooseOne(candidates)
				 
				bot.Send(generateAnswerWithChoice(choice))
			}
		}

		if update.Message != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		}
	}
}

func textIntoCandidates(text string) ([]candidate, error) {
		
	return nil, nil
}

func chooseOne(candidates []candidate) selection {}

func generateAnswerWithChoice(choice selection) {
	switch rand.Intn(3) {
	case 0:
		return fmt.Sprintf("쉽네, \"%s\" 이지. no doubt", selection.text)
	case 1:
		return fmt.Sprintf("Just, \"%s\"", selection.text)
	case 2:
		return fmt.Sprintf("한번만 말한다. \"%s\" 가 답이다.", selection.text)
	}
}
