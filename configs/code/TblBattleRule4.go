//------------------------------------------------------------------------------
// <auto-generated>
//     This code was generated by a tool.
//     Changes to this file may cause incorrect behavior and will be lost if
//     the code is regenerated.
// </auto-generated>
//------------------------------------------------------------------------------

package cfg

type TblBattleRule4 struct {
	_dataMap  map[int32]*BattleRule4
	_dataList []*BattleRule4
}

func NewTblBattleRule4(_buf []map[string]interface{}) (*TblBattleRule4, error) {
	_dataList := make([]*BattleRule4, 0, len(_buf))
	dataMap := make(map[int32]*BattleRule4)
	for _, _ele_ := range _buf {
		if _v, err2 := DeserializeBattleRule4(_ele_); err2 != nil {
			return nil, err2
		} else {
			_dataList = append(_dataList, _v)
			dataMap[_v.ID] = _v
		}
	}
	return &TblBattleRule4{_dataList: _dataList, _dataMap: dataMap}, nil
}

func (table *TblBattleRule4) GetDataMap() map[int32]*BattleRule4 {
	return table._dataMap
}

func (table *TblBattleRule4) GetDataList() []*BattleRule4 {
	return table._dataList
}

func (table *TblBattleRule4) Get(key int32) *BattleRule4 {
	return table._dataMap[key]
}
