package tft

import (
	"fmt"
	"math/rand"
)

type champion struct {
	attrs
	dmgType    int
	locks      []validI
	buffs      map[string]*buff
	passives   map[string]*passive
	attributes []attribute
	skill      *skill
	mana       int
	baseAtk    float64 // 基础攻击力
	baseSpeed  float64 // 基础攻速
	casting    bool    // true表示，施法中。施法中不获取法力值
}

func (g *champion) ability() int {
	ap := g.ap
	for _, a := range g.attributes {
		if a.valid() {
			ap += a.body().ap
		}
	}
	return ap
}

func (g *champion) speed() float64 {
	speed := g.baseSpeed
	haste := g.as
	for _, a := range g.attributes {
		if a.valid() {
			haste += a.body().as
		}
	}
	speed *= float64(haste) / 100
	if speed > 2.5 && !innerTesting {
		return 2.5
	}
	return speed
}

func (g *champion) atkMana() int {
	hitMana := g.hitMana
	for _, a := range g.attributes {
		if a.valid() {
			hitMana += a.body().hitMana
		}
	}
	return g.charge(hitMana)
}

func (g *champion) charge(mana int) int {
	amp := g.manaAmp
	for _, a := range g.attributes {
		if a.valid() {
			amp += a.body().manaAmp
		}
	}
	return mana * amp / 100
}

func (g *champion) bonusAS() int {
	bonus := g.as - 100
	for _, a := range g.attributes {
		if a.valid() {
			bonus += a.body().as
		}
	}
	return bonus
}

func (g *champion) critable() bool {
	if g.skillCrit {
		return true
	}
	for _, a := range g.attributes {
		if a.valid() && a.body().skillCrit {
			return true
		}
	}
	return false
}

func (g *champion) effect() float64 {
	ae := g.ae
	for _, a := range g.attributes {
		if a.valid() {
			ae += a.body().ae
		}
	}
	return float64(ae)
}

func (g *champion) crit() float64 {
	amp, rate := g.critAmp, g.critRate
	for _, a := range g.attributes {
		if a.valid() {
			rate += a.body().critRate
			amp += a.body().critAmp
		}
	}
	if rate > 100 {
		rate = 100
	}
	return float64(100+amp*rate/100) / 100
}

func (g *champion) magic() float64 {
	ap, amp := g.ap, g.amp
	for _, a := range g.attributes {
		if a.valid() {
			ap += a.body().ap
			amp += a.body().amp
		}
	}
	return float64(ap*amp) / 100 / 100
}

func (g *champion) physics() float64 {
	ad, amp := g.ad, g.amp
	for _, a := range g.attributes {
		if a.valid() {
			ad += a.body().ad
			amp += a.body().amp
		}
	}
	return float64(ad*amp) / 100 / 100
}

func (g *champion) spellAmp() float64 {
	amp := g.castAmp
	for _, a := range g.attributes {
		if a.valid() {
			amp += a.body().castAmp
		}
	}
	return float64(amp) / 100
}

// 双重打击
func (g *ground) double() bool {
	amp := g.strike
	for _, a := range g.attributes {
		if a.valid() {
			amp += a.body().strike
		}
	}
	if amp == 0 {
		return false
	}
	rate := float64(amp) / 100
	return rand.Float64() <= rate
}

// 普攻增强
func (g *champion) explode() float64 {
	amp := g.atkAmp
	for _, a := range g.attributes {
		if a.valid() {
			amp += a.body().atkAmp
		}
	}
	return float64(amp) / 100
}

func (g *champion) attackDmg() float64 {
	return g.crit() * g.baseAtk * g.physics() * g.explode()
}

// todo 消除重复代码
func (g *champion) effectDmg() float64 {
	effect := g.effect()
	if g.apAE() {
		effect *= g.magic()
	} else {
		effect *= g.physics()
	}
	if g.critable() {
		return g.crit() * effect
	}
	return effect
}

func (g *champion) attackFull() float64 {
	effect := g.effect()
	if g.apAE() {
		effect *= g.magic()
	} else {
		effect *= g.physics()
	}
	if g.critable() {
		return g.crit()*effect + g.attackDmg()
	}
	return effect + g.attackDmg()
}

func (g *ground) buff(bf *buff) *buff {
	key := bf.key()
	if outputLevel >= 3 {
		fmt.Printf("%2d.0秒:添加buff %s\n", g.now, key[:6])
	}
	var match *buff
	if old, ok := g.buffs[key]; ok {
		old.remain = bf.remain
		match = old
	} else {
		bf.g = g
		g.buffs[key] = bf
		g.attributes = append(g.attributes, bf)
		g.handlers = append(g.handlers, bf)
		bf.name = key
		match = bf
	}
	if match.reduce == TimeGoA {
		match.end = g.ticks + match.remain*tick
	}
	return match
}

func (g *ground) addPassive(p *passive) {
	key := p.key()
	if _, ok := g.passives[key]; ok {
		panic(fmt.Sprintf("被动重复%v", *p))
	} else {
		p.g = g
		g.passives[p.name] = p
		g.attributes = append(g.attributes, p)
		g.handlers = append(g.handlers, p)
		p.name = key
	}
}
