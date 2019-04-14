package service

import (
	"github.com/yakaa/cuter/lib/logx"
	"github.com/yakaa/cuter/lib/messages"
)

type loggedConsumer struct {
	ContentConsumer
}

func NewLoggedConsumer(consumer ContentConsumer) ContentConsumer {
	return &loggedConsumer{
		ContentConsumer: consumer,
	}
}

func (consumer *loggedConsumer) Consume(message *messages.JsonMessage) error {
	logx.Info("=>", message.Raw)
	return consumer.ContentConsumer.Consume(message)
}
