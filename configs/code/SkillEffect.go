//------------------------------------------------------------------------------
// <auto-generated>
//     This code was generated by a tool.
//     Changes to this file may cause incorrect behavior and will be lost if
//     the code is regenerated.
// </auto-generated>
//------------------------------------------------------------------------------

package cfg

import "errors"

type SkillEffect struct {
	ID                  int32
	Name                string
	TargetType          []int32
	Range               []int32
	EffectDelay         int32
	Duration            int32
	Condition           []int32
	Interval            int32
	MoveType            int32
	Speed               int32
	SkillDamage         []float32
	StrikeDamage        int32
	ExtraCriticalDamage int32
	BuffID              []int32
	BuffPro             []int32
	EffectID            []int32
	EffectPro           []int32
	Screen              int32
}

const TypeId_SkillEffect = -795519198

func (*SkillEffect) GetTypeId() int32 {
	return -795519198
}

func (_v *SkillEffect) Deserialize(_buf map[string]interface{}) (err error) {
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
		var _arr_ []interface{}
		var _ok_ bool
		if _arr_, _ok_ = _buf["TargetType"].([]interface{}); !_ok_ {
			err = errors.New("TargetType error")
			return
		}

		_v.TargetType = make([]int32, 0, len(_arr_))

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
			_v.TargetType = append(_v.TargetType, _list_v_)
		}
	}

	{
		var _arr_ []interface{}
		var _ok_ bool
		if _arr_, _ok_ = _buf["Range"].([]interface{}); !_ok_ {
			err = errors.New("Range error")
			return
		}

		_v.Range = make([]int32, 0, len(_arr_))

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
			_v.Range = append(_v.Range, _list_v_)
		}
	}

	{
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["EffectDelay"].(float64); !_ok_ {
			err = errors.New("EffectDelay error")
			return
		}
		_v.EffectDelay = int32(_tempNum_)
	}
	{
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["Duration"].(float64); !_ok_ {
			err = errors.New("Duration error")
			return
		}
		_v.Duration = int32(_tempNum_)
	}
	{
		var _arr_ []interface{}
		var _ok_ bool
		if _arr_, _ok_ = _buf["Condition"].([]interface{}); !_ok_ {
			err = errors.New("Condition error")
			return
		}

		_v.Condition = make([]int32, 0, len(_arr_))

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
			_v.Condition = append(_v.Condition, _list_v_)
		}
	}

	{
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["Interval"].(float64); !_ok_ {
			err = errors.New("Interval error")
			return
		}
		_v.Interval = int32(_tempNum_)
	}
	{
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["MoveType"].(float64); !_ok_ {
			err = errors.New("MoveType error")
			return
		}
		_v.MoveType = int32(_tempNum_)
	}
	{
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["Speed"].(float64); !_ok_ {
			err = errors.New("Speed error")
			return
		}
		_v.Speed = int32(_tempNum_)
	}
	{
		var _arr_ []interface{}
		var _ok_ bool
		if _arr_, _ok_ = _buf["SkillDamage"].([]interface{}); !_ok_ {
			err = errors.New("SkillDamage error")
			return
		}

		_v.SkillDamage = make([]float32, 0, len(_arr_))

		for _, _e_ := range _arr_ {
			var _list_v_ float32
			{
				var _ok_ bool
				var _x_ float64
				if _x_, _ok_ = _e_.(float64); !_ok_ {
					err = errors.New("_list_v_ error")
					return
				}
				_list_v_ = float32(_x_)
			}
			_v.SkillDamage = append(_v.SkillDamage, _list_v_)
		}
	}

	{
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["StrikeDamage"].(float64); !_ok_ {
			err = errors.New("StrikeDamage error")
			return
		}
		_v.StrikeDamage = int32(_tempNum_)
	}
	{
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["ExtraCriticalDamage"].(float64); !_ok_ {
			err = errors.New("ExtraCriticalDamage error")
			return
		}
		_v.ExtraCriticalDamage = int32(_tempNum_)
	}
	{
		var _arr_ []interface{}
		var _ok_ bool
		if _arr_, _ok_ = _buf["BuffID"].([]interface{}); !_ok_ {
			err = errors.New("BuffID error")
			return
		}

		_v.BuffID = make([]int32, 0, len(_arr_))

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
			_v.BuffID = append(_v.BuffID, _list_v_)
		}
	}

	{
		var _arr_ []interface{}
		var _ok_ bool
		if _arr_, _ok_ = _buf["BuffPro"].([]interface{}); !_ok_ {
			err = errors.New("BuffPro error")
			return
		}

		_v.BuffPro = make([]int32, 0, len(_arr_))

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
			_v.BuffPro = append(_v.BuffPro, _list_v_)
		}
	}

	{
		var _arr_ []interface{}
		var _ok_ bool
		if _arr_, _ok_ = _buf["EffectID"].([]interface{}); !_ok_ {
			err = errors.New("EffectID error")
			return
		}

		_v.EffectID = make([]int32, 0, len(_arr_))

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
			_v.EffectID = append(_v.EffectID, _list_v_)
		}
	}

	{
		var _arr_ []interface{}
		var _ok_ bool
		if _arr_, _ok_ = _buf["EffectPro"].([]interface{}); !_ok_ {
			err = errors.New("EffectPro error")
			return
		}

		_v.EffectPro = make([]int32, 0, len(_arr_))

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
			_v.EffectPro = append(_v.EffectPro, _list_v_)
		}
	}

	{
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["Screen"].(float64); !_ok_ {
			err = errors.New("Screen error")
			return
		}
		_v.Screen = int32(_tempNum_)
	}
	return
}

func DeserializeSkillEffect(_buf map[string]interface{}) (*SkillEffect, error) {
	v := &SkillEffect{}
	if err := v.Deserialize(_buf); err == nil {
		return v, nil
	} else {
		return nil, err
	}
}
