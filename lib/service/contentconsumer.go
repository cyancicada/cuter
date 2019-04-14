package service

import "github.com/yakaa/cuter/lib/messages"

type ContentConsumer interface {
	Consume(*messages.JsonMessage) error
}
