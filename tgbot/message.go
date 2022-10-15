package tgbot

import "strconv"

type Message struct {
	Id int `json:"message_id"`
	Sender *User `json:"from"`
	SenderChat *Chat `json:"sender_chat"`
	Date int `json:"date"`
	Chat *Chat `json:"chat"`
	ForwardFrom *User `json:"forward_from"`
	ForwardFromChat *Chat `json:"forward_from_chat"`
	ForwardFromMessageId int `json:"forward_from_message_id"`
	ForwardSignature string `json:"forward_signature"`
	ForwardSenderName string `json:"forward_sender_name"`
	Text string `json:"text"`
}

type User struct {
	Id int `json:"id"`
	IsBot bool `json:"is_bot"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Username string `json:"username"`
}

func (u *User) Recipient() string {
	return strconv.Itoa(u.Id)
}

type Chat struct {
	Id int `json:"id"`
	Type string `json:"type"`
}

