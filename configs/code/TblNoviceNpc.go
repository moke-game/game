//------------------------------------------------------------------------------
// <auto-generated>
//     This code was generated by a tool.
//     Changes to this file may cause incorrect behavior and will be lost if
//     the code is regenerated.
// </auto-generated>
//------------------------------------------------------------------------------

package cfg

type TblNoviceNpc struct {
	_dataMap  map[int32]*NoviceNpc
	_dataList []*NoviceNpc
}

func NewTblNoviceNpc(_buf []map[string]interface{}) (*TblNoviceNpc, error) {
	_dataList := make([]*NoviceNpc, 0, len(_buf))
	dataMap := make(map[int32]*NoviceNpc)
	for _, _ele_ := range _buf {
		if _v, err2 := DeserializeNoviceNpc(_ele_); err2 != nil {
			return nil, err2
		} else {
			_dataList = append(_dataList, _v)
			dataMap[_v.ID] = _v
		}
	}
	return &TblNoviceNpc{_dataList: _dataList, _dataMap: dataMap}, nil
}

func (table *TblNoviceNpc) GetDataMap() map[int32]*NoviceNpc {
	return table._dataMap
}

func (table *TblNoviceNpc) GetDataList() []*NoviceNpc {
	return table._dataList
}

func (table *TblNoviceNpc) Get(key int32) *NoviceNpc {
	return table._dataMap[key]
}
