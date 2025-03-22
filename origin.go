package tft

import "fmt"

type techItem int

const (
	RED techItem = iota
	BLUE
	GREEN
	YELLOW
	VIOLET
)

func (g *ground) template(actives ...int) *ground {
	active := 2
	if len(actives) != 0 {
		active = actives[0]
	}
	switch active {
	case 2:

	case 4:

	case 6:

	default:
		panic("wrong origin num")
	}
	p := attrPassive(TimeGoA, AP(10))
	g.addPassive(p)
	return g
}

func (g *ground) Faerie() *ground {
	g.ad += 30
	g.ap += 30
	g.OncePassive(AttackA, 12, AMP(30))
	return g
}

// 高频打击的短腿英雄会因为逛街而浪费输出时间
func (g *ground) Walk() *ground {
	g.walk = true
	return g
}

// 幻灵战队
func (g *ground) Animal(actives ...int) *ground {
	active := 3
	if len(actives) != 0 {
		active = actives[0]
	}
	switch active {
	case 3:
		g.Add(AMP(5))
	case 5:
		g.Add(AMP(10))
	case 7:
		g.Add(AMP(15))
	default:
		panic("wrong origin num")
	}
	return g
}

// 战地机甲
func (g *ground) BoomBots(actives ...int) *ground {
	active := 2
	if len(actives) != 0 {
		active = actives[0]
	}
	switch active {
	case 2:
		g.Add(AtkAmp(37))
		g.Add(CastAmp(37))
	case 4:
		g.Add(AtkAmp(82))
		g.Add(CastAmp(82))
	case 6:
		g.Add(AtkAmp(100))
		g.Add(CastAmp(100))
	default:
		panic("wrong origin num")
	}
	return g
}

// 最终boss
func (g *ground) Cyberboss(actives ...int) *ground {
	active := 2
	if len(actives) != 0 {
		active = actives[0]
	}
	switch active {
	case 2:
		g.Add(AP(15))
	case 3:
		g.Add(AP(25))
	case 4:
		g.Add(AP(30))
	default:
		panic("wrong origin num")
	}
	return g
}

func (g *ground) Cypher(actives ...int) *ground {
	active := 3
	if len(actives) != 0 {
		active = actives[0]
	}
	switch active {
	case 3:
		g.Add(AP(25))
		g.Add(AD(25))
	case 4:
		g.Add(AP(40))
		g.Add(AD(40))
	case 5:
		g.Add(AP(50))
		g.Add(AD(50))
	default:
		panic("wrong origin num")
	}
	return g
}

// 神法兵团本地人
func (g *ground) Divinicorps(actives int, atrs ...*attrs) *ground {
	increse := 100
	switch actives {
	case 1:
		increse = 100
	case 2:
		increse = 110
	case 3:
		increse = 125
	case 4:
		increse = 150
	case 5:
		increse = 175
	case 6:
		increse = 200
	case 7:
		increse = 225
	default:
		panic("wrong origin num")
	}
	increse *= 2
	for _, a := range atrs {
		a.factor = increse
		g.Add(a)
	}
	return g
}

// 福牛，未计算金币
func (g *ground) GoldenOx(actives ...int) *ground {
	active := 2
	if len(actives) != 0 {
		active = actives[0]
	}
	gold := 60
	switch active {
	case 2:
		g.Add(AMP(18 + gold*20/100))
	case 4:
		g.Add(AMP(20 + gold*40/100))
	case 6:
		g.Add(AMP(22 + gold*66/100))
	default:
		panic("wrong origin num")
	}
	return g
}

// 妖怪艺术家
// 本地人踩签名格
func (g *ground) StreetDemon(actives ...int) *ground {
	active := 3
	if len(actives) != 0 {
		active = actives[0]
	}
	increase := 3
	switch active {
	case 3:
		increase *= 6
	case 5:
		increase *= 10
	case 7:
		increase *= 15
	default:
		panic("wrong origin num")
	}
	g.Add(AP(increase))
	g.Add(AD(increase))
	return g
}

func (g *ground) Syndicate(actives ...int) *ground {
	active := 3
	if len(actives) != 0 {
		active = actives[0]
	}
	switch active {
	case 3:
		g.Add(AMP(5))
	case 5:
		g.Add(AMP(20))
	case 7:
		g.Add(AMP(30))
	default:
		panic("wrong origin num")
	}
	return g
}

