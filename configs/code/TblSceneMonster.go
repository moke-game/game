//------------------------------------------------------------------------------
// <auto-generated>
//     This code was generated by a tool.
//     Changes to this file may cause incorrect behavior and will be lost if
//     the code is regenerated.
// </auto-generated>
//------------------------------------------------------------------------------

package cfg

type TblSceneMonster struct {
	_dataMap  map[int32]*SceneMonster
	_dataList []*SceneMonster
}

func NewTblSceneMonster(_buf []map[string]interface{}) (*TblSceneMonster, error) {
	_dataList := make([]*SceneMonster, 0, len(_buf))
	dataMap := make(map[int32]*SceneMonster)
	for _, _ele_ := range _buf {
		if _v, err2 := DeserializeSceneMonster(_ele_); err2 != nil {
			return nil, err2
		} else {
			_dataList = append(_dataList, _v)
			dataMap[_v.ID] = _v
		}
	}
	return &TblSceneMonster{_dataList: _dataList, _dataMap: dataMap}, nil
}

func (table *TblSceneMonster) GetDataMap() map[int32]*SceneMonster {
	return table._dataMap
}

func (table *TblSceneMonster) GetDataList() []*SceneMonster {
	return table._dataList
}

func (table *TblSceneMonster) Get(key int32) *SceneMonster {
	return table._dataMap[key]
}