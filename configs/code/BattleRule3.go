//------------------------------------------------------------------------------
// <auto-generated>
//     This code was generated by a tool.
//     Changes to this file may cause incorrect behavior and will be lost if
//     the code is regenerated.
// </auto-generated>
//------------------------------------------------------------------------------

package cfg

import "errors"

type BattleRule3 struct {
	ID               int32
	Mode             int32
	PlayerExp1       int32
	PlayerExp2       int32
	SkillExp1        int32
	SkillExp2        int32
	Gold1            int32
	Gold2            int32
	GoldContinue     int32
	GoldMvp          int32
	Time             int32
	ReviveTimes      int32
	ReviveWait       int32
	Tips1            string
	Tips2            string
	Tips3            string
	Tips4            string
	FirstAwardLowest int32
	TeamMember       int32
}

const TypeId_BattleRule3 = 906584767

func (*BattleRule3) GetTypeId() int32 {
	return 906584767
}

func (_v *BattleRule3) Deserialize(_buf map[string]interface{}) (err error) {
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
		if _tempNum_, _ok_ = _buf["Mode"].(float64); !_ok_ {
			err = errors.New("Mode error")
			return
		}
		_v.Mode = int32(_tempNum_)
	}
	{
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["PlayerExp1"].(float64); !_ok_ {
			err = errors.New("PlayerExp1 error")
			return
		}
		_v.PlayerExp1 = int32(_tempNum_)
	}
	{
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["PlayerExp2"].(float64); !_ok_ {
			err = errors.New("PlayerExp2 error")
			return
		}
		_v.PlayerExp2 = int32(_tempNum_)
	}
	{
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["SkillExp1"].(float64); !_ok_ {
			err = errors.New("SkillExp1 error")
			return
		}
		_v.SkillExp1 = int32(_tempNum_)
	}
	{
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["SkillExp2"].(float64); !_ok_ {
			err = errors.New("SkillExp2 error")
			return
		}
		_v.SkillExp2 = int32(_tempNum_)
	}
	{
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["Gold1"].(float64); !_ok_ {
			err = errors.New("Gold1 error")
			return
		}
		_v.Gold1 = int32(_tempNum_)
	}
	{
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["Gold2"].(float64); !_ok_ {
			err = errors.New("Gold2 error")
			return
		}
		_v.Gold2 = int32(_tempNum_)
	}
	{
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["GoldContinue"].(float64); !_ok_ {
			err = errors.New("GoldContinue error")
			return
		}
		_v.GoldContinue = int32(_tempNum_)
	}
	{
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["GoldMvp"].(float64); !_ok_ {
			err = errors.New("GoldMvp error")
			return
		}
		_v.GoldMvp = int32(_tempNum_)
	}
	{
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["Time"].(float64); !_ok_ {
			err = errors.New("Time error")
			return
		}
		_v.Time = int32(_tempNum_)
	}
	{
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["ReviveTimes"].(float64); !_ok_ {
			err = errors.New("ReviveTimes error")
			return
		}
		_v.ReviveTimes = int32(_tempNum_)
	}
	{
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["ReviveWait"].(float64); !_ok_ {
			err = errors.New("ReviveWait error")
			return
		}
		_v.ReviveWait = int32(_tempNum_)
	}
	{
		var _ok_ bool
		if _v.Tips1, _ok_ = _buf["Tips1"].(string); !_ok_ {
			err = errors.New("Tips1 error")
			return
		}
	}
	{
		var _ok_ bool
		if _v.Tips2, _ok_ = _buf["Tips2"].(string); !_ok_ {
			err = errors.New("Tips2 error")
			return
		}
	}
	{
		var _ok_ bool
		if _v.Tips3, _ok_ = _buf["Tips3"].(string); !_ok_ {
			err = errors.New("Tips3 error")
			return
		}
	}
	{
		var _ok_ bool
		if _v.Tips4, _ok_ = _buf["Tips4"].(string); !_ok_ {
			err = errors.New("Tips4 error")
			return
		}
	}
	{
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["FirstAwardLowest"].(float64); !_ok_ {
			err = errors.New("FirstAwardLowest error")
			return
		}
		_v.FirstAwardLowest = int32(_tempNum_)
	}
	{
		var _ok_ bool
		var _tempNum_ float64
		if _tempNum_, _ok_ = _buf["TeamMember"].(float64); !_ok_ {
			err = errors.New("TeamMember error")
			return
		}
		_v.TeamMember = int32(_tempNum_)
	}
	return
}

func DeserializeBattleRule3(_buf map[string]interface{}) (*BattleRule3, error) {
	v := &BattleRule3{}
	if err := v.Deserialize(_buf); err == nil {
		return v, nil
	} else {
		return nil, err
	}
}
