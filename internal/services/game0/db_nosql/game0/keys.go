package game0

import "github.com/gstones/moke-kit/orm/nosql/key"

func NewDemoKey(id string) (key.Key, error) {
	return key.NewKeyFromParts("game", id)
}