// A.M.P
func Overload(actives ...int) float64 {
	active := 2
	if len(actives) != 0 {
		active = actives[0]
	}
	switch active {
	case 2:
		return 1
	case 3:
		return 2
	case 4:
		return 3
	case 5:
		return 4
	default:
		panic("wrong origin num")
	}
}

func (g *ground) Exotech(item techItem, actives ...int) *ground {
	active := 3
	if len(actives) != 0 {
		active = actives[0]
	}
	switch active {
	case 3:

	case 5:

	case 7:

	default:
		panic("wrong origin num")
	}
	return g
}

// 人造人本地人
func (g *ground) Dynamo(actives ...int) *ground {
	active := 2
	if len(actives) != 0 {
		active = actives[0]
	}
	p := &passive{
		Trigger: TimeGoA,
		Freq:    3,
	}
	restore := 0
	switch active {
	case 2:
		restore = 5 * 2
	case 3:
		restore = 8 * 2
	case 4:
		restore = 12 * 2
	default:
		panic("wrong origin num")
	}
	p.call = func(g *ground) {
		g.mana += restore
		if outputLevel >= 3 {
			fmt.Printf("%4.1f秒:人造人羁绊恢复%d点法力值\n", g.current(), restore)
		}
	}
	g.addPassive(p)
	return g
}

// 裁决使
func (g *ground) Executioner(actives ...int) *ground {
	active := 2
	if len(actives) != 0 {
		active = actives[0]
	}
	g.skillCrit = true
	switch active {
	case 2:
		g.Add(CritRate(25))
		g.Add(CritAmp(8))
	case 3:
		g.Add(CritRate(35))
		g.Add(CritAmp(15))
	case 4:
		g.Add(CritRate(45))
		g.Add(CritAmp(20))
	case 5:
		g.Add(CritRate(55))
		g.Add(CritAmp(25))
	default:
		panic("wrong origin num")
	}
	return g
}

// 强袭射手
func (g *ground) Marksman(actives ...int) *ground {
	active := 2
	if len(actives) != 0 {
		active = actives[0]
	}
	increase := 0
	switch active {
	case 2:
		increase = 18
	case 4:
		increase = 64
	default:
		panic("wrong origin num")
	}
	g.Add(AD(increase))
	g.OncePassive(TimeGoA, 8, AD(increase))
	return g
}

// 迅捷射手
func (g *ground) Rapidfire(actives ...int) *ground {
	active := 2
	if len(actives) != 0 {
		active = actives[0]
	}
	g.Add(AS(10))
	increase := 0
	switch active {
	case 2:
		increase = 4
	case 4:
		increase = 10
	case 6:
		increase = 22
	default:
		panic("wrong origin num")
	}
	g.StackPassive0(AttackA, 10, 1, AS(increase))
	return g
}

// 杀手
func (g *ground) Slayer(actives ...int) *ground {
	active := 2
	if len(actives) != 0 {
		active = actives[0]
	}
	switch active {
	case 2:
		g.Add(AD(15))
	case 4:
		g.Add(AD(40))
	case 6:
		g.Add(AD(70))
	default:
		panic("wrong origin num")
	}
	return g
}

// 战略分析师本地人
func (g *ground) Strategist(actives ...int) *ground {
	active := 2
	if len(actives) != 0 {
		active = actives[0]
	}
	switch active {
	case 2:
		g.Add(AMP(6 * 3))
	case 3:
		g.Add(AMP(9 * 3))
	case 4:
		g.Add(AMP(12 * 3))
	case 5:
		g.Add(AMP(15 * 3))
	default:
		panic("wrong origin num")
	}
	return g
}

// 战略分析师外地人
func (g *ground) Strategist0(actives ...int) *ground {
	active := 2
	if len(actives) != 0 {
		active = actives[0]
	}
	switch active {
	case 2:
		g.Add(AMP(6))
	case 3:
		g.Add(AMP(10))
	case 4:
		g.Add(AMP(14))
	case 5:
		g.Add(AMP(18))
	default:
		panic("wrong origin num")
	}
	return g
}

// 高级工程师
func (g *ground) Techie(actives ...int) *ground {
	active := 2
	if len(actives) != 0 {
		active = actives[0]
	}
	switch active {
	case 2:
		g.Add(AP(15))
	case 4:
		g.Add(AP(45))
	case 6:
		g.Add(AP(80))
	case 8:
		g.Add(AP(110))
	default:
		panic("wrong origin num")
	}
	return g
}
