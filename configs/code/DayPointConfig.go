//------------------------------------------------------------------------------
// <auto-generated>
//     This code was generated by a tool.
//     Changes to this file may cause incorrect behavior and will be lost if
//     the code is regenerated.
// </auto-generated>
//------------------------------------------------------------------------------

package cfg

import "errors"

type DayPointConfig struct {
	ID          int32
	PointTarget int32
	Reward      []*ItemReward
}

const TypeId_DayPointConfig = 1486893014

func (*DayPointConfig) GetTypeId() int32 {
	return 1486893014
}

func (_v *DayPointConfig) Deserialize(_buf map[string]interface{}) (err error) {
	{
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["ID"].(float64); !_ok_ {
			err = errors.New("ID error")
			return
		}
		_v.ID = int32(_tempNum_)
	}
	{
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["PointTarget"].(float64); !_ok_ {
			err = errors.New("PointTarget error")
			return
		}
		_v.PointTarget = int32(_tempNum_)
	}
	{
		var _arr_ []interface{}
		var _ok_ bool
		if _arr_, _ok_ = _buf["Reward"].([]interface{}); !_ok_ {
			err = errors.New("Reward error")
			return
		}

		_v.Reward = make([]*ItemReward, 0, len(_arr_))

		for _, _e_ := range _arr_ {
			var _list_v_ *ItemReward
			{
				var _ok_ bool
				var _x_ map[string]interface{}
				if _x_, _ok_ = _e_.(map[string]interface{}); !_ok_ {
					err = errors.New("_list_v_ error")
					return
				}
				if _list_v_, err = DeserializeItemReward(_x_); err != nil {
					return
				}
			}
			_v.Reward = append(_v.Reward, _list_v_)
		}
	}

	return
}

func DeserializeDayPointConfig(_buf map[string]interface{}) (*DayPointConfig, error) {
	v := &DayPointConfig{}
	if err := v.Deserialize(_buf); err == nil {
		return v, nil
	} else {
		return nil, err
	}
}
