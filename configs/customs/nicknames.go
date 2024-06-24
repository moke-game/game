package customs

import (
	"github.com/duke-git/lancet/v2/random"

	cfg "github.com/moke-game/game/configs/code"
)

type Nicknames struct {
	NamePrefixArr []string //前缀
	NameInfixArr  []string //词根
	NameSuffixArr []string //后缀
}

func CreateNickNames() *Nicknames {
	return &Nicknames{
		NamePrefixArr: make([]string, 0),
		NameInfixArr:  make([]string, 0),
		NameSuffixArr: make([]string, 0),
	}
}

func (n *Nicknames) Init(tbl *cfg.TblNickName) {
	for _, v := range tbl.GetDataMap() {
		n.NamePrefixArr = append(n.NamePrefixArr, v.Prefix)
		//n.NameInfixArr = append(n.NameInfixArr, v.Infix)
		//n.NameSuffixArr = append(n.NameSuffixArr, v.Suffix)
	}
}

func (n *Nicknames) Random() string {
	pIndex := random.RandInt(0, len(n.NamePrefixArr))
	//iIndex := random.RandInt(0, len(n.NameInfixArr))
	//sIndex := random.RandInt(0, len(n.NameSuffixArr))
	//return n.NamePrefixArr[pIndex] + n.NameInfixArr[iIndex] + n.NameSuffixArr[sIndex]
	return n.NamePrefixArr[pIndex]
}
