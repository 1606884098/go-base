package main

// SMS短信
// Email

type AbstractMessage interface {
	SendMessage(text, to string)
}

type MessageImlementer interface {
	Send(text, to string)
}
