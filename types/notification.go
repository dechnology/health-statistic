package types

import "github.com/eesoymilk/health-statistic-api/ent"

type NotificationWithRecipient struct {
	ent.Notification
	Recipient ent.User `json:"recipient"`
}
