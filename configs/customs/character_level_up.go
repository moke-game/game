package customs

import cfg "github.com/moke-game/game/configs/code"

type CharacterLevelUp struct {
	characterLevelUp map[int32]map[int32]*cfg.CharacterLvup
}

func CreateCharacterLevelUp() *CharacterLevelUp {
	return &CharacterLevelUp{characterLevelUp: make(map[int32]map[int32]*cfg.CharacterLvup)}
}

func (c *CharacterLevelUp) Init(tblCharacterLvUp *cfg.TblCharacterLvup) {
	lvUpMap := make(map[int32]map[int32]*cfg.CharacterLvup)
	for _, data := range tblCharacterLvUp.GetDataList() {
		dt, ok := lvUpMap[data.HeroID]
		if !ok {
			dt = make(map[int32]*cfg.CharacterLvup)
			lvUpMap[data.HeroID] = dt
		}
		dt[data.Level] = data
	}
	c.characterLevelUp = lvUpMap
}

func (c *CharacterLevelUp) GetCharacterLevelUp(heroId, level int32) (*cfg.CharacterLvup, bool) {
	levelMap, ok := c.characterLevelUp[heroId]
	if !ok {
		return nil, ok
	}
	lvUpCfg, ok := levelMap[level]
	return lvUpCfg, ok
}

func (c *CharacterLevelUp) GetCharacterLevelUpMax(heroId int32) (*cfg.CharacterLvup, bool) {
	levelMap, ok := c.characterLevelUp[heroId]
	if !ok {
		return nil, ok
	}
	var max *cfg.CharacterLvup = nil
	for _, lvUp := range levelMap {
		if max == nil || lvUp.Level > max.Level {
			max = lvUp
		}
	}
	return max, ok
}
