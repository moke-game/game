//------------------------------------------------------------------------------
// <auto-generated>
//     This code was generated by a tool.
//     Changes to this file may cause incorrect behavior and will be lost if
//     the code is regenerated.
// </auto-generated>
//------------------------------------------------------------------------------

package cfg

type TblDailyTask struct {
	_dataMap  map[int32]*DailyTask
	_dataList []*DailyTask
}

func NewTblDailyTask(_buf []map[string]interface{}) (*TblDailyTask, error) {
	_dataList := make([]*DailyTask, 0, len(_buf))
	dataMap := make(map[int32]*DailyTask)
	for _, _ele_ := range _buf {
		if _v, err2 := DeserializeDailyTask(_ele_); err2 != nil {
			return nil, err2
		} else {
			_dataList = append(_dataList, _v)
			dataMap[_v.ID] = _v
		}
	}
	return &TblDailyTask{_dataList: _dataList, _dataMap: dataMap}, nil
}

func (table *TblDailyTask) GetDataMap() map[int32]*DailyTask {
	return table._dataMap
}

func (table *TblDailyTask) GetDataList() []*DailyTask {
	return table._dataList
}

func (table *TblDailyTask) Get(key int32) *DailyTask {
	return table._dataMap[key]
}