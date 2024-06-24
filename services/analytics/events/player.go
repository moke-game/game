package events

import pb "github.com/moke-game/platform/api/gen/profile"

type optionPlayer func(*Player)
type Player struct {
	Uid    string
	Name   string
	Avatar string
	HeroId int32
}

func NewPlayer(uid string, opts ...optionPlayer) *Player {
	p := &Player{
		Uid: uid,
	}
	for _, opt := range opts {
		opt(p)
	}
	return p
}

func WithProfile(profile *pb.Profile) optionPlayer {
	return func(p *Player) {
		if profile == nil {
			return
		}
		p.Name = profile.Nickname
		p.Avatar = profile.Avatar
		p.HeroId = profile.HeroId
	}
}
