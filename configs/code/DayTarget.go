//------------------------------------------------------------------------------
// <auto-generated>
//     This code was generated by a tool.
//     Changes to this file may cause incorrect behavior and will be lost if
//     the code is regenerated.
// </auto-generated>
//------------------------------------------------------------------------------

package cfg

import "errors"

type DayTarget struct {
	ID                  int32
	TaskID              []int32
	CoinGiftID          int32
	ItemID1             int32
	CoinGiftOriginal    int32
	CoinGiftCurrent     int32
	DiamondGiftID       int32
	ItemID2             int32
	DiamondGiftOriginal int32
	DiamondGiftCurrent  int32
}

const TypeId_DayTarget = 1412750381

func (*DayTarget) GetTypeId() int32 {
	return 1412750381
}

func (_v *DayTarget) Deserialize(_buf map[string]interface{}) (err error) {
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
		var _arr_ []interface{}
		var _ok_ bool
		if _arr_, _ok_ = _buf["TaskID"].([]interface{}); !_ok_ {
			err = errors.New("TaskID error")
			return
		}

		_v.TaskID = make([]int32, 0, len(_arr_))

		for _, _e_ := range _arr_ {
			var _list_v_ int32
			{
				var _ok_ bool
				var _x_ float64
				if _x_, _ok_ = _e_.(float64); !_ok_ {
					err = errors.New("_list_v_ error")
					return
				}
				_list_v_ = int32(_x_)
			}
			_v.TaskID = append(_v.TaskID, _list_v_)
		}
	}

	{
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["CoinGiftID"].(float64); !_ok_ {
			err = errors.New("CoinGiftID error")
			return
		}
		_v.CoinGiftID = int32(_tempNum_)
	}
	{
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["ItemID1"].(float64); !_ok_ {
			err = errors.New("ItemID1 error")
			return
		}
		_v.ItemID1 = int32(_tempNum_)
	}
	{
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["CoinGiftOriginal"].(float64); !_ok_ {
			err = errors.New("CoinGiftOriginal error")
			return
		}
		_v.CoinGiftOriginal = int32(_tempNum_)
	}
	{
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["CoinGiftCurrent"].(float64); !_ok_ {
			err = errors.New("CoinGiftCurrent error")
			return
		}
		_v.CoinGiftCurrent = int32(_tempNum_)
	}
	{
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["DiamondGiftID"].(float64); !_ok_ {
			err = errors.New("DiamondGiftID error")
			return
		}
		_v.DiamondGiftID = int32(_tempNum_)
	}
	{
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["ItemID2"].(float64); !_ok_ {
			err = errors.New("ItemID2 error")
			return
		}
		_v.ItemID2 = int32(_tempNum_)
	}
	{
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["DiamondGiftOriginal"].(float64); !_ok_ {
			err = errors.New("DiamondGiftOriginal error")
			return
		}
		_v.DiamondGiftOriginal = int32(_tempNum_)
	}
	{
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["DiamondGiftCurrent"].(float64); !_ok_ {
			err = errors.New("DiamondGiftCurrent error")
			return
		}
		_v.DiamondGiftCurrent = int32(_tempNum_)
	}
	return
}

func DeserializeDayTarget(_buf map[string]interface{}) (*DayTarget, error) {
	v := &DayTarget{}
	if err := v.Deserialize(_buf); err == nil {
		return v, nil
	} else {
		return nil, err
	}
}
