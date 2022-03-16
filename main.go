package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Candidate struct {
	text string
	note string
}

type Selection struct {
	text string
	note string
	idx  int
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
		if update.Message == nil {
			continue
		}
		if update.Message.IsCommand() {
			switch update.Message.Command() {

			case "saysome":
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, saysome()))
			case "vs":
				candidates, err := textIntoCandidates(update.Message.CommandArguments())

				if err != nil {
					bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Usage /vs sth vs sth"))
					continue
				}

				choice := chooseOne(candidates)

				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, generateAnswerWithChoice(choice)))
			}
		}

		if update.Message != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		}
	}
}

func saysome() string {
	i := rand.Intn(5)

	switch i {
	case 0:
		return "뭐.. 쯥.."
	case 1:
		return "ㅎㅎ"
	case 2:
		return ".. 딱히 ㅎㅎ"
	case 3:
		return "어 그랴"
	case 4:
		return "^.^"
	}

	return "ㅇㅇ"
}

func textIntoCandidates(text string) ([]Candidate, error) {
	chunks := strings.Split(text, "vs")

	if len(chunks) <= 1 {
		return nil, errors.New("no delimiter")
	}

	var candidates []Candidate

	for _, c := range chunks {
		candidates = append(candidates, Candidate{text: c})
	}

	return candidates, nil
}

func chooseOne(candidates []Candidate) Selection {
	i := rand.Intn(len(candidates))

	return Selection{
		text: strings.TrimSpace(candidates[i].text),
		idx:  i,
	}
}

func generateAnswerWithChoice(choice Selection) string {
	switch rand.Intn(3) {

	case 0:
		return fmt.Sprintf("\"%s\"로 가시죠. no doubt", choice.text)
	case 1:
		return fmt.Sprintf("Just, \"%s\"", choice.text)
	case 2:
		return fmt.Sprintf("한번만 말한다. \"%s\" 가 답이다.", choice.text)
	}

	return "Nothing to say"
}
