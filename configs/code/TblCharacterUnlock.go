//------------------------------------------------------------------------------
// <auto-generated>
//     This code was generated by a tool.
//     Changes to this file may cause incorrect behavior and will be lost if
//     the code is regenerated.
// </auto-generated>
//------------------------------------------------------------------------------

package cfg

type TblCharacterUnlock struct {
	_dataMap  map[int32]*CharacterUnlock
	_dataList []*CharacterUnlock
}

func NewTblCharacterUnlock(_buf []map[string]interface{}) (*TblCharacterUnlock, error) {
	_dataList := make([]*CharacterUnlock, 0, len(_buf))
	dataMap := make(map[int32]*CharacterUnlock)
	for _, _ele_ := range _buf {
		if _v, err2 := DeserializeCharacterUnlock(_ele_); err2 != nil {
			return nil, err2
		} else {
			_dataList = append(_dataList, _v)
			dataMap[_v.ID] = _v
		}
	}
	return &TblCharacterUnlock{_dataList: _dataList, _dataMap: dataMap}, nil
}

func (table *TblCharacterUnlock) GetDataMap() map[int32]*CharacterUnlock {
	return table._dataMap
}

func (table *TblCharacterUnlock) GetDataList() []*CharacterUnlock {
	return table._dataList
}

func (table *TblCharacterUnlock) Get(key int32) *CharacterUnlock {
	return table._dataMap[key]
}
