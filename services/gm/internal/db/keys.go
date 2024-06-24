package db

import "github.com/gstones/moke-kit/orm/nosql/key"

func makeBlockedListKey() (key.Key, error) {
	return key.NewKeyFromParts("gm", "blockedlist")
}

func makeCashBattleKey(id string) (key.Key, error) {
	return key.NewKeyFromParts("gm", "cashbattle", "info", id)
}

func makeCashBattleSignKey(id string) (key.Key, error) {
	return key.NewKeyFromParts("gm", "cashbattle", "sign", id)
}

func makeCashBattleResultKey(id string) (key.Key, error) {
	return key.NewKeyFromParts("gm", "cashbattle", "result", id)
}
