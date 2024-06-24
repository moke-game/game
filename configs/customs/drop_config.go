package customs

import (
	"github.com/duke-git/lancet/v2/random"

	cfg "github.com/moke-game/game/configs/code"
)

type Library struct {
	ID       int32
	DropType int32
	DropId   int32
	DropRate int32
}

type DropLibrary struct {
	Libraries  map[int32][]*cfg.DropLibrary
	DropGroups map[int32]map[int32]*cfg.DropGroup
}

func CreateDropLibrary() *DropLibrary {
	return &DropLibrary{
		Libraries:  make(map[int32][]*cfg.DropLibrary),
		DropGroups: make(map[int32]map[int32]*cfg.DropGroup),
	}
}

func (dl *DropLibrary) Init(tbl *cfg.TblDropLibrary, dgTbl *cfg.TblDropGroup) {
	for _, v := range tbl.GetDataList() {
		if _, ok := dl.Libraries[v.DropID]; !ok {
			dl.Libraries[v.DropID] = make([]*cfg.DropLibrary, 0)
		}
		dl.Libraries[v.DropID] = append(dl.Libraries[v.DropID], v)
	}

	for _, v := range dgTbl.GetDataList() {
		if _, ok := dl.DropGroups[v.DropGroupID]; !ok {
			dl.DropGroups[v.DropGroupID] = make(map[int32]*cfg.DropGroup)
		}
		dl.DropGroups[v.DropGroupID][v.ID] = v
	}
}

type DropConf struct {
	Id          int32
	Count       int32
	AtLeastItem int32
	IsRepeat    bool
}

func (dl *DropLibrary) Random(tp int32) map[int32]int32 {
	items, drops := dl.randomDropGroups(tp)
	dItems := dl.randomDrops(drops)
	for k, v := range items {
		dItems[k] += v
	}
	return dItems
}

func randomWithWeight(weights map[int32]int32) int32 {
	weightSum := 0
	for _, v := range weights {
		weightSum += int(v)
	}

	weight := random.RandInt(0, weightSum)
	for k, v := range weights {
		if weight > int(v) {
			weight -= int(v)
			continue
		}
		return k
	}
	return -1
}

func (dl *DropLibrary) randomItems(drops map[int32]*cfg.DropGroup, times int32, isRepeat bool) map[int32]int32 {
	res := make(map[int32]int32)
	if times >= int32(len(drops)) && !isRepeat {
		res = dl.addAllItems(drops)
		times -= int32(len(drops))
	}

	weights := map[int32]int32{}
	for k, v := range drops {
		weights[k] = v.DropWeights
	}

	for i := int32(0); i < times; i++ {
		dropId := randomWithWeight(weights)
		itemId := drops[dropId].ItemID
		itemNum := random.RandInt(int(drops[dropId].MinNum), int(drops[dropId].MaxNum))
		res[itemId] = int32(itemNum)
		if !isRepeat {
			delete(weights, itemId)
		}
	}
	return res
}

func (dl *DropLibrary) addAllItems(drops map[int32]*cfg.DropGroup) map[int32]int32 {
	res := make(map[int32]int32)
	for _, v := range drops {
		num := random.RandInt(int(v.MinNum), int(v.MaxNum))
		res[v.ItemID] += int32(num)
	}
	return res
}

func (dl *DropLibrary) randomDrops(dropsGrops map[int32]*DropConf) map[int32]int32 {
	res := make(map[int32]int32)
	for k, v := range dropsGrops {
		if dg, ok := dl.DropGroups[k]; ok {
			items := dl.randomItems(dg, v.Count, v.IsRepeat)
			for ik, iv := range items {
				res[ik] += iv
			}
		}
	}
	return res
}

// randomDropGroups 返回  道具组id, 掉落组id
func (dl *DropLibrary) randomDropGroups(tp int32) (map[int32]int32, map[int32]*DropConf) {
	libs, ok := dl.Libraries[tp]
	if !ok {
		return nil, nil
	}
	items := make(map[int32]int32)
	drops := make(map[int32]*DropConf)

	for _, v := range libs {
		dgId := v.DropValue
		rate := random.RandInt(0, 10000)
		if rate > int(v.DropRate) {
			if v.BaseLineID == 0 {
				continue
			}
			dgId = v.BaseLineID
		}
		count := random.RandInt(int(v.MinDropTimes), int(v.MaxDropTimes))
		if v.DropType == 2 {
			items[dgId] += int32(count)
		} else {
			drops[dgId] = &DropConf{
				Id:          dgId,
				Count:       int32(count),
				AtLeastItem: v.BaseLineID,
				IsRepeat:    v.IsDitto == 1,
			}
		}
	}
	return items, drops
}
