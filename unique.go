package tft

import "fmt"

func (g *ground) KBK() *ground {
	g.BuffPassive(BeforeCastA, 1, 8, AS(100))
	p := g.OncePassive(BeforeCastA, 1)
	p.call = func(g *ground) {
		g.mana += 100
	}
	return g
}

func (g *ground) YuumiAmp(actives ...int) *ground {
	active := 2
	if len(actives) != 0 {
		active = actives[0]
	}
	p := &passive{
		Trigger: TimeGoA,
		Freq:    1,
	}
	restore := 0
	switch active {
	case 2, 3, 4, 5:
		restore = (active - 1) * 3
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
