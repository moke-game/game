//------------------------------------------------------------------------------
// <auto-generated>
//     This code was generated by a tool.
//     Changes to this file may cause incorrect behavior and will be lost if
//     the code is regenerated.
// </auto-generated>
//------------------------------------------------------------------------------

package cfg

type TblMonster struct {
	_dataMap  map[int32]*Monster
	_dataList []*Monster
}

func NewTblMonster(_buf []map[string]interface{}) (*TblMonster, error) {
	_dataList := make([]*Monster, 0, len(_buf))
	dataMap := make(map[int32]*Monster)
	for _, _ele_ := range _buf {
		if _v, err2 := DeserializeMonster(_ele_); err2 != nil {
			return nil, err2
		} else {
			_dataList = append(_dataList, _v)
			dataMap[_v.ID] = _v
		}
	}
	return &TblMonster{_dataList: _dataList, _dataMap: dataMap}, nil
}

func (table *TblMonster) GetDataMap() map[int32]*Monster {
	return table._dataMap
}

func (table *TblMonster) GetDataList() []*Monster {
	return table._dataList
}

func (table *TblMonster) Get(key int32) *Monster {
	return table._dataMap[key]
}
