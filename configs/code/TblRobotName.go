//------------------------------------------------------------------------------
// <auto-generated>
//     This code was generated by a tool.
//     Changes to this file may cause incorrect behavior and will be lost if
//     the code is regenerated.
// </auto-generated>
//------------------------------------------------------------------------------

package cfg

type TblRobotName struct {
	_dataMap  map[int32]*RobotName
	_dataList []*RobotName
}

func NewTblRobotName(_buf []map[string]interface{}) (*TblRobotName, error) {
	_dataList := make([]*RobotName, 0, len(_buf))
	dataMap := make(map[int32]*RobotName)
	for _, _ele_ := range _buf {
		if _v, err2 := DeserializeRobotName(_ele_); err2 != nil {
			return nil, err2
		} else {
			_dataList = append(_dataList, _v)
			dataMap[_v.ID] = _v
		}
	}
	return &TblRobotName{_dataList: _dataList, _dataMap: dataMap}, nil
}

func (table *TblRobotName) GetDataMap() map[int32]*RobotName {
	return table._dataMap
}

func (table *TblRobotName) GetDataList() []*RobotName {
	return table._dataList
}

func (table *TblRobotName) Get(key int32) *RobotName {
	return table._dataMap[key]
}
