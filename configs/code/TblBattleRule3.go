//------------------------------------------------------------------------------
// <auto-generated>
//     This code was generated by a tool.
//     Changes to this file may cause incorrect behavior and will be lost if
//     the code is regenerated.
// </auto-generated>
//------------------------------------------------------------------------------

package cfg

type TblBattleRule3 struct {
	_dataMap  map[int32]*BattleRule3
	_dataList []*BattleRule3
}

func NewTblBattleRule3(_buf []map[string]interface{}) (*TblBattleRule3, error) {
	_dataList := make([]*BattleRule3, 0, len(_buf))
	dataMap := make(map[int32]*BattleRule3)
	for _, _ele_ := range _buf {
		if _v, err2 := DeserializeBattleRule3(_ele_); err2 != nil {
			return nil, err2
		} else {
			_dataList = append(_dataList, _v)
			dataMap[_v.ID] = _v
		}
	}
	return &TblBattleRule3{_dataList: _dataList, _dataMap: dataMap}, nil
}

func (table *TblBattleRule3) GetDataMap() map[int32]*BattleRule3 {
	return table._dataMap
}

func (table *TblBattleRule3) GetDataList() []*BattleRule3 {
	return table._dataList
}

func (table *TblBattleRule3) Get(key int32) *BattleRule3 {
	return table._dataMap[key]
}
