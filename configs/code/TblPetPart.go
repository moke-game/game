//------------------------------------------------------------------------------
// <auto-generated>
//     This code was generated by a tool.
//     Changes to this file may cause incorrect behavior and will be lost if
//     the code is regenerated.
// </auto-generated>
//------------------------------------------------------------------------------

package cfg

type TblPetPart struct {
	_dataMap  map[int32]*PetPart
	_dataList []*PetPart
}

func NewTblPetPart(_buf []map[string]interface{}) (*TblPetPart, error) {
	_dataList := make([]*PetPart, 0, len(_buf))
	dataMap := make(map[int32]*PetPart)
	for _, _ele_ := range _buf {
		if _v, err2 := DeserializePetPart(_ele_); err2 != nil {
			return nil, err2
		} else {
			_dataList = append(_dataList, _v)
			dataMap[_v.ID] = _v
		}
	}
	return &TblPetPart{_dataList: _dataList, _dataMap: dataMap}, nil
}

func (table *TblPetPart) GetDataMap() map[int32]*PetPart {
	return table._dataMap
}

func (table *TblPetPart) GetDataList() []*PetPart {
	return table._dataList
}

func (table *TblPetPart) Get(key int32) *PetPart {
	return table._dataMap[key]
}
