//------------------------------------------------------------------------------
// <auto-generated>
//     This code was generated by a tool.
//     Changes to this file may cause incorrect behavior and will be lost if
//     the code is regenerated.
// </auto-generated>
//------------------------------------------------------------------------------

package cfg

type TblCumulativeLogin struct {
	_dataMap  map[int32]*CumulativeLogin
	_dataList []*CumulativeLogin
}

func NewTblCumulativeLogin(_buf []map[string]interface{}) (*TblCumulativeLogin, error) {
	_dataList := make([]*CumulativeLogin, 0, len(_buf))
	dataMap := make(map[int32]*CumulativeLogin)
	for _, _ele_ := range _buf {
		if _v, err2 := DeserializeCumulativeLogin(_ele_); err2 != nil {
			return nil, err2
		} else {
			_dataList = append(_dataList, _v)
			dataMap[_v.ID] = _v
		}
	}
	return &TblCumulativeLogin{_dataList: _dataList, _dataMap: dataMap}, nil
}

func (table *TblCumulativeLogin) GetDataMap() map[int32]*CumulativeLogin {
	return table._dataMap
}

func (table *TblCumulativeLogin) GetDataList() []*CumulativeLogin {
	return table._dataList
}

func (table *TblCumulativeLogin) Get(key int32) *CumulativeLogin {
	return table._dataMap[key]
}
