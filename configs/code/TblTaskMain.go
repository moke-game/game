//------------------------------------------------------------------------------
// <auto-generated>
//     This code was generated by a tool.
//     Changes to this file may cause incorrect behavior and will be lost if
//     the code is regenerated.
// </auto-generated>
//------------------------------------------------------------------------------

package cfg

type TblTaskMain struct {
	_dataMap  map[int32]*TaskMain
	_dataList []*TaskMain
}

func NewTblTaskMain(_buf []map[string]interface{}) (*TblTaskMain, error) {
	_dataList := make([]*TaskMain, 0, len(_buf))
	dataMap := make(map[int32]*TaskMain)
	for _, _ele_ := range _buf {
		if _v, err2 := DeserializeTaskMain(_ele_); err2 != nil {
			return nil, err2
		} else {
			_dataList = append(_dataList, _v)
			dataMap[_v.ID] = _v
		}
	}
	return &TblTaskMain{_dataList: _dataList, _dataMap: dataMap}, nil
}

func (table *TblTaskMain) GetDataMap() map[int32]*TaskMain {
	return table._dataMap
}

func (table *TblTaskMain) GetDataList() []*TaskMain {
	return table._dataList
}

func (table *TblTaskMain) Get(key int32) *TaskMain {
	return table._dataMap[key]
}
