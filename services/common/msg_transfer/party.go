package msg_transfer

import (
	"fmt"
	"strconv"

	ptpb "github.com/moke-game/platform/api/gen/party"
	ppb "github.com/moke-game/platform/api/gen/profile"

	bffpb "github.com/moke-game/game/api/gen/bff"
)

func Party(party *ptpb.PartyInfo) (*bffpb.PartyInfo, error) {
	if party == nil {
		return nil, fmt.Errorf("party is nil")
	}
	ret := &bffpb.PartyInfo{
		PartyMembers: make(map[int64]*bffpb.PartyMember),
	}
	ret.PlayId = party.Party.Type
	if party.Party != nil {
		ret.PartyId = party.Party.Id
		if uid, err := strconv.ParseInt(party.Party.Owner, 10, 64); err != nil {
			return nil, fmt.Errorf("parse owner:%s error: %w", party.Party.Owner, err)
		} else {
			ret.PartyOwner = uid
		}
	}
	for _, v := range party.GetMembers() {
		if member, err := PartyMember(v); err != nil {
			return nil, err
		} else {
			ret.PartyMembers[member.SimpleInfo.Uid] = member
		}
	}
	return ret, nil
}

func PartyMember(member *ptpb.Member) (*bffpb.PartyMember, error) {
	uid, err := strconv.ParseInt(member.GetUid(), 10, 64)
	if err != nil {
		return nil, err
	}
	sp := &bffpb.PlayerSimpleInfo{
		Uid:    uid,
		Name:   member.Nickname,
		Head:   member.Avatar,
		HeroId: member.HeroId,
		Online: !member.IsOffline,
	}
	return &bffpb.PartyMember{
		SimpleInfo: sp,
		Status:     bffpb.PartyMemberStatus(member.Status),
	}, nil
}

func ProfileToPartyMember(profile *ppb.ProfileBasic) *ptpb.Member {
	return &ptpb.Member{
		Uid:      profile.Uid,
		Nickname: profile.Nickname,
		Avatar:   profile.Avatar,
		HeroId:   profile.HeroId,
	}
}
