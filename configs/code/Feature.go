//------------------------------------------------------------------------------
// <auto-generated>
//     This code was generated by a tool.
//     Changes to this file may cause incorrect behavior and will be lost if
//     the code is regenerated.
// </auto-generated>
//------------------------------------------------------------------------------

package cfg

import "errors"

type Feature struct {
	ID            int32
	Name          string
	Text          string
	Type          int32
	MoneySettings []int32
	DefaultSystem int32
	Desc          string
}

const TypeId_Feature = 685445846

func (*Feature) GetTypeId() int32 {
	return 685445846
}

func (_v *Feature) Deserialize(_buf map[string]interface{}) (err error) {
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
		var _ok_ bool
		if _v.Text, _ok_ = _buf["Text"].(string); !_ok_ {
			err = errors.New("Text error")
			return
		}
	}
	{
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["Type"].(float64); !_ok_ {
			err = errors.New("Type error")
			return
		}
		_v.Type = int32(_tempNum_)
	}
	{
		var _arr_ []interface{}
		var _ok_ bool
		if _arr_, _ok_ = _buf["MoneySettings"].([]interface{}); !_ok_ {
			err = errors.New("MoneySettings error")
			return
		}

		_v.MoneySettings = make([]int32, 0, len(_arr_))

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
			_v.MoneySettings = append(_v.MoneySettings, _list_v_)
		}
	}

	{
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["DefaultSystem"].(float64); !_ok_ {
			err = errors.New("DefaultSystem error")
			return
		}
		_v.DefaultSystem = int32(_tempNum_)
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

func DeserializeFeature(_buf map[string]interface{}) (*Feature, error) {
	v := &Feature{}
	if err := v.Deserialize(_buf); err == nil {
		return v, nil
	} else {
		return nil, err
	}
}
