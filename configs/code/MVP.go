//------------------------------------------------------------------------------
// <auto-generated>
//     This code was generated by a tool.
//     Changes to this file may cause incorrect behavior and will be lost if
//     the code is regenerated.
// </auto-generated>
//------------------------------------------------------------------------------

package cfg

import "errors"

type MVP struct {
	ID      int32
	Kill    int32
	Dead    int32
	Assist  int32
	Goal    int32
	Jewel   int32
	GemPrem int32
	GemNorm int32
}

const TypeId_MVP = 76743

func (*MVP) GetTypeId() int32 {
	return 76743
}

func (_v *MVP) Deserialize(_buf map[string]interface{}) (err error) {
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
		if _tempNum_, _ok_ = _buf["Kill"].(float64); !_ok_ {
			err = errors.New("Kill error")
			return
		}
		_v.Kill = int32(_tempNum_)
	}
	{
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["Dead"].(float64); !_ok_ {
			err = errors.New("Dead error")
			return
		}
		_v.Dead = int32(_tempNum_)
	}
	{
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["Assist"].(float64); !_ok_ {
			err = errors.New("Assist error")
			return
		}
		_v.Assist = int32(_tempNum_)
	}
	{
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["Goal"].(float64); !_ok_ {
			err = errors.New("Goal error")
			return
		}
		_v.Goal = int32(_tempNum_)
	}
	{
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["Jewel"].(float64); !_ok_ {
			err = errors.New("Jewel error")
			return
		}
		_v.Jewel = int32(_tempNum_)
	}
	{
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["GemPrem"].(float64); !_ok_ {
			err = errors.New("GemPrem error")
			return
		}
		_v.GemPrem = int32(_tempNum_)
	}
	{
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["GemNorm"].(float64); !_ok_ {
			err = errors.New("GemNorm error")
			return
		}
		_v.GemNorm = int32(_tempNum_)
	}
	return
}

func DeserializeMVP(_buf map[string]interface{}) (*MVP, error) {
	v := &MVP{}
	if err := v.Deserialize(_buf); err == nil {
		return v, nil
	} else {
		return nil, err
	}
}
