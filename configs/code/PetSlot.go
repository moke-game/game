//------------------------------------------------------------------------------
// <auto-generated>
//     This code was generated by a tool.
//     Changes to this file may cause incorrect behavior and will be lost if
//     the code is regenerated.
// </auto-generated>
//------------------------------------------------------------------------------

package cfg

import "errors"

type PetSlot struct {
	ID      int32
	CostID  int32
	CostNum int32
}

const TypeId_PetSlot = 987103133

func (*PetSlot) GetTypeId() int32 {
	return 987103133
}

func (_v *PetSlot) Deserialize(_buf map[string]interface{}) (err error) {
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
		if _tempNum_, _ok_ = _buf["CostID"].(float64); !_ok_ {
			err = errors.New("CostID error")
			return
		}
		_v.CostID = int32(_tempNum_)
	}
	{
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["CostNum"].(float64); !_ok_ {
			err = errors.New("CostNum error")
			return
		}
		_v.CostNum = int32(_tempNum_)
	}
	return
}

func DeserializePetSlot(_buf map[string]interface{}) (*PetSlot, error) {
	v := &PetSlot{}
	if err := v.Deserialize(_buf); err == nil {
		return v, nil
	} else {
		return nil, err
	}
}
