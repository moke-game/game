//------------------------------------------------------------------------------
// <auto-generated>
//     This code was generated by a tool.
//     Changes to this file may cause incorrect behavior and will be lost if
//     the code is regenerated.
// </auto-generated>
//------------------------------------------------------------------------------

package cfg

import "errors"

type ChatEmoji struct {
	ID         int32
	Name       string
	ResName    []string
	Tag        int32
	Interval   []float32
	UnlockType []int32
}

const TypeId_ChatEmoji = -695110738

func (*ChatEmoji) GetTypeId() int32 {
	return -695110738
}

func (_v *ChatEmoji) Deserialize(_buf map[string]interface{}) (err error) {
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
		if _v.Name, _ok_ = _buf["Name"].(string); !_ok_ {
			err = errors.New("Name error")
			return
		}
	}
	{
		var _arr_ []interface{}
		var _ok_ bool
		if _arr_, _ok_ = _buf["ResName"].([]interface{}); !_ok_ {
			err = errors.New("ResName error")
			return
		}

		_v.ResName = make([]string, 0, len(_arr_))

		for _, _e_ := range _arr_ {
			var _list_v_ string
			{
				if _list_v_, _ok_ = _e_.(string); !_ok_ {
					err = errors.New("_list_v_ error")
					return
				}
			}
			_v.ResName = append(_v.ResName, _list_v_)
		}
	}

	{
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["Tag"].(float64); !_ok_ {
			err = errors.New("Tag error")
			return
		}
		_v.Tag = int32(_tempNum_)
	}
	{
		var _arr_ []interface{}
		var _ok_ bool
		if _arr_, _ok_ = _buf["Interval"].([]interface{}); !_ok_ {
			err = errors.New("Interval error")
			return
		}

		_v.Interval = make([]float32, 0, len(_arr_))

		for _, _e_ := range _arr_ {
			var _list_v_ float32
			{
				var _ok_ bool
				var _x_ float64
				if _x_, _ok_ = _e_.(float64); !_ok_ {
					err = errors.New("_list_v_ error")
					return
				}
				_list_v_ = float32(_x_)
			}
			_v.Interval = append(_v.Interval, _list_v_)
		}
	}

	{
		var _arr_ []interface{}
		var _ok_ bool
		if _arr_, _ok_ = _buf["UnlockType"].([]interface{}); !_ok_ {
			err = errors.New("UnlockType error")
			return
		}

		_v.UnlockType = make([]int32, 0, len(_arr_))

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
			_v.UnlockType = append(_v.UnlockType, _list_v_)
		}
	}

	return
}

func DeserializeChatEmoji(_buf map[string]interface{}) (*ChatEmoji, error) {
	v := &ChatEmoji{}
	if err := v.Deserialize(_buf); err == nil {
		return v, nil
	} else {
		return nil, err
	}
}
