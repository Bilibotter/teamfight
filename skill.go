package tft

type stackType int

const (
	NoStk stackType = iota
	AsStk
	CastStk
	AtkStk
)

type skill struct {
	g        *ground
	apDmg    float64
	adDmg    float64
	stackBy  stackType
	perStack int
	stackDmg float64
	solidDmg int // 无加成伤害
	call     []func(g *ground)
	next     *skill
	mana     int
	wait     int // 施法前摇，单位100ms
}

func (s *skill) swing() int {
	return s.wait * tick / 10
}

func (s *skill) cost() int {
	g := s.g
	reduce, shrink := g.manaDec, g.manaSrk
	for _, a := range g.attributes {
		if a.valid() {
			reduce += a.body().manaDec
			shrink += a.body().manaSrk
		}
	}
	return (s.mana - reduce) * (100 - shrink) / 100
}

func (s *skill) cast() {
	for _, call := range s.call {
		call(s.g)
	}
	if s.next != nil {
		s.g.skill = s.next
	}
}

func (s *skill) normal(g *ground) {
	ap, ad := 0.0, 0.0
	if g.dmgType == phyHero {
		ap = g.magic() * (s.apDmg)
		ad = g.physics() * (s.adDmg + s.stack())
	} else if g.dmgType == magHero {
		ap = g.magic() * (s.apDmg + s.stack())
		ad = g.physics() * (s.adDmg)
	} else {
		panic("unknown skill type")
	}
	if g.critable() {
		amp := g.crit()
		ap *= amp
		ad *= amp
	}
	g.recordDmg((ap+ad)*g.spellAmp(), fromCast)
}

func (s *skill) stack() float64 {
	g := s.g
	switch s.stackBy {
	case AtkStk:
		return s.stackDmg * float64(g.atkTimes/s.perStack)
	case CastStk:
		// 施法前已将castTimes+1
		return s.stackDmg * float64((g.castTimes-1)/s.perStack)
	case AsStk:
		return s.stackDmg * float64(g.bonusAS()/s.perStack)
	case NoStk:
		return 0
	default:
		panic("unknown stack type")
	}
}

func (g *ground) Skill(mana int, dmg float64, minors ...float64) *ground {
	sk := &skill{mana: mana}
	minor := 0.0
	if len(minors) > 0 {
		minor = minors[0]
	}
	if g.dmgType == phyHero {
		sk.apDmg = minor
		sk.adDmg = dmg
	} else if g.dmgType == magHero {
		sk.apDmg = dmg
		sk.adDmg = minor
	} else {
		panic("unknown skill type")
	}
	sk.call = []func(g *ground){sk.normal}
	sk.wait = 5
	return g.addSkill(sk)
}

func (g *ground) Stack(typ stackType, stackDmg float64, perStacks ...int) *ground {
	// 默认叠加型不会有2阶段
	sk := g.tail()
	sk.stackBy = typ
	sk.stackDmg = stackDmg
	perStack := 1
	if len(perStacks) == 1 {
		perStack = perStacks[0]
	}
	sk.perStack = perStack
	return g
}

// 释放技能进入暴怒，暴怒期间无法获得法力
func (g *ground) Fury(duration int, reduce action, bf ...*attrs) *ground {
	g.tail().call = append(g.skill.call, addLockBuff(duration, reduce, bf...))
	return g
}

// 释放技能进入振奋，振奋期间可以获得法力
func (g *ground) Inspire(duration int, reduce action, bf ...*attrs) *ground {
	g.tail().call = append(g.skill.call, addFreeBuff(duration, reduce, bf...))
	return g
}

func (g *ground) Swing(_100ms int) *ground {
	g.tail().wait = _100ms
	return g
}

func (g *ground) Loop() *ground {
	g.tail().next = g.skill
	return g
}

func (g *ground) tail() *skill {
	tail := g.skill
	for ; tail.next != nil; tail = tail.next {
	}
	return tail
}

func (g *ground) addSkill(skill *skill) *ground {
	skill.g = g
	if g.skill == nil {
		g.skill = skill
		return g
	}
	tail := g.skill
	for ; tail.next != nil; tail = tail.next {
	}
	tail.next = skill
	return g
}
