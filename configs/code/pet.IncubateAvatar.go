//------------------------------------------------------------------------------
// <auto-generated>
//     This code was generated by a tool.
//     Changes to this file may cause incorrect behavior and will be lost if
//     the code is regenerated.
// </auto-generated>
//------------------------------------------------------------------------------

package cfg

import "errors"

type PetIncubateAvatar struct {
	PlanetId int32
	Rate     int32
	RandomId int32
}

const TypeId_PetIncubateAvatar = 832689359

func (*PetIncubateAvatar) GetTypeId() int32 {
	return 832689359
}

func (_v *PetIncubateAvatar) Deserialize(_buf map[string]interface{}) (err error) {
	{
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["planetId"].(float64); !_ok_ {
			err = errors.New("planetId error")
			return
		}
		_v.PlanetId = int32(_tempNum_)
	}
	{
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["rate"].(float64); !_ok_ {
			err = errors.New("rate error")
			return
		}
		_v.Rate = int32(_tempNum_)
	}
	{
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["randomId"].(float64); !_ok_ {
			err = errors.New("randomId error")
			return
		}
		_v.RandomId = int32(_tempNum_)
	}
	return
}

func DeserializePetIncubateAvatar(_buf map[string]interface{}) (*PetIncubateAvatar, error) {
	v := &PetIncubateAvatar{}
	if err := v.Deserialize(_buf); err == nil {
		return v, nil
	} else {
		return nil, err
	}
}
