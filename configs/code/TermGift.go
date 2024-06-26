//------------------------------------------------------------------------------
// <auto-generated>
//     This code was generated by a tool.
//     Changes to this file may cause incorrect behavior and will be lost if
//     the code is regenerated.
// </auto-generated>
//------------------------------------------------------------------------------

package cfg

import "errors"

type TermGift struct {
	ID         int32
	Name       string
	Res        string
	Type       int32
	SaleOff    int32
	PurChaseID int32
	ItemBoxID  int32
	LimitTime  int32
	CD         int32
	TermLv1    int32
	TermMun1   []int32
	TermLv2    int32
	TermMun2   []int32
	TermLv3    int32
	TermMun3   []int32
}

const TypeId_TermGift = -1117622308

func (*TermGift) GetTypeId() int32 {
	return -1117622308
}

func (_v *TermGift) Deserialize(_buf map[string]interface{}) (err error) {
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
		if _v.Res, _ok_ = _buf["Res"].(string); !_ok_ {
			err = errors.New("Res error")
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
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["SaleOff"].(float64); !_ok_ {
			err = errors.New("SaleOff error")
			return
		}
		_v.SaleOff = int32(_tempNum_)
	}
	{
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["PurChaseID"].(float64); !_ok_ {
			err = errors.New("PurChaseID error")
			return
		}
		_v.PurChaseID = int32(_tempNum_)
	}
	{
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["ItemBoxID"].(float64); !_ok_ {
			err = errors.New("ItemBoxID error")
			return
		}
		_v.ItemBoxID = int32(_tempNum_)
	}
	{
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["LimitTime"].(float64); !_ok_ {
			err = errors.New("LimitTime error")
			return
		}
		_v.LimitTime = int32(_tempNum_)
	}
	{
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["CD"].(float64); !_ok_ {
			err = errors.New("CD error")
			return
		}
		_v.CD = int32(_tempNum_)
	}
	{
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["TermLv1"].(float64); !_ok_ {
			err = errors.New("TermLv1 error")
			return
		}
		_v.TermLv1 = int32(_tempNum_)
	}
	{
		var _arr_ []interface{}
		var _ok_ bool
		if _arr_, _ok_ = _buf["TermMun1"].([]interface{}); !_ok_ {
			err = errors.New("TermMun1 error")
			return
		}

		_v.TermMun1 = make([]int32, 0, len(_arr_))

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
			_v.TermMun1 = append(_v.TermMun1, _list_v_)
		}
	}

	{
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["TermLv2"].(float64); !_ok_ {
			err = errors.New("TermLv2 error")
			return
		}
		_v.TermLv2 = int32(_tempNum_)
	}
	{
		var _arr_ []interface{}
		var _ok_ bool
		if _arr_, _ok_ = _buf["TermMun2"].([]interface{}); !_ok_ {
			err = errors.New("TermMun2 error")
			return
		}

		_v.TermMun2 = make([]int32, 0, len(_arr_))

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
			_v.TermMun2 = append(_v.TermMun2, _list_v_)
		}
	}

	{
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["TermLv3"].(float64); !_ok_ {
			err = errors.New("TermLv3 error")
			return
		}
		_v.TermLv3 = int32(_tempNum_)
	}
	{
		var _arr_ []interface{}
		var _ok_ bool
		if _arr_, _ok_ = _buf["TermMun3"].([]interface{}); !_ok_ {
			err = errors.New("TermMun3 error")
			return
		}

		_v.TermMun3 = make([]int32, 0, len(_arr_))

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
			_v.TermMun3 = append(_v.TermMun3, _list_v_)
		}
	}

	return
}

func DeserializeTermGift(_buf map[string]interface{}) (*TermGift, error) {
	v := &TermGift{}
	if err := v.Deserialize(_buf); err == nil {
		return v, nil
	} else {
		return nil, err
	}
}
