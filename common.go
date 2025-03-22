package tft

func addBuff(remain int, a ...*attrs) func(*ground) {
	return func(g *ground) {
		g.buff(newB(remain, a...))
	}
}

func addMinion(remain int, reduce action, a ...*attrs) func(*ground) {
	return func(g *ground) {
		bf := newB(remain, a...)
		bf.reduce = reduce
		bf.unique = true
		g.buff(bf)
	}
}

func addFreeBuff(remain int, reduce action, a ...*attrs) func(*ground) {
	return func(g *ground) {
		bf := newB(remain, a...)
		bf.reduce = reduce
		g.buff(bf)
	}
}

func addLockBuff(remain int, reduce action, a ...*attrs) func(*ground) {
	return func(g *ground) {
		bf := newB(remain, a...)
		bf.reduce = reduce
		cp := g.buff(bf)
		g.locks = append(g.locks, cp)
	}
}
