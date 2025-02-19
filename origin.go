package tft

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
	g.passive(p)
	return g
}

func (g *ground) Walk() *ground {
	g.walk = true
	return g
}
