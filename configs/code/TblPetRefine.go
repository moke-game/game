//------------------------------------------------------------------------------
// <auto-generated>
//     This code was generated by a tool.
//     Changes to this file may cause incorrect behavior and will be lost if
//     the code is regenerated.
// </auto-generated>
//------------------------------------------------------------------------------

package cfg

type TblPetRefine struct {
	_dataMap  map[int32]*PetRefine
	_dataList []*PetRefine
}

func NewTblPetRefine(_buf []map[string]interface{}) (*TblPetRefine, error) {
	_dataList := make([]*PetRefine, 0, len(_buf))
	dataMap := make(map[int32]*PetRefine)
	for _, _ele_ := range _buf {
		if _v, err2 := DeserializePetRefine(_ele_); err2 != nil {
			return nil, err2
		} else {
			_dataList = append(_dataList, _v)
			dataMap[_v.ID] = _v
		}
	}
	return &TblPetRefine{_dataList: _dataList, _dataMap: dataMap}, nil
}

func (table *TblPetRefine) GetDataMap() map[int32]*PetRefine {
	return table._dataMap
}

func (table *TblPetRefine) GetDataList() []*PetRefine {
	return table._dataList
}

func (table *TblPetRefine) Get(key int32) *PetRefine {
	return table._dataMap[key]
}
