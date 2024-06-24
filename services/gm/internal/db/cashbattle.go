package db

import (
	"context"
	"encoding/json"

	pb "github.com/moke-game/game/api/gen/gm"
)

func (db *Database) GetCashBattle(id string) (*pb.WatchGMResponse, error) {
	res := &pb.WatchGMResponse{}
	if key, err := makeCashBattleKey(id); err != nil {
		return nil, err
	} else if err := db.HGetAll(context.Background(), key.String()).Scan(res); err != nil {
		return nil, err
	}
	return res, nil
}

func (db *Database) AddCashBattleInfo(id string, info *pb.CashBattleInfo) error {
	if key, err := makeCashBattleKey(id); err != nil {
		return err
	} else if data, err := json.Marshal(info); err != nil {
		return err
	} else if err := db.HSet(
		context.Background(),
		key.String(),
		"info",
		data,
	).Err(); err != nil {
		return err
	}
	return nil
}

func (db *Database) AddCashBattleMatch(id string, info *pb.CashBattleMatch) error {
	if key, err := makeCashBattleKey(id); err != nil {
		return err
	} else if data, err := json.Marshal(info); err != nil {
		return err
	} else if err := db.HSet(
		context.Background(),
		key.String(),
		"match",
		data,
	).Err(); err != nil {
		return err
	}
	return nil
}

func (db *Database) AddCashBattleCoe(id string, info *pb.CashBattleCoe) error {
	if key, err := makeCashBattleKey(id); err != nil {
		return err
	} else if data, err := json.Marshal(info); err != nil {
		return err
	} else if err := db.HSet(
		context.Background(),
		key.String(),
		"coe",
		data,
	).Err(); err != nil {
		return err
	}
	return nil
}

func (db *Database) GetCashBattleSignInfo(id string) ([]*pb.CashPlayerInfo, error) {
	res := make([]*pb.CashPlayerInfo, 0)
	if key, err := makeCashBattleSignKey(id); err != nil {
		return nil, err
	} else if data, err := db.HGetAll(context.Background(), key.String()).Result(); err != nil {
		return nil, err
	} else {
		for _, v := range data {
			var player pb.CashPlayerInfo
			if err := json.Unmarshal([]byte(v), &player); err != nil {
				return nil, err
			}
			res = append(res, &player)
		}
	}
	return res, nil
}

func (db *Database) SignCashBattle(id string, player *pb.CashPlayerInfo) error {
	if key, err := makeCashBattleSignKey(id); err != nil {
		return err
	} else if data, err := json.Marshal(player); err != nil {
		return err
	} else if err := db.HSet(
		context.Background(),
		key.String(),
		player.Uid,
		data,
	).Err(); err != nil {
		return err
	}
	return nil
}

func (db *Database) ReportCashBattleResult(id string, result *pb.ResultInfo) error {
	if key, err := makeCashBattleResultKey(id); err != nil {
		return err
	} else if data, err := json.Marshal(result); err != nil {
		return err
	} else if err := db.HSet(
		context.Background(),
		key.String(),
		result.RoomId,
		data,
	).Err(); err != nil {
		return err
	}
	return nil
}

func (db *Database) GetCashBattleResult(id string) ([]*pb.ResultInfo, error) {
	res := make([]*pb.ResultInfo, 0)
	if key, err := makeCashBattleResultKey(id); err != nil {
		return nil, err
	} else if data, err := db.HGetAll(context.Background(), key.String()).Result(); err != nil {
		return nil, err
	} else {
		for _, v := range data {
			var result pb.ResultInfo
			if err := json.Unmarshal([]byte(v), &result); err != nil {
				return nil, err
			}
			res = append(res, &result)
		}
	}
	return res, nil
}
