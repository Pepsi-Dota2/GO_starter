package model

import "sync"

type Message struct {
	Data string `json:"data"`
}

type Subscriber struct {
	ID      int
	Channel chan Message
}

type Pubsub struct {
	subscribers  map[int]*Subscriber
	mu           sync.Mutex
	subscriberID int
}
