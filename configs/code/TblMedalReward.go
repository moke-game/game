//------------------------------------------------------------------------------
// <auto-generated>
//     This code was generated by a tool.
//     Changes to this file may cause incorrect behavior and will be lost if
//     the code is regenerated.
// </auto-generated>
//------------------------------------------------------------------------------

package cfg

type TblMedalReward struct {
	_dataMap  map[int32]*MedalReward
	_dataList []*MedalReward
}

func NewTblMedalReward(_buf []map[string]interface{}) (*TblMedalReward, error) {
	_dataList := make([]*MedalReward, 0, len(_buf))
	dataMap := make(map[int32]*MedalReward)
	for _, _ele_ := range _buf {
		if _v, err2 := DeserializeMedalReward(_ele_); err2 != nil {
			return nil, err2
		} else {
			_dataList = append(_dataList, _v)
			dataMap[_v.ID] = _v
		}
	}
	return &TblMedalReward{_dataList: _dataList, _dataMap: dataMap}, nil
}

func (table *TblMedalReward) GetDataMap() map[int32]*MedalReward {
	return table._dataMap
}

func (table *TblMedalReward) GetDataList() []*MedalReward {
	return table._dataList
}

func (table *TblMedalReward) Get(key int32) *MedalReward {
	return table._dataMap[key]
}