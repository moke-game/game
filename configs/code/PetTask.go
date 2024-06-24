//------------------------------------------------------------------------------
// <auto-generated>
//     This code was generated by a tool.
//     Changes to this file may cause incorrect behavior and will be lost if
//     the code is regenerated.
// </auto-generated>
//------------------------------------------------------------------------------

package cfg

import "errors"

type PetTask struct {
	ID            int32
	FinishCost    int32
	Refresh       []*ItemReward
	TaskCount     int32
	TaskExtension []*ItemReward
}

const TypeId_PetTask = 987122468

func (*PetTask) GetTypeId() int32 {
	return 987122468
}

func (_v *PetTask) Deserialize(_buf map[string]interface{}) (err error) {
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
		if _tempNum_, _ok_ = _buf["FinishCost"].(float64); !_ok_ {
			err = errors.New("FinishCost error")
			return
		}
		_v.FinishCost = int32(_tempNum_)
	}
	{
		var _arr_ []interface{}
		var _ok_ bool
		if _arr_, _ok_ = _buf["Refresh"].([]interface{}); !_ok_ {
			err = errors.New("Refresh error")
			return
		}

		_v.Refresh = make([]*ItemReward, 0, len(_arr_))

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
			_v.Refresh = append(_v.Refresh, _list_v_)
		}
	}

	{
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["TaskCount"].(float64); !_ok_ {
			err = errors.New("TaskCount error")
			return
		}
		_v.TaskCount = int32(_tempNum_)
	}
	{
		var _arr_ []interface{}
		var _ok_ bool
		if _arr_, _ok_ = _buf["TaskExtension"].([]interface{}); !_ok_ {
			err = errors.New("TaskExtension error")
			return
		}

		_v.TaskExtension = make([]*ItemReward, 0, len(_arr_))

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
			_v.TaskExtension = append(_v.TaskExtension, _list_v_)
		}
	}

	return
}

func DeserializePetTask(_buf map[string]interface{}) (*PetTask, error) {
	v := &PetTask{}
	if err := v.Deserialize(_buf); err == nil {
		return v, nil
	} else {
		return nil, err
	}
}
