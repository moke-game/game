package db

import (
	"context"
	"encoding/json"
	"time"

	pb "github.com/moke-game/game/api/gen/gm"
)

func (db *Database) GetBlockedById(uid string) (*pb.BannedInfo, error) {
	res := &pb.BannedInfo{}
	if key, err := makeBlockedListKey(); err != nil {
		return nil, err
	} else if data, err := db.HGet(context.Background(), key.String(), uid).Result(); err != nil {
		return nil, err
	} else if err := json.Unmarshal([]byte(data), res); err != nil {
		return nil, err
	}
	return res, nil
}

func (db *Database) GetBlockedList(uids ...string) (map[string]*pb.BannedInfo, error) {
	if len(uids) == 0 {
		return make(map[string]*pb.BannedInfo), nil
	}
	if key, err := makeBlockedListKey(); err != nil {
		return nil, err
	} else if data, err := db.HMGet(context.Background(), key.String(), uids...).Result(); err != nil {
		return nil, err
	} else {
		res := make(map[string]*pb.BannedInfo)
		for i, v := range data {
			if v != nil {
				var info pb.BannedInfo
				if err := json.Unmarshal([]byte(v.(string)), &info); err != nil {
					return nil, err
				}
				res[uids[i]] = &info
			}
		}
		if err := db.checkAndRemoveUnblock(res); err != nil {
			return nil, err
		}
		return res, nil
	}
}

func (db *Database) checkAndRemoveUnblock(infos map[string]*pb.BannedInfo) error {
	var removeUids []string
	for k, v := range infos {
		if v.UnsealTime > 0 && v.UnsealTime < time.Now().Unix() {
			removeUids = append(removeUids, k)
			delete(infos, k)
		}
	}
	if len(removeUids) <= 0 {
		return nil
	} else if key, err := makeBlockedListKey(); err != nil {
		return err
	} else if err := db.HDel(context.Background(), key.String(), removeUids...).Err(); err != nil {
		return err
	}
	return nil
}

func (db *Database) AddBlockedList(uid string, info *pb.BannedInfo) error {
	if key, err := makeBlockedListKey(); err != nil {
		return err
	} else if data, err := json.Marshal(info); err != nil {
		return err
	} else if err := db.HSet(
		context.Background(),
		key.String(),
		uid,
		data,
	).Err(); err != nil {
		return err
	}
	return nil
}

func (db *Database) RemoveBlockedList(uid string) error {
	if key, err := makeBlockedListKey(); err != nil {
		return err
	} else if err := db.HDel(context.Background(), key.String(), uid).Err(); err != nil {
		return err
	}
	return nil
}
