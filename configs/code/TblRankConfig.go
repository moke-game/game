//------------------------------------------------------------------------------
// <auto-generated>
//     This code was generated by a tool.
//     Changes to this file may cause incorrect behavior and will be lost if
//     the code is regenerated.
// </auto-generated>
//------------------------------------------------------------------------------

package cfg

type TblRankConfig struct {
	_dataMap  map[int32]*RankConfig
	_dataList []*RankConfig
}

func NewTblRankConfig(_buf []map[string]interface{}) (*TblRankConfig, error) {
	_dataList := make([]*RankConfig, 0, len(_buf))
	dataMap := make(map[int32]*RankConfig)
	for _, _ele_ := range _buf {
		if _v, err2 := DeserializeRankConfig(_ele_); err2 != nil {
			return nil, err2
		} else {
			_dataList = append(_dataList, _v)
			dataMap[_v.ID] = _v
		}
	}
	return &TblRankConfig{_dataList: _dataList, _dataMap: dataMap}, nil
}

func (table *TblRankConfig) GetDataMap() map[int32]*RankConfig {
	return table._dataMap
}

func (table *TblRankConfig) GetDataList() []*RankConfig {
	return table._dataList
}

func (table *TblRankConfig) Get(key int32) *RankConfig {
	return table._dataMap[key]
}
