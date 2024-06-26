//------------------------------------------------------------------------------
// <auto-generated>
//     This code was generated by a tool.
//     Changes to this file may cause incorrect behavior and will be lost if
//     the code is regenerated.
// </auto-generated>
//------------------------------------------------------------------------------

package cfg

import "errors"

type Item struct {
	ID              int32
	Desc            string
	Name            string
	Deps            string
	Icon            string
	ShopIcon        string
	Quality         int32
	ItemType        int32
	DiamondExchange int32
	Default         int32
	Itemparameters  int32
	Itemparameters2 int32
}

const TypeId_Item = 2289459

func (*Item) GetTypeId() int32 {
	return 2289459
}

func (_v *Item) Deserialize(_buf map[string]interface{}) (err error) {
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
		if _v.Desc, _ok_ = _buf["Desc"].(string); !_ok_ {
			err = errors.New("Desc error")
			return
		}
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
		if _v.Deps, _ok_ = _buf["Deps"].(string); !_ok_ {
			err = errors.New("Deps error")
			return
		}
	}
	{
		var _ok_ bool
		if _v.Icon, _ok_ = _buf["Icon"].(string); !_ok_ {
			err = errors.New("Icon error")
			return
		}
	}
	{
		var _ok_ bool
		if _v.ShopIcon, _ok_ = _buf["ShopIcon"].(string); !_ok_ {
			err = errors.New("ShopIcon error")
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
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["ItemType"].(float64); !_ok_ {
			err = errors.New("ItemType error")
			return
		}
		_v.ItemType = int32(_tempNum_)
	}
	{
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["DiamondExchange"].(float64); !_ok_ {
			err = errors.New("DiamondExchange error")
			return
		}
		_v.DiamondExchange = int32(_tempNum_)
	}
	{
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["Default"].(float64); !_ok_ {
			err = errors.New("Default error")
			return
		}
		_v.Default = int32(_tempNum_)
	}
	{
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["Itemparameters"].(float64); !_ok_ {
			err = errors.New("Itemparameters error")
			return
		}
		_v.Itemparameters = int32(_tempNum_)
	}
	{
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["Itemparameters2"].(float64); !_ok_ {
			err = errors.New("Itemparameters2 error")
			return
		}
		_v.Itemparameters2 = int32(_tempNum_)
	}
	return
}

func DeserializeItem(_buf map[string]interface{}) (*Item, error) {
	v := &Item{}
	if err := v.Deserialize(_buf); err == nil {
		return v, nil
	} else {
		return nil, err
	}
}
