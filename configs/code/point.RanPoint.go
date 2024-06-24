//------------------------------------------------------------------------------
// <auto-generated>
//     This code was generated by a tool.
//     Changes to this file may cause incorrect behavior and will be lost if
//     the code is regenerated.
// </auto-generated>
//------------------------------------------------------------------------------

package cfg

import "errors"

type PointRanPoint struct {
	Camp   int32
	X      float32
	Y      float32
	Z      float32
	Radius int32
	FaceTo int32
}

const TypeId_PointRanPoint = -986043537

func (*PointRanPoint) GetTypeId() int32 {
	return -986043537
}

func (_v *PointRanPoint) Deserialize(_buf map[string]interface{}) (err error) {
	{
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["camp"].(float64); !_ok_ {
			err = errors.New("camp error")
			return
		}
		_v.Camp = int32(_tempNum_)
	}
	{
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["x"].(float64); !_ok_ {
			err = errors.New("x error")
			return
		}
		_v.X = float32(_tempNum_)
	}
	{
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["y"].(float64); !_ok_ {
			err = errors.New("y error")
			return
		}
		_v.Y = float32(_tempNum_)
	}
	{
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["z"].(float64); !_ok_ {
			err = errors.New("z error")
			return
		}
		_v.Z = float32(_tempNum_)
	}
	{
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["radius"].(float64); !_ok_ {
			err = errors.New("radius error")
			return
		}
		_v.Radius = int32(_tempNum_)
	}
	{
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["faceTo"].(float64); !_ok_ {
			err = errors.New("faceTo error")
			return
		}
		_v.FaceTo = int32(_tempNum_)
	}
	return
}

func DeserializePointRanPoint(_buf map[string]interface{}) (*PointRanPoint, error) {
	v := &PointRanPoint{}
	if err := v.Deserialize(_buf); err == nil {
		return v, nil
	} else {
		return nil, err
	}
}
