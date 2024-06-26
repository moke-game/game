//------------------------------------------------------------------------------
// <auto-generated>
//     This code was generated by a tool.
//     Changes to this file may cause incorrect behavior and will be lost if
//     the code is regenerated.
// </auto-generated>
//------------------------------------------------------------------------------

package cfg

import "errors"

type Shop struct {
	ID         int32
	Name       string
	Desp       string
	ActType    int32
	ItemID     int32
	WindowType int32
	Price      int32
	PuchaseID  int32
	AwardType  int32
	AwardIDs   []int32
	Closed     int32
	SaleOff    int32
	BuyLimit   []int32
	Res        []int32
	BeginTime  string
	EndTime    string
	ShopType   int32
	Display    int32
}

const TypeId_Shop = 2576150

func (*Shop) GetTypeId() int32 {
	return 2576150
}

func (_v *Shop) Deserialize(_buf map[string]interface{}) (err error) {
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
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["ActType"].(float64); !_ok_ {
			err = errors.New("ActType error")
			return
		}
		_v.ActType = int32(_tempNum_)
	}
	{
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["ItemID"].(float64); !_ok_ {
			err = errors.New("ItemID error")
			return
		}
		_v.ItemID = int32(_tempNum_)
	}
	{
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["WindowType"].(float64); !_ok_ {
			err = errors.New("WindowType error")
			return
		}
		_v.WindowType = int32(_tempNum_)
	}
	{
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["Price"].(float64); !_ok_ {
			err = errors.New("Price error")
			return
		}
		_v.Price = int32(_tempNum_)
	}
	{
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["PuchaseID"].(float64); !_ok_ {
			err = errors.New("PuchaseID error")
			return
		}
		_v.PuchaseID = int32(_tempNum_)
	}
	{
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["AwardType"].(float64); !_ok_ {
			err = errors.New("AwardType error")
			return
		}
		_v.AwardType = int32(_tempNum_)
	}
	{
		var _arr_ []interface{}
		var _ok_ bool
		if _arr_, _ok_ = _buf["AwardIDs"].([]interface{}); !_ok_ {
			err = errors.New("AwardIDs error")
			return
		}

		_v.AwardIDs = make([]int32, 0, len(_arr_))

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
			_v.AwardIDs = append(_v.AwardIDs, _list_v_)
		}
	}

	{
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["Closed"].(float64); !_ok_ {
			err = errors.New("Closed error")
			return
		}
		_v.Closed = int32(_tempNum_)
	}
	{
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["SaleOff"].(float64); !_ok_ {
			err = errors.New("SaleOff error")
			return
		}
		_v.SaleOff = int32(_tempNum_)
	}
	{
		var _arr_ []interface{}
		var _ok_ bool
		if _arr_, _ok_ = _buf["BuyLimit"].([]interface{}); !_ok_ {
			err = errors.New("BuyLimit error")
			return
		}

		_v.BuyLimit = make([]int32, 0, len(_arr_))

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
			_v.BuyLimit = append(_v.BuyLimit, _list_v_)
		}
	}

	{
		var _arr_ []interface{}
		var _ok_ bool
		if _arr_, _ok_ = _buf["Res"].([]interface{}); !_ok_ {
			err = errors.New("Res error")
			return
		}

		_v.Res = make([]int32, 0, len(_arr_))

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
			_v.Res = append(_v.Res, _list_v_)
		}
	}

	{
		var _ok_ bool
		if _v.BeginTime, _ok_ = _buf["BeginTime"].(string); !_ok_ {
			err = errors.New("BeginTime error")
			return
		}
	}
	{
		var _ok_ bool
		if _v.EndTime, _ok_ = _buf["EndTime"].(string); !_ok_ {
			err = errors.New("EndTime error")
			return
		}
	}
	{
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["ShopType"].(float64); !_ok_ {
			err = errors.New("ShopType error")
			return
		}
		_v.ShopType = int32(_tempNum_)
	}
	{
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["Display"].(float64); !_ok_ {
			err = errors.New("Display error")
			return
		}
		_v.Display = int32(_tempNum_)
	}
	return
}

func DeserializeShop(_buf map[string]interface{}) (*Shop, error) {
	v := &Shop{}
	if err := v.Deserialize(_buf); err == nil {
		return v, nil
	} else {
		return nil, err
	}
}
