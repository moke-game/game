//------------------------------------------------------------------------------
// <auto-generated>
//     This code was generated by a tool.
//     Changes to this file may cause incorrect behavior and will be lost if
//     the code is regenerated.
// </auto-generated>
//------------------------------------------------------------------------------

package cfg

import "errors"

type BattleSpecialReward struct {
	MapId  int32
	Rank   int32
	DropId int32
}

const TypeId_BattleSpecialReward = 1670932082

func (*BattleSpecialReward) GetTypeId() int32 {
	return 1670932082
}

func (_v *BattleSpecialReward) Deserialize(_buf map[string]interface{}) (err error) {
	{
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["mapId"].(float64); !_ok_ {
			err = errors.New("mapId error")
			return
		}
		_v.MapId = int32(_tempNum_)
	}
	{
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["rank"].(float64); !_ok_ {
			err = errors.New("rank error")
			return
		}
		_v.Rank = int32(_tempNum_)
	}
	{
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["dropId"].(float64); !_ok_ {
			err = errors.New("dropId error")
			return
		}
		_v.DropId = int32(_tempNum_)
	}
	return
}

func DeserializeBattleSpecialReward(_buf map[string]interface{}) (*BattleSpecialReward, error) {
	v := &BattleSpecialReward{}
	if err := v.Deserialize(_buf); err == nil {
		return v, nil
	} else {
		return nil, err
	}
}
