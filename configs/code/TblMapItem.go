//------------------------------------------------------------------------------
// <auto-generated>
//     This code was generated by a tool.
//     Changes to this file may cause incorrect behavior and will be lost if
//     the code is regenerated.
// </auto-generated>
//------------------------------------------------------------------------------

package cfg

type TblMapItem struct {
	_dataMap  map[int32]*MapItem
	_dataList []*MapItem
}

func NewTblMapItem(_buf []map[string]interface{}) (*TblMapItem, error) {
	_dataList := make([]*MapItem, 0, len(_buf))
	dataMap := make(map[int32]*MapItem)
	for _, _ele_ := range _buf {
		if _v, err2 := DeserializeMapItem(_ele_); err2 != nil {
			return nil, err2
		} else {
			_dataList = append(_dataList, _v)
			dataMap[_v.ID] = _v
		}
	}
	return &TblMapItem{_dataList: _dataList, _dataMap: dataMap}, nil
}

func (table *TblMapItem) GetDataMap() map[int32]*MapItem {
	return table._dataMap
}

func (table *TblMapItem) GetDataList() []*MapItem {
	return table._dataList
}

func (table *TblMapItem) Get(key int32) *MapItem {
	return table._dataMap[key]
}
