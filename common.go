package tft

func addBuff(remain int, a ...*attrs) func(*ground) {
	return func(g *ground) {
		g.buff(newB(remain, a...))
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
		g.locks = append(g.locks, g.buff(bf))
	}
}
