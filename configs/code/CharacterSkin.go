//------------------------------------------------------------------------------
// <auto-generated>
//     This code was generated by a tool.
//     Changes to this file may cause incorrect behavior and will be lost if
//     the code is regenerated.
// </auto-generated>
//------------------------------------------------------------------------------

package cfg

import "errors"

type CharacterSkin struct {
	ID         int32
	Name       string
	Desp       string
	Res        string
	Quality    int32
	Price      []*ItemReward
	SourceNote int32
	View       int32
	Goto       int32
	GotoPara   []string
	UnlockDesc string
}

const TypeId_CharacterSkin = -991000538

func (*CharacterSkin) GetTypeId() int32 {
	return -991000538
}

func (_v *CharacterSkin) Deserialize(_buf map[string]interface{}) (err error) {
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
		if _v.Desp, _ok_ = _buf["Desp"].(string); !_ok_ {
			err = errors.New("Desp error")
			return
		}
	}
	{
		var _ok_ bool
		if _v.Res, _ok_ = _buf["Res"].(string); !_ok_ {
			err = errors.New("Res error")
			return
		}
	}
	{
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["Quality"].(float64); !_ok_ {
			err = errors.New("Quality error")
			return
		}
		_v.Quality = int32(_tempNum_)
	}
	{
		var _arr_ []interface{}
		var _ok_ bool
		if _arr_, _ok_ = _buf["Price"].([]interface{}); !_ok_ {
			err = errors.New("Price error")
			return
		}

		_v.Price = make([]*ItemReward, 0, len(_arr_))

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
			_v.Price = append(_v.Price, _list_v_)
		}
	}

	{
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["SourceNote"].(float64); !_ok_ {
			err = errors.New("SourceNote error")
			return
		}
		_v.SourceNote = int32(_tempNum_)
	}
	{
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["View"].(float64); !_ok_ {
			err = errors.New("View error")
			return
		}
		_v.View = int32(_tempNum_)
	}
	{
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["Goto"].(float64); !_ok_ {
			err = errors.New("Goto error")
			return
		}
		_v.Goto = int32(_tempNum_)
	}
	{
		var _arr_ []interface{}
		var _ok_ bool
		if _arr_, _ok_ = _buf["GotoPara"].([]interface{}); !_ok_ {
			err = errors.New("GotoPara error")
			return
		}

		_v.GotoPara = make([]string, 0, len(_arr_))

		for _, _e_ := range _arr_ {
			var _list_v_ string
			{
				if _list_v_, _ok_ = _e_.(string); !_ok_ {
					err = errors.New("_list_v_ error")
					return
				}
			}
			_v.GotoPara = append(_v.GotoPara, _list_v_)
		}
	}

	{
		var _ok_ bool
		if _v.UnlockDesc, _ok_ = _buf["UnlockDesc"].(string); !_ok_ {
			err = errors.New("UnlockDesc error")
			return
		}
	}
	return
}

func DeserializeCharacterSkin(_buf map[string]interface{}) (*CharacterSkin, error) {
	v := &CharacterSkin{}
	if err := v.Deserialize(_buf); err == nil {
		return v, nil
	} else {
		return nil, err
	}
}
