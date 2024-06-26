//------------------------------------------------------------------------------
// <auto-generated>
//     This code was generated by a tool.
//     Changes to this file may cause incorrect behavior and will be lost if
//     the code is regenerated.
// </auto-generated>
//------------------------------------------------------------------------------

package cfg

import "errors"

type AiMatchFile struct {
	ID         int32
	PlayTypeID int32
	MapID      int32
	CupsNum    []int32
	ScriptID   []int32
	DelayCount int32
	Desc       string
}

const TypeId_AiMatchFile = -1674198663

func (*AiMatchFile) GetTypeId() int32 {
	return -1674198663
}

func (_v *AiMatchFile) Deserialize(_buf map[string]interface{}) (err error) {
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
		if _tempNum_, _ok_ = _buf["PlayTypeID"].(float64); !_ok_ {
			err = errors.New("PlayTypeID error")
			return
		}
		_v.PlayTypeID = int32(_tempNum_)
	}
	{
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["MapID"].(float64); !_ok_ {
			err = errors.New("MapID error")
			return
		}
		_v.MapID = int32(_tempNum_)
	}
	{
		var _arr_ []interface{}
		var _ok_ bool
		if _arr_, _ok_ = _buf["CupsNum"].([]interface{}); !_ok_ {
			err = errors.New("CupsNum error")
			return
		}

		_v.CupsNum = make([]int32, 0, len(_arr_))

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
			_v.CupsNum = append(_v.CupsNum, _list_v_)
		}
	}

	{
		var _arr_ []interface{}
		var _ok_ bool
		if _arr_, _ok_ = _buf["ScriptID"].([]interface{}); !_ok_ {
			err = errors.New("ScriptID error")
			return
		}

		_v.ScriptID = make([]int32, 0, len(_arr_))

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
			_v.ScriptID = append(_v.ScriptID, _list_v_)
		}
	}

	{
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["DelayCount"].(float64); !_ok_ {
			err = errors.New("DelayCount error")
			return
		}
		_v.DelayCount = int32(_tempNum_)
	}
	{
		var _ok_ bool
		if _v.Desc, _ok_ = _buf["Desc"].(string); !_ok_ {
			err = errors.New("Desc error")
			return
		}
	}
	return
}

func DeserializeAiMatchFile(_buf map[string]interface{}) (*AiMatchFile, error) {
	v := &AiMatchFile{}
	if err := v.Deserialize(_buf); err == nil {
		return v, nil
	} else {
		return nil, err
	}
}
