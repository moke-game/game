package service

import (
	"fmt"

	"github.com/gstones/moke-kit/mq/common"
)

func makeBlockedListTopic(uid string) string {
	topic := fmt.Sprintf("blockedlist.%s", uid)
	return common.NatsHeader.CreateTopic(topic)
}
