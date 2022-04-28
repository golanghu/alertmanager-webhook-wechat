package notify

import "context"

type Notifier interface {
	Notify(context.Context, interface{}) (bool, error)
}
