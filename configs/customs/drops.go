package customs

import "github.com/duke-git/lancet/v2/random"

// DropItemPool 掉落物品池
type DropItemPool struct {
	dropType uint
	entries  []*DropItemEntry
}

func (d *DropItemPool) DropType() uint {
	return d.dropType
}

func (d *DropItemPool) Entries() []*DropItemEntry {
	return d.entries
}

// DropItemEntry 掉落物品实体
type DropItemEntry struct {
	itemId   int32 // 物品id
	min, max int32 // 数量范围
	weight   int32 // 权重
}

func NewDropItemPool(dropType uint, entries []*DropItemEntry) *DropItemPool {
	return &DropItemPool{
		dropType: dropType,
		entries:  entries,
	}
}

func (d *DropItemPool) Drops() [][]int32 {
	ret := make([][]int32, 0)
	if d.GetDropType() == 0 {
		// 固定掉落
		for _, entry := range d.Entries() {
			itemId := entry.GetItemId()
			itemCount := random.RandInt(int(entry.Min()), int(entry.Max()))
			ret = append(ret, []int32{itemId, int32(itemCount)})
		}
	} else if d.GetDropType() == 1 {
		itemPool := make([]IRandomWeightItem, 0, len(d.Entries()))
		for _, entry := range d.Entries() {
			item := NewRandomWeightItem(entry, entry.Weight())
			itemPool = append(itemPool, item)
		}
		if item, ok := RandomWeight(itemPool); ok {
			if en, ok := item.(*DropItemEntry); ok {
				itemCount := random.RandInt(int(en.Min()), int(en.Max()))
				ret = append(ret, []int32{en.ItemId(), int32(itemCount)})
			}
		}
	}
	return ret
}

func NewDropItemEntry(itemId, min, max, weight int32) *DropItemEntry {
	return &DropItemEntry{
		itemId: itemId,
		min:    min,
		max:    max,
		weight: weight,
	}
}

func (d *DropItemPool) GetDropType() uint {
	return d.dropType
}

func (d *DropItemPool) GetEntries() []*DropItemEntry {
	return d.entries
}

func (d *DropItemEntry) GetItemId() int32 {
	return d.itemId
}

func (d *DropItemEntry) ItemId() int32 {
	return d.itemId
}

func (d *DropItemEntry) SetItemId(itemId int32) {
	d.itemId = itemId
}

func (d *DropItemEntry) Min() int32 {
	return d.min
}

func (d *DropItemEntry) SetMin(min int32) {
	d.min = min
}

func (d *DropItemEntry) Max() int32 {
	return d.max
}

func (d *DropItemEntry) SetMax(max int32) {
	d.max = max
}

func (d *DropItemEntry) Weight() int32 {
	return d.weight
}

func (d *DropItemEntry) SetWeight(weight int32) {
	d.weight = weight
}
