//------------------------------------------------------------------------------
// <auto-generated>
//     This code was generated by a tool.
//     Changes to this file may cause incorrect behavior and will be lost if
//     the code is regenerated.
// </auto-generated>
//------------------------------------------------------------------------------

package cfg

import "errors"

type ItemDropItemPool struct {
	Typ   int32
	Items []*ItemDropItem
}

const TypeId_ItemDropItemPool = 858248441

func (*ItemDropItemPool) GetTypeId() int32 {
	return 858248441
}

func (_v *ItemDropItemPool) Deserialize(_buf map[string]interface{}) (err error) {
	{
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["typ"].(float64); !_ok_ {
			err = errors.New("typ error")
			return
		}
		_v.Typ = int32(_tempNum_)
	}
	{
		var _arr_ []interface{}
		var _ok_ bool
		if _arr_, _ok_ = _buf["Items"].([]interface{}); !_ok_ {
			err = errors.New("Items error")
			return
		}

		_v.Items = make([]*ItemDropItem, 0, len(_arr_))

		for _, _e_ := range _arr_ {
			var _list_v_ *ItemDropItem
			{
				var _ok_ bool
				var _x_ map[string]interface{}
				if _x_, _ok_ = _e_.(map[string]interface{}); !_ok_ {
					err = errors.New("_list_v_ error")
					return
				}
				if _list_v_, err = DeserializeItemDropItem(_x_); err != nil {
					return
				}
			}
			_v.Items = append(_v.Items, _list_v_)
		}
	}

	return
}

func DeserializeItemDropItemPool(_buf map[string]interface{}) (*ItemDropItemPool, error) {
	v := &ItemDropItemPool{}
	if err := v.Deserialize(_buf); err == nil {
		return v, nil
	} else {
		return nil, err
	}
}
