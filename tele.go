// A package for handling a telegram bot.
//
// Copyright 2019 Maxim R. <epiqmax@gmail.com>
package tele

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/epiqm/seq"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strings"
	"time"
)

var (
	StatsFile = "stats.txt"
	DebugMode bool
)

type Assist struct {
	Actions []Action `json:"actions"`
	Done    Action   `json:"done"`
}
type Option struct {
	Name  string   `json:"name"`
	Text  string   `json:"text"`
	Words []string `json:"words"`
}
type Action struct {
	Name       string   `json:"name"`
	Text       string   `json:"text"`
	FirstText  string   `json:"first_text"`
	Validation bool     `json:"validation"`
	Options    []Option `json:"options"`
}

var Assister Assist

type Progress struct {
	Users []UserProgress `json:"users"`
}
type UserProgress struct {
	Id        int64  `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Step      string `json:"step"`
}

var Usergress Progress

func (a *Progress) Get(par int64) UserProgress {
	for _, u := range a.Users {
		if u.Id == par {
			return u
		}
	}
	return UserProgress{}
}

func (a *Progress) Set(par int64, step string) bool {
	for k, u := range a.Users {
		if u.Id == par {
			ku := &a.Users[k]
			ku.Step = step
			return true
		}
	}
	return false
}

func (a *Assist) Get(par string) Action {
	for _, act := range a.Actions {
		if act.Name == par {
			return act
		}
	}
	return Action{}
}

func (a *Assist) GetString(par string) string {
	for _, c := range a.Actions {
		x := reflect.ValueOf(c).Elem()
		typeOfT := x.Type()
		for i := 0; i < x.NumField(); i++ {
			f := x.Field(i)
			n := strings.ToLower(typeOfT.Field(i).Name)
			if n == par {
				return f.String()
			}
		}
		return ""
	}
	return ""
}

// Building URL for request.
func (bot *Bot) buildCommand(command string) string {
	str := fmt.Sprintf("%s%s/%s", bot.Spawn, bot.Au, command)
	return str
}

type GetMeResult struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	Username  string `json:"username"`
}

type GetMeResponse struct {
	Ok     bool        `json:"ok"`
	Result GetMeResult `json:"result"`
}

type From struct {
	Id           int64  `json:"id"`
	IsBot        bool   `json:"is_bot"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Username     string `json:"username"`
	LanguageCode string `json:"language_code"`
}

type Chat struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	Username  string `json:"username"`
	Type      string `json:"type"`
}

type UpdatesMessage struct {
	MessageId int64  `json:"message_id"`
	From      From   `json:"from"`
	Chat      Chat   `json:"chat"`
	Date      int64  `json:"date"`
	Text      string `json:"text"`
}

type UpdatesResult struct {
	UpdateId int64          `json:"update_id"`
	Message  UpdatesMessage `json:"message"`
	Date     int64          `json:"date"`
	Text     string         `json:"text"`
}

type GetUpdatesResponse struct {
	Ok     bool            `json:"ok"`
	Result []UpdatesResult `json:"result"`
}

type ReplyMessage struct {
	ChatId                int64  `json:"chat_id"`
	Text                  string `json:"text"`
	ParseMode             string `json:"parse_mode"`
	DisableWebPagePreview bool   `json:"disable_web_page_preview"`
	DisableNotification   bool   `json:"disable_notification"`
	ReplyToMessageId      int64  `json:"reply_to_message_id"`
	ReplyMarkup           string `json:"reply_markup"`
}
type NewMessage struct {
	ChatId                string `json:"chat_id"`
	Text                  string `json:"text"`
	ParseMode             string `json:"parse_mode"`
	DisableWebPagePreview bool   `json:"disable_web_page_preview"`
	DisableNotification   bool   `json:"disable_notification"`
	ReplyToMessageId      int64  `json:"reply_to_message_id"`
	ReplyMarkup           string `json:"reply_markup"`
}
type SimpleMessage struct {
	ChatId int64  `json:"chat_id"`
	Text   string `json:"text"`
}
type AdditionalMessage struct {
	ChatId    int64  `json:"chat_id"`
	Text      string `json:"text"`
	ParseMode string `json:"parse_mode"`
}

