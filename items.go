package tft

// 高频施法
func (g *ground) Invoker() *ground {
	if g.skill.cost() >= 60 {
		g.Shojin()
	} else {
		g.Blue()
	}
	g.Nashor().Jeweled()
	return g
}

// 爆发施法
func (g *ground) Sorcerer() *ground {
	if g.skill.cost() >= 60 {
		g.Shojin()
	} else {
		g.Blue()
	}
	g.Breaker().Jeweled()
	return g
}

func (g *ground) Fighter() *ground {
	g.Sword().sterak().Titan().beaten(9)
	return g
}

func (g *ground) Beaten(i ...int) *ground {
	if len(i) == 0 {
		g.beaten(9)
	} else {
		g.beaten(i[0])
	}
	return g
}

func (g *ground) Ranger() *ground {
	g.Guinsoo().Infinity().Whisper()
	return g
}

func (g *ground) Blaster() *ground {
	g.Shojin().Infinity().Whisper()
	return g
}

// 青龙刀
func (g *ground) Shojin() *ground {
	g.hitMana += 5
	g.ad += 15
	g.ap += 15
	g.mana += 15
	return g
}

// 羊刀
func (g *ground) Guinsoo() *ground {
	g.ap += 10
	g.as += 10
	p := stackPassive(AttackA, 1, AS(5))
	g.addPassive(p)
	return g
}

func (g *ground) DoubleGuinsoo() *ground {
	g.ap += 20
	g.as += 20
	p := stackPassive(AttackA, 1, AS(10))
	g.addPassive(p)
	return g
}

// 无尽
func (g *ground) Infinity() *ground {
	if g.critable() {
		g.critAmp += 10
	} else {
		g.skillCrit = true
	}
	g.critRate += 35
	g.ad += 35
	return g
}

// 法爆
func (g *ground) Jeweled() *ground {
	if g.critable() {
		g.critAmp += 10
	} else {
		g.skillCrit = true
	}
	g.critRate += 35
	g.ap += 35
	return g
}

// 炽天使
func (g *ground) Staff() *ground {
	g.ap += 20
	g.mana += 15
	p := stackPassive(TimeGoA, 5, AP(30))
	g.addPassive(p)
	return g
}

// 鬼书
func (g *ground) Book() *ground {
	g.as += 10
	g.ap += 25
	return g
}

// 挨揍回蓝
func (g *ground) beaten(i int) *ground {
	p := attrPassive(TimeGoA)
	p.call = func(g *ground) {
		restore := g.charge(i)
		if !g.locking() {
			g.mana += restore
		}
	}
	g.addPassive(p)
	return g
}

// 帽子
func (g *ground) DeathCap() *ground {
	g.ap += 50
	g.amp += 15
	return g
}

// 巨杀
func (g *ground) GiantSlayer() *ground {
	g.ad += 25
	g.ap += 25
	g.as += 10
	g.amp += 18
	return g
}

// 纳什
func (g *ground) Nashor() *ground {
	g.ap += 10
	g.as += 10
	p := buffPassive(BeforeCastA, 5, AS(50))
	g.addPassive(p)
	return g
}

// 轻语
func (g *ground) Whisper() *ground {
	g.ad += 15
	g.as += 20
	g.critRate += 20
	return g
}

// 蓝buff
func (g *ground) Blue() *ground {
	g.mana += 30
	g.ad += 15
	g.ap += 15

	p0 := buffPassive(AfterCastA, 8, AMP(5))

	f := func(g *ground) {
		restore := g.charge(10)
		g.mana += restore
	}
	p1 := attrPassive(AfterCastA)
	p1.call = f

	g.addPassive(p0)
	g.addPassive(p1)
	return g
}

// 饮血剑
func (g *ground) Sword() *ground {
	g.ad += 15
	g.ap += 15
	return g
}

// 正义
func (g *ground) Justice() *ground {
	g.mana += 15
	g.critRate += 20
	g.ad += 23
	g.ap += 23
	return g
}

// 血手
func (g *ground) sterak() *ground {
	g.ad += 15
	p := attrPassive(TimeGoA, AD(35))
	p.Left = 8
	p.Right = 30
	return g
}

// 泰坦
func (g *ground) Titan() *ground {
	g.as += 10
	p0 := stackPassive0(AttackA, 12, 1, AD(2), AP(2))
	p1 := stackPassive0(TimeGoA, 13, 1, AD(2), AP(2))
	g.addPassive(p0)
	g.addPassive(p1)
	return g
}

// 水银
func (g *ground) Silver() *ground {
	g.critRate += 20
	g.as += 30
	p := stackPassive0(TimeGoA, 9, 2, AS(3))
	g.addPassive(p)
	return g
}

// 破防
func (g *ground) Breaker() *ground {
	g.as += 20
	g.ap += 10
	g.critRate += 20
	g.amp += 10
	g.BuffPassive(BeforeCastA, 1, 3, AMP(15))
	return g
}

func (g *ground) Red() *ground {
	g.as += 35
	g.amp += 6
	return g
}

// 复活甲
func (g *ground) Night() *ground {
	g.as += 10
	p := attrPassive(TimeGoA, AS(15))
	p.Left = 8
	p.Right = 30
	g.addPassive(p)
	return g
}

// 科技枪
func (g *ground) Gunblade() *ground {
	g.ad += 20
	g.ap += 20
	return g
}

// 杀人剑
func (g *ground) DeathBlade() *ground {
	g.ad += 55
	g.amp += 10
	return g
}
