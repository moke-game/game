//------------------------------------------------------------------------------
// <auto-generated>
//     This code was generated by a tool.
//     Changes to this file may cause incorrect behavior and will be lost if
//     the code is regenerated.
// </auto-generated>
//------------------------------------------------------------------------------

package cfg

import "errors"

type PetShape struct {
	ID      int32
	ResName []string
}

const TypeId_PetShape = 535293410

func (*PetShape) GetTypeId() int32 {
	return 535293410
}

func (_v *PetShape) Deserialize(_buf map[string]interface{}) (err error) {
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

	return
}

func DeserializePetShape(_buf map[string]interface{}) (*PetShape, error) {
	v := &PetShape{}
	if err := v.Deserialize(_buf); err == nil {
		return v, nil
	} else {
		return nil, err
	}
}
