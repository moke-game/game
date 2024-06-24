package notification

import (
	"fmt"

	"github.com/gstones/moke-kit/mq/common"
)

// 游戏内事件通知

type NotifyEvent int32

const (
	NotifyEventNone          NotifyEvent = 0
	NotifyEventLoginOverride NotifyEvent = 1 // 重复登录, 踢掉前一个登录
	NotifyEventKickOffline   NotifyEvent = 2 // GM踢掉下线
)

func MakeGamePrivateNotifyTopic(uid string) string {
	return common.NatsHeader.CreateTopic(fmt.Sprintf("rumble.notify.%s", uid))
}
func MakeGamePublicNotifyTopic() string {
	return common.NatsHeader.CreateTopic("rumble.notify.public")
}
