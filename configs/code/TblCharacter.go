//------------------------------------------------------------------------------
// <auto-generated>
//     This code was generated by a tool.
//     Changes to this file may cause incorrect behavior and will be lost if
//     the code is regenerated.
// </auto-generated>
//------------------------------------------------------------------------------

package cfg

type TblCharacter struct {
	_dataMap  map[int32]*Character
	_dataList []*Character
}

func NewTblCharacter(_buf []map[string]interface{}) (*TblCharacter, error) {
	_dataList := make([]*Character, 0, len(_buf))
	dataMap := make(map[int32]*Character)
	for _, _ele_ := range _buf {
		if _v, err2 := DeserializeCharacter(_ele_); err2 != nil {
			return nil, err2
		} else {
			_dataList = append(_dataList, _v)
			dataMap[_v.ID] = _v
		}
	}
	return &TblCharacter{_dataList: _dataList, _dataMap: dataMap}, nil
}

func (table *TblCharacter) GetDataMap() map[int32]*Character {
	return table._dataMap
}

func (table *TblCharacter) GetDataList() []*Character {
	return table._dataList
}

func (table *TblCharacter) Get(key int32) *Character {
	return table._dataMap[key]
}