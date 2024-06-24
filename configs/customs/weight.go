package customs

import (
	"math/rand"
)

func RandomRange(min, max int32) int32 {
	if min == max {
		return min
	}
	a, b := int64(min), int64(max)
	if a > b {
		b, a = a, b
	}
	return int32(rand.Int63n(b-a+1) + a)
}

type IRandomWeightItem interface {
	GetWeight() int32
	GetRanWeight() int32
	SetRanWeight(weight int32)
	GetItem() any
}

type RandomWeightItem struct {
	Item      any
	Weight    int32
	RanWeight int32
}

func NewRandomWeightItem(item any, weight int32) IRandomWeightItem {
	return &RandomWeightItem{
		Item:   item,
		Weight: weight,
	}
}

func (r *RandomWeightItem) GetWeight() int32 {
	return r.Weight
}

func (r *RandomWeightItem) GetRanWeight() int32 {
	return r.RanWeight
}

func (r *RandomWeightItem) GetItem() any {
	return r.Item
}

func (r *RandomWeightItem) SetRanWeight(weight int32) {
	r.RanWeight = weight
}

func RandomWeight(pool []IRandomWeightItem) (any, bool) {
	total := int32(0)
	for _, p := range pool {
		total += p.GetWeight()
		p.SetRanWeight(total)
	}
	ran := RandomRange(0, int32(int(total)))
	for _, p := range pool {
		if p.GetRanWeight() >= ran {
			return p.GetItem(), true
		}
	}
	return nil, false
}
