package tft

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

func (g *ground) Dynamo(actives ...int) *ground {
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
	return g
}

func (g *ground) Executioner(actives ...int) *ground {
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
	return g
}

func (g *ground) Marksman(actives ...int) *ground {
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
	return g
}

func (g *ground) Rapidfire(actives ...int) *ground {
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
	return g
}

func (g *ground) Techie(actives ...int) *ground {
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
	return g
}
