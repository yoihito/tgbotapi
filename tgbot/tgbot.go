package tgbot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Bot struct {
	Token string
	ApiUrl string
	Handlers map[string]HandlerFunc
	Poller Poller
	StopClient chan struct{}
	Stop chan chan struct{}
}

type UpdatesResponse struct {
	Result []Update
}

type MessageResponse struct {
	Result *Message
}

func NewBot(pref Settings) (*Bot, error) {
	bot := &Bot{
		Token: pref.Token,
		ApiUrl: pref.ApiUrl,
		Poller: pref.Poller,
		Handlers: make(map[string]HandlerFunc),
	}
	return bot, nil
}

func (bot *Bot) Start() {
	stop := make(chan struct{})
	updates := make(chan Update)
	go func() {
		bot.Poller.Listen(bot, updates, stop)
	}()

	for {
		select {
		case upd := <-updates:
			bot.ProcessUpdate(upd)
		}
	}
}

func (bot *Bot) SendMessage(to Recipient, message string) (*Message, error) {
	values := map[string]string{"chat_id": to.Recipient(), "text": message}
	jsonData, err := bot.MakeRequest("sendMessage", values)
	if err != nil {
		return nil, err
	}
	var resp MessageResponse
	if err := json.Unmarshal(jsonData, &resp); err != nil {
		return nil, err
	}
	return resp.Result, nil
}

func (bot *Bot) GetUpdates(offset UpdateId) ([]Update, error) {
	values :=  map[string]string{ "offset": strconv.Itoa(int(offset))}
	jsonData, err := bot.MakeRequest("getUpdates", values)
	if err != nil {
		return nil, err
	}
	var resp UpdatesResponse
	if err := json.Unmarshal(jsonData, &resp); err != nil {
		return nil, err
	}
	return resp.Result, nil
}

func (bot *Bot) MakeRequest(command string, values map[string]string) ([]byte, error) {
	request := fmt.Sprintf("%s%s/%s", bot.ApiUrl, bot.Token, command)
	jsonData, _ := json.Marshal(values)
	resp, err := http.Post(request, "application/json", bytes.NewReader(jsonData))
	if err != nil {
		log.Fatal(err)
		return nil, err
  }

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
			return nil, err
		}
		log.Fatalln(string(b))
		return nil, fmt.Errorf("HTTP Error: %d", resp.StatusCode)
	}
	rawData, err := ioutil.ReadAll(resp.Body)
	log.Println(string(rawData))
	if err != nil {
		return nil, err
	}
	return rawData, nil
}

func (bot *Bot) OnText(text string, handlerFn HandlerFunc) {
	bot.Handlers[text] = handlerFn
}

func (bot *Bot) ProcessUpdate(update Update) {
	go func () {
		if update.Message == nil {
			log.Println("Unknown command")
			return
		}

		handler, ok := bot.Handlers[update.Message.Text]
		if !ok {
			log.Printf("Handler not found: %s\n", update.Message.Text)
			return
		}
		handler(HandlerContext{bot, update})
	}()
}