func (bot *Bot) SendMessage(chat_id int64, text string) (responseBody string, err error) {
	var m SimpleMessage
	m.ChatId = chat_id
	m.Text = text

	var mJson []byte
	mJson, err = json.Marshal(m)
	if err != nil {
		if DebugMode {
			log.Println("send message marshal error:", err.Error())
		}
	}

	newCommand := bot.buildCommand("sendMessage")
	req, err := http.NewRequest("POST", newCommand, bytes.NewBuffer(mJson))
	req.Header.Set("User-agent", "Google Bot+")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		if DebugMode {
			log.Println("failed to send message:", err)
		}
	}
	defer resp.Body.Close()

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	return string(body), nil
}

func (bot *Bot) SendMarkdownMessage(chat_id int64, text string) (responseBody string, err error) {
	var m AdditionalMessage
	m.ChatId = chat_id
	m.Text = text
	m.ParseMode = "Markdown"

	var mJson []byte
	mJson, err = json.Marshal(m)
	if err != nil {
		if DebugMode {
			log.Println("send message marshal error:", err.Error())
		}
	}

	newCommand := bot.buildCommand("sendMessage")
	req, err := http.NewRequest("POST", newCommand, bytes.NewBuffer(mJson))
	req.Header.Set("User-agent", "Google Bot+")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		if DebugMode {
			log.Println("failed to send message:", err)
		}
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	return string(body), nil
}

func (c *UpdatesResult) Dump() (s string) {
	x := reflect.ValueOf(c).Elem()
	typeOfT := x.Type()
	for i := 0; i < x.NumField(); i++ {
		f := x.Field(i)
		s = s + fmt.Sprintf("%s=%v\n",
			strings.ToLower(typeOfT.Field(i).Name), f.Interface())
	}
	return
}

func ReadUpdates(ur *[]UpdatesResult) error {
	txt, err := seq.ReadFile(StatsFile)
	if err != nil {
		return err
	}
	txt, err = seq.Decode(txt, "")
	if err != nil {
		return err
	}
	seq.Unmarshal(txt, &ur)
	return nil
}

func SaveUpdates(ur *[]UpdatesResult) error {
	json := seq.Marshal(ur)
	json, err := seq.Encode(json, "")
	if err != nil {
		return err
	}

	err = seq.WriteFile(StatsFile, json)
	if err != nil {
		return err
	}
	return nil
}

func (c *GetUpdatesResponse) Dump() (s string) {
	x := reflect.ValueOf(c).Elem()
	typeOfT := x.Type()
	for i := 0; i < x.NumField(); i++ {
		f := x.Field(i)
		s = s + fmt.Sprintf("%s=%v\n",
			strings.ToLower(typeOfT.Field(i).Name), f.Interface())
	}
	return
}

func (bot *Bot) GetUpdates(lastUpdateId int64) (response GetUpdatesResponse) {
	newCommand := bot.buildCommand("getUpdates")
	resp, err := http.Get(fmt.Sprintf("%s?offset=%d", newCommand, lastUpdateId))
	if err != nil {
		bot.FailedRequests++
		if DebugMode {
			log.Println("error on spawn request:", newCommand, err.Error())
		}
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		bot.FailedRequests++
		if DebugMode {
			log.Println("failed to read response:", err.Error())
		}
		return
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		bot.FailedRequests++
		if DebugMode {
			log.Println("spawn response unmarshal error: ", err.Error(), string(body))
		}
		return
	}

	bot.SuccessfullRequests++

	bot.LastUpdateId = lastUpdateId
	return response
}

type Bot struct {
	Id                  string
	Au                  string
	Spawn               string
	LastUpdateId        int64
	Created             time.Time
	FailedRequests      int64
	SuccessfullRequests int64
}
type Instances struct {
	Bots []Bot
}

var BotsInstances Instances

func Create(au string, spawn string, lastUpdateId int64) (bot Bot, err error) {
	ausp := fmt.Sprintf("%s%s", au, spawn)
	bot.Id = seq.Hash(ausp)
	bot.Au = au
	bot.Spawn = spawn
	bot.LastUpdateId = lastUpdateId
	bot.Created = time.Now()

	for _, v := range BotsInstances.Bots {
		if v.Id == bot.Id {
			err := errors.New("create: bot instance was already created")
			return bot, err
		}
	}

	BotsInstances.Bots = append(BotsInstances.Bots, bot)

	return bot, nil
}

func GetBots() (bots *[]Bot) {
	return &BotsInstances.Bots
}
