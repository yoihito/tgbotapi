package tgbot

type UpdateId int

type Update struct {
	UpdateId UpdateId `json:"update_id"`
	Message *Message
}
